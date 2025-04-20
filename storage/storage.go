package storage

import (
	"context"
	"sync"
	"time"

	"github.com/crosszan/a2a/errorx"
	"github.com/crosszan/a2a/schema"
)

// Storage is an interface for manage tasks and their associated data.
type Storage interface {
	// GetHistory returns the history of the session.
	// if historyLenght is less than 0, all history is returned.
	GetHistory(ctx context.Context, sessionID string, historyLenght int) ([]schema.Message, error)

	// AppendHistory appends a new message to the history.
	AppendHistory(ctx context.Context, sessionID string, message schema.Message) error

	// CreateTask creates a new task with the given name and description.
	CreateTask(ctx context.Context, task *schema.Task) error

	// GetTask returns the task for the given taskID.
	GetTask(ctx context.Context, taskID string) (string, error)

	// UpdateTask updates the task for the given taskID.
	UpdateStatus(ctx context.Context, taskID string, status schema.TaskStatus) error

	// UpdateArtifact updates the artifact for the given taskID.
	UpdateArtifact(ctx context.Context, taskID string, artifact schema.Artifact) error
}

type InMemoryStorage struct {
	mu       sync.RWMutex
	initOnce sync.Once

	tasks   map[string]*schema.Task
	history map[string][]schema.Message
}

func (s *InMemoryStorage) init() {
	s.initOnce.Do(func() {
		s.tasks = make(map[string]*schema.Task)
		s.history = make(map[string][]schema.Message)
	})
}

func (s *InMemoryStorage) CreateTask(ctx context.Context, task *schema.Task) error {
	s.init()
	s.mu.Lock()
	defer s.mu.Unlock()

	cloned := *task
	s.tasks[cloned.ID] = &cloned
	return nil
}

func (s *InMemoryStorage) GetTask(ctx context.Context, taskID string) (*schema.Task, error) {
	s.init()
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, ok := s.tasks[taskID]; ok {
		cloned := *task
		return &cloned, nil
	}
	return nil, errorx.ErrorTaskNotFound
}

func (s *InMemoryStorage) AppendHistory(ctx context.Context, sessionID string, message schema.Message) error {
	s.init()
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.history[sessionID]; !ok {
		s.history[sessionID] = []schema.Message{}
	}

	s.history[sessionID] = append(s.history[sessionID], message)
	return nil
}

func (s *InMemoryStorage) GetHistory(ctx context.Context, sessionID string, historyLenght int) ([]schema.Message, error) {
	s.init()
	s.mu.RLock()
	defer s.mu.RUnlock()

	if messages, ok := s.history[sessionID]; ok {
		if historyLenght < 0 {
			return messages, nil
		}
		if len(messages) > historyLenght {
			return messages[len(messages)-historyLenght:], nil
		}
		return messages, nil
	}
	return []schema.Message{}, nil
}

func (s *InMemoryStorage) UpdateStatus(ctx context.Context, taskID string, status schema.TaskStatus) error {
	s.init()
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[taskID]
	if !ok {
		return errorx.ErrorTaskNotFound
	}
	if status.Timestamp == nil {
		// if timestamp is not set, set it to the current time
		timestamp := time.Now().Format(time.RFC3339)
		status.Timestamp = &timestamp
	}
	if task.Status.Timestamp != nil {
		before, err := time.Parse(time.RFC3339, *task.Status.Timestamp)
		if err != nil {
			return err
		}
		after, err := time.Parse(time.RFC3339, *status.Timestamp)
		if err != nil {
			return err
		}
		if after.Before(before) {
			// ignore the update if the timestamp is before the previous timestamp
			return nil
		}
	}
	task.Status = status
	s.tasks[taskID] = task
	return nil
}

func (s *InMemoryStorage) UpdateArtifact(ctx context.Context, taskID string, artifact schema.Artifact) error {
	s.init()
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, ok := s.tasks[taskID]; ok {
		if artifact.Index < len(task.Artifacts) {
			task.Artifacts[artifact.Index] = artifact
		} else {
			artifacts := make([]schema.Artifact, artifact.Index+1)
			copy(artifacts, task.Artifacts)
			artifacts[artifact.Index] = artifact
			task.Artifacts = artifacts
		}
		s.tasks[taskID] = task
		return nil
	}
	return errorx.ErrorTaskNotFound
}

package storage

import (
	"context"
	"sync"

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
	})
}

package schema

type Task struct {
	ID        string                 `json:"id"`
	SessionID string                 `json:"sessionId,omitempty"`
	Status    TaskStatus             `json:"status"`
	History   []Message              `json:"history,omitempty"`
	Artifacts []Artifact             `json:"artifacts,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type TaskStatus struct {
}

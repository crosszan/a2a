package schema

type Artifact struct {
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	Parts       []Part         `json:"parts"`
	Index       int            `json:"index"`
	Append      *bool          `json:"append,omitempty"`
	LastChunk   *bool          `json:"last_chunk,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

type PartType string

const (
	PartTypeText PartType = "text"
	PartTypeFile PartType = "file"
	PartTypeData PartType = "data"
)

type FileContent struct {
	Name     *string `json:"name,omitempty"`
	MimeType *string `json:"mime_type,omitempty"`
	Bytes    *string `json:"bytes,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

type Part struct {
	Type     PartType       `json:"type"`
	Text     *string        `json:"text,omitempty"`
	File     *FileContent   `json:"file,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

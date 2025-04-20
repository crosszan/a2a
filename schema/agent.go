package schema

// AgentCard represents key information about an agent.
// Specification: https://google.github.io/A2A/#/documentation?id=agent-card
type AgentCard struct {
	// Human readable name of the agent.
	Name string `json:"name"`
	// Human readable description of the agent's functionality.
	Description string `json:"description"`
	// URL where agent is hosted.
	URL string `json:"url"`
	// Provider of the agent.
	Provider *AgentProvider `json:"provider"`
	// Version of the agent.
	Version string `json:"version"`
	// URL of the agent's documentation.
	DocumentationURL *string `json:"documentationUrl,omitempty"`
	// Capabilities of the agent.
	Capabilities Capabilities `json:"capabilities"`
	// Authentication information for the agent.
	Authentication Authentication `json:"authentication"`
	// Default input modes for the agent.
	DefaultInputModes []string `json:"defaultInputModes"`
	// Default output modes for the agent.
	DefaultOutputModes []string `json:"defaultOutputModes"`
	// Skills of the agent.
	Skills []Skill `json:"skills"`
}

type AgentProvider struct {
	// Organization that created the agent.
	Organization string `json:"organization"`
	// URL of the organization's website.
	URL string `json:"url"`
}

type Capabilities struct {
	// if true, the agent supports SSE.
	Streaming *bool `json:"streaming,omitempty"`
	// if true, the agent supports push notifications.
	PushNotifications *bool `json:"pushNotifications,omitempty"`
	// if true, the agent supports state transition history.
	StateTransitionHistory *bool `json:"stateTransitionHistory,omitempty"`
}

type Authentication struct {
	// Schemes of the authentication.
	// Supported schemes:
	// - Bearer Token
	// - Basic Authentication
	Schemes []string `json:"schemes"`
	// Credentials of the authentication.
	Credentials *string `json:"credentials,omitempty"`
}

type Skill struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Examples    []string `json:"examples,omitempty"`
	InputModes  []string `json:"inputModes"`
	OutputModes []string `json:"outputModes"`
}

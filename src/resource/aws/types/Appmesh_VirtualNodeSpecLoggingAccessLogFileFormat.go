package types

type Appmesh_VirtualNodeSpecLoggingAccessLogFileFormat struct {
	// The logging format for JSON.
	Jsons []Appmesh_VirtualNodeSpecLoggingAccessLogFileFormatJson `json:"jsons,omitempty" yaml:"jsons,omitempty"`

	// The logging format for text. Must be between 1 and 1000 characters in length.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}

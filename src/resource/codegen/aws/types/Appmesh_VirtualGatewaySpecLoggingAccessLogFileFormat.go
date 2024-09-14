package types

type Appmesh_VirtualGatewaySpecLoggingAccessLogFileFormat struct {
	// The logging format for JSON.
	Jsons []Appmesh_VirtualGatewaySpecLoggingAccessLogFileFormatJson `json:"jsons,omitempty" yaml:"jsons,omitempty"`

	// The logging format for text. Must be between 1 and 1000 characters in length.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}

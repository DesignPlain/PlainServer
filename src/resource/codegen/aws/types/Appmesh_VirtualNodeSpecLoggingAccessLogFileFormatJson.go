package types

type Appmesh_VirtualNodeSpecLoggingAccessLogFileFormatJson struct {
	// The specified value for the JSON. Must be between 1 and 100 characters in length.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The specified key for the JSON. Must be between 1 and 100 characters in length.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}

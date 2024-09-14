package types

type Medialive_InputSource struct {
	// The key used to extract the password from EC2 Parameter store.
	PasswordParam string `json:"passwordParam,omitempty" yaml:"passwordParam,omitempty"`

	// The URL where the stream is pulled from.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The username for the input source.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

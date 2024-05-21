package types

type Compute_SecurityPolicyAdvancedOptionsConfig struct {
	// Log level to use. Defaults to `NORMAL`.
	LogLevel string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`

	// An optional list of case-insensitive request header names to use for resolving the callers client IP address.
	UserIpRequestHeaders []string `json:"userIpRequestHeaders,omitempty" yaml:"userIpRequestHeaders,omitempty"`

	/*
	   Custom configuration to apply the JSON parsing. Only applicable when
	   `json_parsing` is set to `STANDARD`. Structure is documented below.
	*/
	JsonCustomConfig Compute_SecurityPolicyAdvancedOptionsConfigJsonCustomConfig `json:"jsonCustomConfig,omitempty" yaml:"jsonCustomConfig,omitempty"`

	// Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
	JsonParsing string `json:"jsonParsing,omitempty" yaml:"jsonParsing,omitempty"`
}

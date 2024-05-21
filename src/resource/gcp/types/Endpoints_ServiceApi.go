package types

type Endpoints_ServiceApi struct {
	// A list of Method objects; structure is documented below.
	Methods []Endpoints_ServiceApiMethod `json:"methods,omitempty" yaml:"methods,omitempty"`

	// The simple name of the endpoint as described in the config.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// `SYNTAX_PROTO2` or `SYNTAX_PROTO3`.
	Syntax string `json:"syntax,omitempty" yaml:"syntax,omitempty"`

	// A version string for this api. If specified, will have the form major-version.minor-version, e.g. `1.10`.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

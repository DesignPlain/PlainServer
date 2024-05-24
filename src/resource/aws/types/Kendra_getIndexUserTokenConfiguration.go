package types

type Kendra_getIndexUserTokenConfiguration struct {
	// A block that specifies the information about the JSON token type configuration.
	JsonTokenTypeConfigurations []Kendra_getIndexUserTokenConfigurationJsonTokenTypeConfiguration `json:"jsonTokenTypeConfigurations,omitempty" yaml:"jsonTokenTypeConfigurations,omitempty"`

	// A block that specifies the information about the JWT token type configuration.
	JwtTokenTypeConfigurations []Kendra_getIndexUserTokenConfigurationJwtTokenTypeConfiguration `json:"jwtTokenTypeConfigurations,omitempty" yaml:"jwtTokenTypeConfigurations,omitempty"`
}

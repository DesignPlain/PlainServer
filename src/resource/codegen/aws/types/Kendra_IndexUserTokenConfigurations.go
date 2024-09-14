package types

type Kendra_IndexUserTokenConfigurations struct {
	// A block that specifies the information about the JSON token type configuration. Detailed below.
	JsonTokenTypeConfiguration Kendra_IndexUserTokenConfigurationsJsonTokenTypeConfiguration `json:"jsonTokenTypeConfiguration,omitempty" yaml:"jsonTokenTypeConfiguration,omitempty"`

	// A block that specifies the information about the JWT token type configuration. Detailed below.
	JwtTokenTypeConfiguration Kendra_IndexUserTokenConfigurationsJwtTokenTypeConfiguration `json:"jwtTokenTypeConfiguration,omitempty" yaml:"jwtTokenTypeConfiguration,omitempty"`
}

package types

type Identityplatform_ConfigSmsRegionConfig struct {
	/*
	   A policy of allowing SMS to every region by default and adding disallowed regions to a disallow list.
	   Structure is documented below.
	*/
	AllowByDefault Identityplatform_ConfigSmsRegionConfigAllowByDefault `json:"allowByDefault,omitempty" yaml:"allowByDefault,omitempty"`

	/*
	   A policy of only allowing regions by explicitly adding them to an allowlist.
	   Structure is documented below.
	*/
	AllowlistOnly Identityplatform_ConfigSmsRegionConfigAllowlistOnly `json:"allowlistOnly,omitempty" yaml:"allowlistOnly,omitempty"`
}

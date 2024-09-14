package types

type Cognito_RiskConfigurationRiskExceptionConfiguration struct {
	/*
	   Overrides the risk decision to always block the pre-authentication requests.
	   The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	   Can contain a maximum of 200 items.
	*/
	BlockedIpRangeLists []string `json:"blockedIpRangeLists,omitempty" yaml:"blockedIpRangeLists,omitempty"`

	/*
	   Risk detection isn't performed on the IP addresses in this range list.
	   The IP range is in CIDR notation.
	   Can contain a maximum of 200 items.
	*/
	SkippedIpRangeLists []string `json:"skippedIpRangeLists,omitempty" yaml:"skippedIpRangeLists,omitempty"`
}

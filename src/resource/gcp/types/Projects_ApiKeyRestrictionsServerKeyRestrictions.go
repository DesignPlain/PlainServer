package types

type Projects_ApiKeyRestrictionsServerKeyRestrictions struct {
	// A list of the caller IP addresses that are allowed to make API calls with this key.
	AllowedIps []string `json:"allowedIps,omitempty" yaml:"allowedIps,omitempty"`
}

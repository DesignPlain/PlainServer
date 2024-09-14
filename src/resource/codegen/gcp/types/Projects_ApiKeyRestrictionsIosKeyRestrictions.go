package types

type Projects_ApiKeyRestrictionsIosKeyRestrictions struct {
	// A list of bundle IDs that are allowed when making API calls with this key.
	AllowedBundleIds []string `json:"allowedBundleIds,omitempty" yaml:"allowedBundleIds,omitempty"`
}

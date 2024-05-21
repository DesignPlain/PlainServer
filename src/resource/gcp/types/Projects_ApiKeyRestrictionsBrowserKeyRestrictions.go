package types

type Projects_ApiKeyRestrictionsBrowserKeyRestrictions struct {
	// A list of regular expressions for the referrer URLs that are allowed to make API calls with this key.
	AllowedReferrers []string `json:"allowedReferrers,omitempty" yaml:"allowedReferrers,omitempty"`
}

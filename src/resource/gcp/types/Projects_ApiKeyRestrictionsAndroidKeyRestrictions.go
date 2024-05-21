package types

type Projects_ApiKeyRestrictionsAndroidKeyRestrictions struct {
	// A list of Android applications that are allowed to make API calls with this key.
	AllowedApplications []Projects_ApiKeyRestrictionsAndroidKeyRestrictionsAllowedApplication `json:"allowedApplications,omitempty" yaml:"allowedApplications,omitempty"`
}

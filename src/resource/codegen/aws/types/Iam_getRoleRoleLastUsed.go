package types

type Iam_getRoleRoleLastUsed struct {
	// The date and time, in RFC 3339 format, that the role was last used.
	LastUsedDate string `json:"lastUsedDate,omitempty" yaml:"lastUsedDate,omitempty"`

	// The name of the AWS Region in which the role was last used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}

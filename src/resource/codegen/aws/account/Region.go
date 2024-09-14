package account

type Region struct {
	// The ID of the target account when managing member accounts. Will manage current user's account by default if omitted. To use this parameter, the caller must be an identity in the organization's management account or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have all features enabled, and the organization must have trusted access enabled for the Account Management service, and optionally a delegated admin account assigned.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Whether the region is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The region name to manage.
	RegionName string `json:"regionName,omitempty" yaml:"regionName,omitempty"`
}

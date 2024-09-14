package types

type Kendra_getIndexUserGroupResolutionConfiguration struct {
	// The identity store provider (mode) you want to use to fetch access levels of groups and users. AWS Single Sign-On is currently the only available mode. Your users and groups must exist in an AWS SSO identity source in order to use this mode. Valid Values are `AWS_SSO` or `NONE`.
	UserGroupResolutionMode string `json:"userGroupResolutionMode,omitempty" yaml:"userGroupResolutionMode,omitempty"`
}

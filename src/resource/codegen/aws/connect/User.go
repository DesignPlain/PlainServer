package connect

import types "libds/aws/types"

type User struct {
	// The identifier of the user account in the directory used for identity management. If Amazon Connect cannot access the directory, you can specify this identifier to authenticate users. If you include the identifier, we assume that Amazon Connect cannot access the directory. Otherwise, the identity information is used to authenticate users from your directory. This parameter is required if you are using an existing directory for identity management in Amazon Connect when Amazon Connect cannot access your directory to authenticate users. If you are using SAML for identity management and include this parameter, an error is returned.
	DirectoryUserId string `json:"directoryUserId,omitempty" yaml:"directoryUserId,omitempty"`

	// A block that contains information about the identity of the user. Documented below.
	IdentityInfo types.Connect_UserIdentityInfo `json:"identityInfo,omitempty" yaml:"identityInfo,omitempty"`

	// The user name for the account. For instances not using SAML for identity management, the user name can include up to 20 characters. If you are using SAML for identity management, the user name can include up to 64 characters from `[a-zA-Z0-9_-.\@]+`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// A block that contains information about the phone settings for the user. Documented below.
	PhoneConfig types.Connect_UserPhoneConfig `json:"phoneConfig,omitempty" yaml:"phoneConfig,omitempty"`

	// The identifier of the routing profile for the user.
	RoutingProfileId string `json:"routingProfileId,omitempty" yaml:"routingProfileId,omitempty"`

	// The identifier of the hierarchy group for the user.
	HierarchyGroupId string `json:"hierarchyGroupId,omitempty" yaml:"hierarchyGroupId,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// A list of identifiers for the security profiles for the user. Specify a minimum of 1 and maximum of 10 security profile ids. For more information, see [Best Practices for Security Profiles](https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-best-practices.html) in the Amazon Connect Administrator Guide.
	SecurityProfileIds []string `json:"securityProfileIds,omitempty" yaml:"securityProfileIds,omitempty"`

	/*
	   Tags to apply to the user. If configured with a provider
	   `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

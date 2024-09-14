package emr

type StudioSessionMapping struct {
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId string `json:"studioId,omitempty" yaml:"studioId,omitempty"`

	// The globally unique identifier (GUID) of the user or group from the Amazon Web Services SSO Identity Store.
	IdentityId string `json:"identityId,omitempty" yaml:"identityId,omitempty"`

	// The name of the user or group from the Amazon Web Services SSO Identity Store.
	IdentityName string `json:"identityName,omitempty" yaml:"identityName,omitempty"`

	// Specifies whether the identity to map to the Amazon EMR Studio is a `USER` or a `GROUP`.
	IdentityType string `json:"identityType,omitempty" yaml:"identityType,omitempty"`

	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. You should specify the ARN for the session policy that you want to apply, not the ARN of your user role.
	SessionPolicyArn string `json:"sessionPolicyArn,omitempty" yaml:"sessionPolicyArn,omitempty"`
}

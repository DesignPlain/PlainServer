package cognito

import types "DesignSphere_Server/src/resource/aws/types"

type IdentityPoolRoleAttachment struct {
	// A List of Role Mapping.
	RoleMappings []types.Cognito_IdentityPoolRoleAttachmentRoleMapping `json:"roleMappings,omitempty" yaml:"roleMappings,omitempty"`

	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles map[string]string `json:"roles,omitempty" yaml:"roles,omitempty"`

	// An identity pool ID in the format `REGION_GUID`.
	IdentityPoolId string `json:"identityPoolId,omitempty" yaml:"identityPoolId,omitempty"`
}

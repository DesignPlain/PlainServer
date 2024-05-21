package iam

import types "DesignSphere_Server/src/resource/gcp/types"

type WorkforcePoolProvider struct {
	/*
	   A [Common Expression Language](https://opensource.google/projects/cel) expression, in
	   plain text, to restrict what otherwise valid authentication credentials issued by the
	   provider should not be accepted.
	   The expression must output a boolean representing whether to allow the federation.
	   The following keywords may be referenced in the expressions:
	*/
	AttributeCondition string `json:"attributeCondition,omitempty" yaml:"attributeCondition,omitempty"`

	// A user-specified display name for the provider. Cannot exceed 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Represents an OpenId Connect 1.0 identity provider.
	   Structure is documented below.
	*/
	Oidc types.Iam_WorkforcePoolProviderOidc `json:"oidc,omitempty" yaml:"oidc,omitempty"`

	/*
	   Represents a SAML identity provider.
	   Structure is documented below.
	*/
	Saml types.Iam_WorkforcePoolProviderSaml `json:"saml,omitempty" yaml:"saml,omitempty"`

	/*
	   Maps attributes from the authentication credentials issued by an external identity provider
	   to Google Cloud attributes, such as `subject` and `segment`.
	   Each key must be a string specifying the Google Cloud IAM attribute to map to.
	   The following keys are supported:
	   - `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings.
	   This is also the subject that appears in Cloud Logging logs. This is a required field and
	   the mapped subject cannot exceed 127 bytes.
	   - `google.groups`: Groups the authenticating user belongs to. You can grant groups access to
	   resources using an IAM `principalSet` binding; access applies to all members of the group.
	   - `google.display_name`: The name of the authenticated user. This is an optional field and
	   the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead.
	   This attribute cannot be referenced in IAM bindings.
	   - `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo.
	   This is an optional field. When set, the image will be visible as the user's profile picture.
	   If not set, a generic user icon will be displayed instead.
	   This attribute cannot be referenced in IAM bindings.
	   You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute}
	   is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes.
	   The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_].
	   You can reference these attributes in IAM policies to define fine-grained access for a workforce pool
	   to Google Cloud resources. For example:
	   - `google.subject`:
	   `principal://iam.googleapis.com/locations/{location}/workforcePools/{pool}/subject/{value}`
	   - `google.groups`:
	   `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/group/{value}`
	   - `attribute.{custom_attribute}`:
	   `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/attribute.{custom_attribute}/{value}`
	   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	   function that maps an identity provider credential to the normalized attribute specified
	   by the corresponding map key.
	   You can use the `assertion` keyword in the expression to access a JSON representation of
	   the authentication credential issued by the provider.
	   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	   the total size of all mapped attributes must not exceed 8KB.
	   For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute.
	   For example, the following maps the sub claim of the incoming credential to the `subject` attribute
	   on a Google token:
	*/
	AttributeMapping map[string]string `json:"attributeMapping,omitempty" yaml:"attributeMapping,omitempty"`

	// A user-specified description of the provider. Cannot exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	   However, existing tokens still grant access.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// The location for the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID for the provider, which becomes the final component of the resource name.
	   This value must be 4-32 characters, and may contain the characters [a-z0-9-].
	   The prefix `gcp-` is reserved for use by Google, and may not be specified.


	   - - -
	*/
	ProviderId string `json:"providerId,omitempty" yaml:"providerId,omitempty"`

	/*
	   The ID to use for the pool, which becomes the final component of the resource name.
	   The IDs must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens.
	   It must start with a letter, and cannot have a trailing hyphen.
	   The prefix `gcp-` is reserved for use by Google, and may not be specified.
	*/
	WorkforcePoolId string `json:"workforcePoolId,omitempty" yaml:"workforcePoolId,omitempty"`
}

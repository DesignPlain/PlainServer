package iam

import types "DesignSphere_Server/src/resource/gcp/types"

type WorkloadIdentityPoolProvider struct {
	/*
	   The ID used for the pool, which is the final component of the pool resource name. This
	   value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	   `gcp-` is reserved for use by Google, and may not be specified.
	*/
	WorkloadIdentityPoolId string `json:"workloadIdentityPoolId,omitempty" yaml:"workloadIdentityPoolId,omitempty"`

	/*
	   An Amazon Web Services identity provider. Not compatible with the property oidc or saml.
	   Structure is documented below.
	*/
	Aws types.Iam_WorkloadIdentityPoolProviderAws `json:"aws,omitempty" yaml:"aws,omitempty"`

	/*
	   Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	   However, existing tokens still grant access.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   An OpenId Connect 1.0 identity provider. Not compatible with the property aws or saml.
	   Structure is documented below.
	*/
	Oidc types.Iam_WorkloadIdentityPoolProviderOidc `json:"oidc,omitempty" yaml:"oidc,omitempty"`

	/*
	   An SAML 2.0 identity provider. Not compatible with the property oidc or aws.
	   Structure is documented below.
	*/
	Saml types.Iam_WorkloadIdentityPoolProviderSaml `json:"saml,omitempty" yaml:"saml,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The ID for the provider, which becomes the final component of the resource name. This
	   value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	   `gcp-` is reserved for use by Google, and may not be specified.


	   - - -
	*/
	WorkloadIdentityPoolProviderId string `json:"workloadIdentityPoolProviderId,omitempty" yaml:"workloadIdentityPoolProviderId,omitempty"`

	/*
	   [A Common Expression Language](https://opensource.google/projects/cel) expression, in
	   plain text, to restrict what otherwise valid authentication credentials issued by the
	   provider should not be accepted.
	   The expression must output a boolean representing whether to allow the federation.
	   The following keywords may be referenced in the expressions:
	*/
	AttributeCondition string `json:"attributeCondition,omitempty" yaml:"attributeCondition,omitempty"`

	/*
	   Maps attributes from authentication credentials issued by an external identity provider
	   to Google Cloud attributes, such as `subject` and `segment`.
	   Each key must be a string specifying the Google Cloud IAM attribute to map to.
	   The following keys are supported:
	   - `google.subject`: The principal IAM is authenticating. You can reference this value
	   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
	   Cannot exceed 127 characters.
	   - `google.groups`: Groups the external identity belongs to. You can grant groups
	   access to resources using an IAM `principalSet` binding; access applies to all
	   members of the group.
	   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
	   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
	   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
	   is 100 characters, and the key may only contain the characters [a-z0-9_].
	   You can reference these attributes in IAM policies to define fine-grained access for a
	   workload to Google Cloud resources. For example:
	   - `google.subject`:
	   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
	   - `google.groups`:
	   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
	   - `attribute.{custom_attribute}`:
	   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
	   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	   function that maps an identity provider credential to the normalized attribute specified
	   by the corresponding map key.
	   You can use the `assertion` keyword in the expression to access a JSON representation of
	   the authentication credential issued by the provider.
	   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	   the total size of all mapped attributes must not exceed 8KB.
	   For AWS providers, the following rules apply:
	   - If no attribute mapping is defined, the following default mapping applies:
	*/
	AttributeMapping map[string]string `json:"attributeMapping,omitempty" yaml:"attributeMapping,omitempty"`

	// A description for the provider. Cannot exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}

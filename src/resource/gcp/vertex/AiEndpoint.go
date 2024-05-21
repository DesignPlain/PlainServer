package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiEndpoint struct {
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
	   Structure is documented below.
	*/
	EncryptionSpec types.Vertex_AiEndpointEncryptionSpec `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`

	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The description of the Endpoint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The location for the resource


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The region for the resource
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

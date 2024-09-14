package vertex

import types "libds/gcp/types"

type AiIndexEndpoint struct {
	/*
	   The labels with user-defined metadata to organize your Indexes.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
	   Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
	   [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`.
	   Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Optional. Configuration for private service connect. `network` and `privateServiceConnectConfig` are mutually exclusive.
	   Structure is documented below.
	*/
	PrivateServiceConnectConfig types.Vertex_AiIndexEndpointPrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty" yaml:"privateServiceConnectConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// If true, the deployed index will be accessible through public endpoint.
	PublicEndpointEnabled bool `json:"publicEndpointEnabled,omitempty" yaml:"publicEndpointEnabled,omitempty"`

	// The region of the index endpoint. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The description of the Index.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}

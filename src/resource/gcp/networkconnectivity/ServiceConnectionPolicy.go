package networkconnectivity

import types "DesignSphere_Server/src/resource/gcp/types"

type ServiceConnectionPolicy struct {
	/*
	   User-defined labels.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of the ServiceConnectionPolicy.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
	   Structure is documented below.
	*/
	PscConfig types.Networkconnectivity_ServiceConnectionPolicyPscConfig `json:"pscConfig,omitempty" yaml:"pscConfig,omitempty"`

	/*
	   The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass.
	   It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
	*/
	ServiceClass string `json:"serviceClass,omitempty" yaml:"serviceClass,omitempty"`

	// Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

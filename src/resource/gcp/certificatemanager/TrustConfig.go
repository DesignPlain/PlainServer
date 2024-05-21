package certificatemanager

import types "DesignSphere_Server/src/resource/gcp/types"

type TrustConfig struct {
	// A user-defined name of the trust config. Trust config names must be unique globally.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Set of trust stores to perform validation against.
	   This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation.
	   Structure is documented below.
	*/
	TrustStores []types.Certificatemanager_TrustConfigTrustStore `json:"trustStores,omitempty" yaml:"trustStores,omitempty"`

	// One or more paragraphs of text description of a trust config.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Set of label tags associated with the trust config.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The trust config location.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}

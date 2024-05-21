package securesourcemanager

import types "DesignSphere_Server/src/resource/gcp/types"

type Instance struct {
	/*
	   The name for the Instance.


	   - - -
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// Customer-managed encryption key name, in the format projects/-/locations/-/keyRings/-/cryptoKeys/-.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	/*
	   Labels as key value pairs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location for the Instance.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Private settings for private instance.
	   Structure is documented below.
	*/
	PrivateConfig types.Securesourcemanager_InstancePrivateConfig `json:"privateConfig,omitempty" yaml:"privateConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}

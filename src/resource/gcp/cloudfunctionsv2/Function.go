package cloudfunctionsv2

import types "DesignSphere_Server/src/resource/gcp/types"

type Function struct {
	// User-provided description of a function.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
	   It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   A set of key/value label pairs associated with this Cloud Function.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of this cloud function.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Describes the Service being deployed.
	   Structure is documented below.
	*/
	ServiceConfig types.Cloudfunctionsv2_FunctionServiceConfig `json:"serviceConfig,omitempty" yaml:"serviceConfig,omitempty"`

	/*
	   Describes the Build step of the function that builds a container
	   from the given source.
	   Structure is documented below.
	*/
	BuildConfig types.Cloudfunctionsv2_FunctionBuildConfig `json:"buildConfig,omitempty" yaml:"buildConfig,omitempty"`

	/*
	   An Eventarc trigger managed by Google Cloud Functions that fires events in
	   response to a condition in another service.
	   Structure is documented below.
	*/
	EventTrigger types.Cloudfunctionsv2_FunctionEventTrigger `json:"eventTrigger,omitempty" yaml:"eventTrigger,omitempty"`

	/*
	   A user-defined name of the function. Function names must
	   be unique globally and match pattern `projects/-/locations/-/functions/-`.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}

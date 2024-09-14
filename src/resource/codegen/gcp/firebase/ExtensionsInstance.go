package firebase

import types "libds/gcp/types"

type ExtensionsInstance struct {
	/*
	   The current Config of the Extension Instance.
	   Structure is documented below.
	*/
	Config types.Firebase_ExtensionsInstanceConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   The ID to use for the Extension Instance, which will become the final
	   component of the instance's name.
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}

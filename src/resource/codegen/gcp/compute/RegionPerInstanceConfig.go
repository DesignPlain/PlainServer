package compute

import types "libds/gcp/types"

type RegionPerInstanceConfig struct {
	/*
	   When true, deleting this config will immediately remove the underlying instance.
	   When false, deleting this config will use the behavior as determined by remove_instance_on_destroy.
	*/
	RemoveInstanceOnDestroy bool `json:"removeInstanceOnDestroy,omitempty" yaml:"removeInstanceOnDestroy,omitempty"`

	/*
	   The minimal action to perform on the instance during an update.
	   Default is `NONE`. Possible values are:
	   - REPLACE
	   - RESTART
	   - REFRESH
	   - NONE
	*/
	MinimalAction string `json:"minimalAction,omitempty" yaml:"minimalAction,omitempty"`

	// The name for this per-instance config and its corresponding instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The region instance group manager this instance config is part of.


	   - - -
	*/
	RegionInstanceGroupManager string `json:"regionInstanceGroupManager,omitempty" yaml:"regionInstanceGroupManager,omitempty"`

	/*
	   When true, deleting this config will immediately remove any specified state from the underlying instance.
	   When false, deleting this config will -not- immediately remove any state from the underlying instance.
	   State will be removed on the next instance recreation or update.
	*/
	RemoveInstanceStateOnDestroy bool `json:"removeInstanceStateOnDestroy,omitempty" yaml:"removeInstanceStateOnDestroy,omitempty"`

	/*
	   The most disruptive action to perform on the instance during an update.
	   Default is `REPLACE`. Possible values are:
	   - REPLACE
	   - RESTART
	   - REFRESH
	   - NONE
	*/
	MostDisruptiveAllowedAction string `json:"mostDisruptiveAllowedAction,omitempty" yaml:"mostDisruptiveAllowedAction,omitempty"`

	/*
	   The preserved state for this instance.
	   Structure is documented below.
	*/
	PreservedState types.Compute_RegionPerInstanceConfigPreservedState `json:"preservedState,omitempty" yaml:"preservedState,omitempty"`

	// Region where the containing instance group manager is located
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}

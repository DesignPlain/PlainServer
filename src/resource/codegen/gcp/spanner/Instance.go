package spanner

import types "libds/gcp/types"

type Instance struct {
	/*
	   An object containing a list of "key": value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The autoscaling configuration. Autoscaling is enabled if this field is set.
	   When autoscaling is enabled, num_nodes and processing_units are treated as,
	   OUTPUT_ONLY fields and reflect the current compute capacity allocated to
	   the instance.
	   Structure is documented below.
	*/
	AutoscalingConfig types.Spanner_InstanceAutoscalingConfig `json:"autoscalingConfig,omitempty" yaml:"autoscalingConfig,omitempty"`

	/*
	   The descriptive name for this instance as it appears in UIs. Must be
	   unique per project and between 4 and 30 characters in length.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   When deleting a spanner instance, this boolean option will delete all backups of this instance.
	   This must be set to true if you created a backup manually in the console.
	*/
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	/*
	   A unique identifier for the instance, which cannot be changed after
	   the instance is created. The name must be between 6 and 30 characters
	   in length.

	   If not provided, a random string starting with `tf-` will be selected.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The number of nodes allocated to this instance. Exactly one of either node_count or processing_units must be present in
	   terraform.
	*/
	NumNodes int `json:"numNodes,omitempty" yaml:"numNodes,omitempty"`

	/*
	   The number of processing units allocated to this instance. Exactly one of processing_units or node_count must be present
	   in terraform.
	*/
	ProcessingUnits int `json:"processingUnits,omitempty" yaml:"processingUnits,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The name of the instance's configuration (similar but not
	   quite the same as a region) which defines the geographic placement and
	   replication of your databases in this instance. It determines where your data
	   is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
	   In order to obtain a valid list please consult the
	   [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	*/
	Config string `json:"config,omitempty" yaml:"config,omitempty"`
}

package compute

import (
	types "DesignSphere_Server/src/resource/gcp/types"
)

type Autoscaler struct {
	// URL of the zone where the instance group resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   The configuration parameters for the autoscaling algorithm. You can
	   define one or more of the policies for an autoscaler: cpuUtilization,
	   customMetricUtilizations, and loadBalancingUtilization.
	   If none of these are specified, the default will be to autoscale based
	   on cpuUtilization to 0.6 or 60%!!(MISSING)
	   (MISSING)Structure is documented below.
	*/
	AutoscalingPolicy types.Compute_AutoscalerAutoscalingPolicy `json:"autoscalingPolicy,omitempty" yaml:"autoscalingPolicy,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. The name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// URL of the managed instance group that this autoscaler will scale.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}

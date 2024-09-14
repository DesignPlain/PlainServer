package cloudrunv2

import types "libds/gcp/types"

type Job struct {
	// Arbitrary version identifier for the API client.
	ClientVersion string `json:"clientVersion,omitempty" yaml:"clientVersion,omitempty"`

	// Name of the Job.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The template used to create executions for this Job.
	   Structure is documented below.
	*/
	Template types.Cloudrunv2_JobTemplate `json:"template,omitempty" yaml:"template,omitempty"`

	/*
	   Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.
	   Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected on new resources.
	   All system annotations in v1 now have a corresponding field in v2 Job.
	   This field follows Kubernetes annotations' namespacing, limits, and rules.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   Settings for the Binary Authorization feature.
	   Structure is documented below.
	*/
	BinaryAuthorization types.Cloudrunv2_JobBinaryAuthorization `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`

	// Arbitrary identifier for the API client.
	Client string `json:"client,omitempty" yaml:"client,omitempty"`

	/*
	   Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component,
	   environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.
	   Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected.
	   All system labels in v1 now have a corresponding field in v2 Job.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.
	   If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.
	   For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
	   Possible values are: `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, `DEPRECATED`.
	*/
	LaunchStage string `json:"launchStage,omitempty" yaml:"launchStage,omitempty"`

	// The location of the cloud run job
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}

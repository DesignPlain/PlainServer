package cloudrunv2

import types "DesignSphere_Server/src/resource/gcp/types"

type Service struct {
	/*
	   Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100%!!(MISSING)t(MISSING)raffic to the latest Ready Revision.
	   Structure is documented below.
	*/
	Traffics []types.Cloudrunv2_ServiceTraffic `json:"traffics,omitempty" yaml:"traffics,omitempty"`

	// Arbitrary identifier for the API client.
	Client string `json:"client,omitempty" yaml:"client,omitempty"`

	/*
	   Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component,
	   environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.
	   Cloud Run API v2 does not support labels with  `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected.
	   All system labels in v1 now have a corresponding field in v2 Service.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Name of the Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.
	   If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.
	   For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
	   Possible values are: `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, `DEPRECATED`.
	*/
	LaunchStage string `json:"launchStage,omitempty" yaml:"launchStage,omitempty"`

	/*
	   Settings for the Binary Authorization feature.
	   Structure is documented below.
	*/
	BinaryAuthorization types.Cloudrunv2_ServiceBinaryAuthorization `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`

	// Arbitrary version identifier for the API client.
	ClientVersion string `json:"clientVersion,omitempty" yaml:"clientVersion,omitempty"`

	/*
	   One or more custom audiences that you want this service to support. Specify each custom audience as the full URL in a string. The custom audiences are encoded in the token and used to authenticate requests.
	   For more information, see https://cloud.google.com/run/docs/configuring/custom-audiences.
	*/
	CustomAudiences []string `json:"customAudiences,omitempty" yaml:"customAudiences,omitempty"`

	// User-provided description of the Service. This field currently has a 512-character limit.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
	   Possible values are: `INGRESS_TRAFFIC_ALL`, `INGRESS_TRAFFIC_INTERNAL_ONLY`, `INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER`.
	*/
	Ingress string `json:"ingress,omitempty" yaml:"ingress,omitempty"`

	/*
	   Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.
	   Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected in new resources.
	   All system annotations in v1 now have a corresponding field in v2 Service.
	   This field follows Kubernetes annotations' namespacing, limits, and rules.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// The location of the cloud run service
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The template used to create revisions for this Service.
	   Structure is documented below.
	*/
	Template types.Cloudrunv2_ServiceTemplate `json:"template,omitempty" yaml:"template,omitempty"`
}

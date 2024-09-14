package bigquery

import types "libds/gcp/types"

type AppProfile struct {
	/*
	   The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]-`.


	   - - -
	*/
	AppProfileId string `json:"appProfileId,omitempty" yaml:"appProfileId,omitempty"`

	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings bool `json:"ignoreWarnings,omitempty" yaml:"ignoreWarnings,omitempty"`

	/*
	   If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	   in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	   consistency to improve availability.
	*/
	MultiClusterRoutingUseAny bool `json:"multiClusterRoutingUseAny,omitempty" yaml:"multiClusterRoutingUseAny,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The standard options used for isolating this app profile's traffic from other use cases.
	   Structure is documented below.
	*/
	StandardIsolation types.Bigquery_AppProfileStandardIsolation `json:"standardIsolation,omitempty" yaml:"standardIsolation,omitempty"`

	// Long form description of the use case for this app profile.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the instance to create the app profile within.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
	   clusters are eligible.
	*/
	MultiClusterRoutingClusterIds []string `json:"multiClusterRoutingClusterIds,omitempty" yaml:"multiClusterRoutingClusterIds,omitempty"`

	/*
	   Use a single-cluster routing policy.
	   Structure is documented below.
	*/
	SingleClusterRouting types.Bigquery_AppProfileSingleClusterRouting `json:"singleClusterRouting,omitempty" yaml:"singleClusterRouting,omitempty"`
}

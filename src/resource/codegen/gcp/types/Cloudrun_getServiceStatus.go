package types

type Cloudrun_getServiceStatus struct {
	/*
	   From ConfigurationStatus. LatestCreatedRevisionName is the last revision that was created
	   from this Service's Configuration. It might not be ready yet, for that use
	   LatestReadyRevisionName.
	*/
	LatestCreatedRevisionName string `json:"latestCreatedRevisionName,omitempty" yaml:"latestCreatedRevisionName,omitempty"`

	/*
	   From ConfigurationStatus. LatestReadyRevisionName holds the name of the latest Revision
	   stamped out from this Service's Configuration that has had its "Ready" condition become
	   "True".
	*/
	LatestReadyRevisionName string `json:"latestReadyRevisionName,omitempty" yaml:"latestReadyRevisionName,omitempty"`

	/*
	   ObservedGeneration is the 'Generation' of the Route that was last processed by the
	   controller.

	   Clients polling for completed reconciliation should poll until observedGeneration =
	   metadata.generation and the Ready condition's status is True or False.
	*/
	ObservedGeneration int `json:"observedGeneration,omitempty" yaml:"observedGeneration,omitempty"`

	/*
	   Traffic specifies how to distribute traffic over a collection of Knative Revisions
	   and Configurations
	*/
	Traffics []Cloudrun_getServiceStatusTraffic `json:"traffics,omitempty" yaml:"traffics,omitempty"`

	/*
	   From RouteStatus. URL holds the url that will distribute traffic over the provided traffic
	   targets. It generally has the form
	   https://{route-hash}-{project-hash}-{cluster-level-suffix}.a.run.app
	*/
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// Array of observed Service Conditions, indicating the current ready state of the service.
	Conditions []Cloudrun_getServiceStatusCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

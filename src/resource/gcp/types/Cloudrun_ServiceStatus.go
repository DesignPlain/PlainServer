package types

type Cloudrun_ServiceStatus struct {
	/*
	   (Output)
	   From ConfigurationStatus. LatestCreatedRevisionName is the last revision that was created
	   from this Service's Configuration. It might not be ready yet, for that use
	   LatestReadyRevisionName.
	*/
	LatestCreatedRevisionName string `json:"latestCreatedRevisionName,omitempty" yaml:"latestCreatedRevisionName,omitempty"`

	/*
	   (Output)
	   From ConfigurationStatus. LatestReadyRevisionName holds the name of the latest Revision
	   stamped out from this Service's Configuration that has had its "Ready" condition become
	   "True".
	*/
	LatestReadyRevisionName string `json:"latestReadyRevisionName,omitempty" yaml:"latestReadyRevisionName,omitempty"`

	/*
	   (Output)
	   ObservedGeneration is the 'Generation' of the Route that was last processed by the
	   controller.
	   Clients polling for completed reconciliation should poll until observedGeneration =
	   metadata.generation and the Ready condition's status is True or False.
	*/
	ObservedGeneration int `json:"observedGeneration,omitempty" yaml:"observedGeneration,omitempty"`

	/*
	   Traffic specifies how to distribute traffic over a collection of Knative Revisions
	   and Configurations
	   Structure is documented below.
	*/
	Traffics []Cloudrun_ServiceStatusTraffic `json:"traffics,omitempty" yaml:"traffics,omitempty"`

	/*
	   (Output)
	   URL displays the URL for accessing tagged traffic targets. URL is displayed in status,
	   and is disallowed on spec. URL must contain a scheme (e.g. http://) and a hostname,
	   but may not contain anything else (e.g. basic auth, url path, etc.)
	*/
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	/*
	   (Output)
	   Array of observed Service Conditions, indicating the current ready state of the service.
	   Structure is documented below.
	*/
	Conditions []Cloudrun_ServiceStatusCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

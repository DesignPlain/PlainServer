package types

type Appengine_FlexibleAppVersionEndpointsApiService struct {
	/*
	   Endpoints service configuration ID as specified by the Service Management API. For example "2016-09-19r1".
	   By default, the rollout strategy for Endpoints is "FIXED". This means that Endpoints starts up with a particular configuration ID.
	   When a new configuration is rolled out, Endpoints must be given the new configuration ID. The configId field is used to give the configuration ID
	   and is required in this case.
	   Endpoints also has a rollout strategy called "MANAGED". When using this, Endpoints fetches the latest configuration and does not need
	   the configuration ID. In this case, configId must be omitted.
	*/
	ConfigId string `json:"configId,omitempty" yaml:"configId,omitempty"`

	// Enable or disable trace sampling. By default, this is set to false for enabled.
	DisableTraceSampling bool `json:"disableTraceSampling,omitempty" yaml:"disableTraceSampling,omitempty"`

	/*
	   Endpoints service name which is the name of the "service" resource in the Service Management API.
	   For example "myapi.endpoints.myproject.cloud.goog"
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Endpoints rollout strategy. If FIXED, configId must be specified. If MANAGED, configId must be omitted.
	   Default value is `FIXED`.
	   Possible values are: `FIXED`, `MANAGED`.
	*/
	RolloutStrategy string `json:"rolloutStrategy,omitempty" yaml:"rolloutStrategy,omitempty"`
}

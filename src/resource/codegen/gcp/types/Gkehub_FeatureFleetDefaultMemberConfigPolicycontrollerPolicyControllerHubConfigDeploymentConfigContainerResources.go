package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigContainerResources struct {
	/*
	   Limits describes the maximum amount of compute resources allowed for use by the running container.
	   Structure is documented below.
	*/
	Limits Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigContainerResourcesLimits `json:"limits,omitempty" yaml:"limits,omitempty"`

	/*
	   Requests describes the amount of compute resources reserved for the container by the kube-scheduler.
	   Structure is documented below.
	*/
	Requests Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigContainerResourcesRequests `json:"requests,omitempty" yaml:"requests,omitempty"`
}

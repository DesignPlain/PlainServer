package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfig struct {
	// The identifier for this object. Format specified above.
	Component string `json:"component,omitempty" yaml:"component,omitempty"`

	/*
	   Container resource requirements.
	   Structure is documented below.
	*/
	ContainerResources Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigContainerResources `json:"containerResources,omitempty" yaml:"containerResources,omitempty"`

	/*
	   Pod affinity configuration.
	   Possible values are: `AFFINITY_UNSPECIFIED`, `NO_AFFINITY`, `ANTI_AFFINITY`.
	*/
	PodAffinity string `json:"podAffinity,omitempty" yaml:"podAffinity,omitempty"`

	/*
	   Pod tolerations of node taints.
	   Structure is documented below.
	*/
	PodTolerations []Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigPodToleration `json:"podTolerations,omitempty" yaml:"podTolerations,omitempty"`

	// Pod replica count.
	ReplicaCount int `json:"replicaCount,omitempty" yaml:"replicaCount,omitempty"`
}

package types

type Clouddeploy_TargetAnthosCluster struct {
	// Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
	Membership string `json:"membership,omitempty" yaml:"membership,omitempty"`
}

package types

type Gkehub_FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBinding struct {
	/*
	   The relative resource name of the binauthz platform policy to audit. GKE
	   platform policies have the following format:
	   `projects/{project_number}/platforms/gke/policies/{policy_id}`.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

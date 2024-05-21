package types

type Gkehub_MembershipEndpoint struct {
	/*
	   If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	   Structure is documented below.
	*/
	GkeCluster Gkehub_MembershipEndpointGkeCluster `json:"gkeCluster,omitempty" yaml:"gkeCluster,omitempty"`
}

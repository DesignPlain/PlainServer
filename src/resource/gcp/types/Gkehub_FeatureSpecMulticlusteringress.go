package types

type Gkehub_FeatureSpecMulticlusteringress struct {
	// Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: `projects/foo-proj/locations/global/memberships/bar`
	ConfigMembership string `json:"configMembership,omitempty" yaml:"configMembership,omitempty"`
}

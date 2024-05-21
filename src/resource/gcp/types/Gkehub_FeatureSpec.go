package types

type Gkehub_FeatureSpec struct {
	/*
	   Clusterupgrade feature spec.
	   Structure is documented below.
	*/
	Clusterupgrade Gkehub_FeatureSpecClusterupgrade `json:"clusterupgrade,omitempty" yaml:"clusterupgrade,omitempty"`

	/*
	   Fleet Observability feature spec.
	   Structure is documented below.
	*/
	Fleetobservability Gkehub_FeatureSpecFleetobservability `json:"fleetobservability,omitempty" yaml:"fleetobservability,omitempty"`

	/*
	   Multicluster Ingress-specific spec.
	   Structure is documented below.
	*/
	Multiclusteringress Gkehub_FeatureSpecMulticlusteringress `json:"multiclusteringress,omitempty" yaml:"multiclusteringress,omitempty"`
}

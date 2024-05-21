package types

type Edgecontainer_ClusterMaintenancePolicy struct {
	/*
	   Specifies the maintenance window in which maintenance may be performed.
	   Structure is documented below.
	*/
	Window Edgecontainer_ClusterMaintenancePolicyWindow `json:"window,omitempty" yaml:"window,omitempty"`
}

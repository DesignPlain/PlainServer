package types

type Dns_ManagedZonePrivateVisibilityConfig struct {
	//
	Networks []Dns_ManagedZonePrivateVisibilityConfigNetwork `json:"networks,omitempty" yaml:"networks,omitempty"`

	/*
	   The list of Google Kubernetes Engine clusters that can see this zone.
	   Structure is documented below.
	*/
	GkeClusters []Dns_ManagedZonePrivateVisibilityConfigGkeCluster `json:"gkeClusters,omitempty" yaml:"gkeClusters,omitempty"`
}

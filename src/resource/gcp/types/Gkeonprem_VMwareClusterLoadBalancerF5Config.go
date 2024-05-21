package types

type Gkeonprem_VMwareClusterLoadBalancerF5Config struct {
	/*
	   (Output)
	   The vCenter IP address.
	*/
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	/*
	   he preexisting partition to be used by the load balancer. T
	   his partition is usually created for the admin cluster for example:
	   'my-f5-admin-partition'.
	*/
	Partition string `json:"partition,omitempty" yaml:"partition,omitempty"`

	// The pool name. Only necessary, if using SNAT.
	SnatPool string `json:"snatPool,omitempty" yaml:"snatPool,omitempty"`
}

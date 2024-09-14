package memcache

import types "libds/gcp/types"

type Instance struct {
	// Number of nodes in the memcache instance.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// A user-visible name for the instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   User-specified parameters for this memcache instance.
	   Structure is documented below.
	*/
	MemcacheParameters types.Memcache_InstanceMemcacheParameters `json:"memcacheParameters,omitempty" yaml:"memcacheParameters,omitempty"`

	/*
	   The major version of Memcached software. If not provided, latest supported version will be used.
	   Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	   determined by our system based on the latest supported minor version.
	   Default value is `MEMCACHE_1_5`.
	   Possible values are: `MEMCACHE_1_5`, `MEMCACHE_1_6_15`.
	*/
	MemcacheVersion string `json:"memcacheVersion,omitempty" yaml:"memcacheVersion,omitempty"`

	/*
	   Configuration for memcache nodes.
	   Structure is documented below.
	*/
	NodeConfig types.Memcache_InstanceNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	/*
	   The full name of the GCE network to connect the instance to.  If not provided,
	   'default' will be used.
	*/
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty" yaml:"authorizedNetwork,omitempty"`

	/*
	   Maintenance policy for an instance.
	   Structure is documented below.
	*/
	MaintenancePolicy types.Memcache_InstanceMaintenancePolicy `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`

	// The resource name of the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Contains the name of allocated IP address ranges associated with
	   the private service access connection for example, "test-default"
	   associated with IP range 10.0.0.0/29.
	*/
	ReservedIpRangeIds []string `json:"reservedIpRangeIds,omitempty" yaml:"reservedIpRangeIds,omitempty"`

	/*
	   Zones where memcache nodes should be provisioned.  If not
	   provided, all zones will be used.
	*/
	Zones []string `json:"zones,omitempty" yaml:"zones,omitempty"`
}

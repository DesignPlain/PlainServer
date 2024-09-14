package bigtable

import types "libds/gcp/types"

type GCPolicy struct {
	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersions []types.Bigtable_GCPolicyMaxVersion `json:"maxVersions,omitempty" yaml:"maxVersions,omitempty"`

	// The name of the table.
	Table string `json:"table,omitempty" yaml:"table,omitempty"`

	// The name of the column family.
	ColumnFamily string `json:"columnFamily,omitempty" yaml:"columnFamily,omitempty"`

	/*
	   The deletion policy for the GC policy.
	   Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted in a replicated instance.

	   Possible values are: `ABANDON`.

	   -----
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	// Serialized JSON object to represent a more complex GC policy. Conflicts with `mode`, `max_age` and `max_version`. Conflicts with `mode`, `max_age` and `max_version`.
	GcRules string `json:"gcRules,omitempty" yaml:"gcRules,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the Bigtable instance.
	InstanceName string `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`

	// GC policy that applies to all cells older than the given age.
	MaxAge types.Bigtable_GCPolicyMaxAge `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`

	// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}

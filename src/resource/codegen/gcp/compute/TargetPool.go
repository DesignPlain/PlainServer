package compute

type TargetPool struct {
	/*
	   URL to the backup target pool. Must also set
	   failover\_ratio.
	*/
	BackupPool string `json:"backupPool,omitempty" yaml:"backupPool,omitempty"`

	/*
	   List of zero or one health check name or self_link. Only
	   legacy `gcp.compute.HttpHealthCheck` is supported.
	*/
	HealthChecks string `json:"healthChecks,omitempty" yaml:"healthChecks,omitempty"`

	/*
	   List of instances in the pool. They can be given as
	   URLs, or in the form of "zone/name". Note that the instances need not exist
	   at the time of target pool creation, so there is no need to use the
	   interpolation to create a dependency on the instances from the
	   target pool.
	*/
	Instances []string `json:"instances,omitempty" yaml:"instances,omitempty"`

	/*
	   Where the target pool resides. Defaults to project
	   region.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The resource URL for the security policy associated with this target pool.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	// Textual description field.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Ratio (0 to 1) of failed nodes before using the
	   backup pool (which must also be set).
	*/
	FailoverRatio float64 `json:"failoverRatio,omitempty" yaml:"failoverRatio,omitempty"`

	/*
	   A unique name for the resource, required by GCE. Changing
	   this forces a new resource to be created.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   How to distribute load. Options are "NONE" (no
	   affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	   "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	*/
	SessionAffinity string `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`
}

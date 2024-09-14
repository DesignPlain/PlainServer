package netapp

type StoragePool struct {
	// Capacity of the storage pool (in GiB).
	CapacityGib string `json:"capacityGib,omitempty" yaml:"capacityGib,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specifies the CMEK policy to be used for volume encryption. Format: `projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}`.
	   The policy needs to be in the same location as the storage pool.
	*/
	KmsConfig string `json:"kmsConfig,omitempty" yaml:"kmsConfig,omitempty"`

	/*
	   Labels as key value pairs. Example: `{ "owner": "Bob", "department": "finance", "purpose": "testing" }`.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   When enabled, the volumes uses Active Directory as LDAP name service for UID/GID lookups. Required to enable extended group support for NFSv3,
	   using security identifiers for NFSv4.1 or principal names for kerberized NFSv4.1.
	*/
	LdapEnabled bool `json:"ldapEnabled,omitempty" yaml:"ldapEnabled,omitempty"`

	/*
	   The resource name of the storage pool. Needs to be unique per location.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// VPC network name with format: `projects/{{project}}/global/networks/{{network}}`
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies the Active Directory policy to be used. Format: `projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}`.
	   The policy needs to be in the same location as the storage pool.
	*/
	ActiveDirectory string `json:"activeDirectory,omitempty" yaml:"activeDirectory,omitempty"`

	// Name of the location. Usually a region name, expect for some STANDARD service level pools which require a zone name.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Service level of the storage pool.
	   Possible values are: `PREMIUM`, `EXTREME`, `STANDARD`.
	*/
	ServiceLevel string `json:"serviceLevel,omitempty" yaml:"serviceLevel,omitempty"`
}

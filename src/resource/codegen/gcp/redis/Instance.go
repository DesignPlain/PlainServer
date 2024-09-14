package redis

import types "libds/gcp/types"

type Instance struct {
	/*
	   The connection mode of the Redis instance.
	   Default value is `DIRECT_PEERING`.
	   Possible values are: `DIRECT_PEERING`, `PRIVATE_SERVICE_ACCESS`.
	*/
	ConnectMode string `json:"connectMode,omitempty" yaml:"connectMode,omitempty"`

	/*
	   The version of Redis software. If not provided, latest supported
	   version will be used. Please check the API documentation linked
	   at the top for the latest valid values.
	*/
	RedisVersion string `json:"redisVersion,omitempty" yaml:"redisVersion,omitempty"`

	/*
	   Optional. The number of replica nodes. The valid range for the Standard Tier with
	   read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
	   for a Standard Tier instance, the only valid value is 1 and the default is 1.
	   The valid value for basic tier is 0 and the default is also 0.
	*/
	ReplicaCount int `json:"replicaCount,omitempty" yaml:"replicaCount,omitempty"`

	/*
	   The full name of the Google Compute Engine network to which the
	   instance is connected. If left unspecified, the default network
	   will be used.
	*/
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty" yaml:"authorizedNetwork,omitempty"`

	/*
	   Optional. The KMS key reference that you want to use to encrypt the data at rest for this Redis
	   instance. If this is provided, CMEK is enabled.
	*/
	CustomerManagedKey string `json:"customerManagedKey,omitempty" yaml:"customerManagedKey,omitempty"`

	// An arbitrary and optional user-provided name for the instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Resource labels to represent user provided metadata.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The ID of the instance or a fully qualified identifier for the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Persistence configuration for an instance.
	   Structure is documented below.
	*/
	PersistenceConfig types.Redis_InstancePersistenceConfig `json:"persistenceConfig,omitempty" yaml:"persistenceConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
	   - SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication
	   Default value is `DISABLED`.
	   Possible values are: `SERVER_AUTHENTICATION`, `DISABLED`.
	*/
	TransitEncryptionMode string `json:"transitEncryptionMode,omitempty" yaml:"transitEncryptionMode,omitempty"`

	/*
	   Only applicable to STANDARD_HA tier which protects the instance
	   against zonal failures by provisioning it across two zones.
	   If provided, it must be a different zone from the one provided in
	   [locationId].
	*/
	AlternativeLocationId string `json:"alternativeLocationId,omitempty" yaml:"alternativeLocationId,omitempty"`

	/*
	   Optional. Indicates whether OSS Redis AUTH is enabled for the
	   instance. If set to "true" AUTH is enabled on the instance.
	   Default value is "false" meaning AUTH is disabled.
	*/
	AuthEnabled bool `json:"authEnabled,omitempty" yaml:"authEnabled,omitempty"`

	/*
	   The zone where the instance will be provisioned. If not provided,
	   the service will choose a zone for the instance. For STANDARD_HA tier,
	   instances will be created across two zones for protection against
	   zonal failures. If [alternativeLocationId] is also provided, it must
	   be different from [locationId].
	*/
	LocationId string `json:"locationId,omitempty" yaml:"locationId,omitempty"`

	/*
	   Optional. Read replica mode. Can only be specified when trying to create the instance.
	   If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
	   - READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the
	   instance cannot scale up or down the number of replicas.
	   - READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance
	   can scale up and down the number of replicas.
	   Possible values are: `READ_REPLICAS_DISABLED`, `READ_REPLICAS_ENABLED`.
	*/
	ReadReplicasMode string `json:"readReplicasMode,omitempty" yaml:"readReplicasMode,omitempty"`

	/*
	   Redis configuration parameters, according to http://redis.io/topics/config.
	   Please check Memorystore documentation for the list of supported parameters:
	   https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	*/
	RedisConfigs map[string]string `json:"redisConfigs,omitempty" yaml:"redisConfigs,omitempty"`

	/*
	   The CIDR range of internal addresses that are reserved for this
	   instance. If not provided, the service will choose an unused /29
	   block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	   unique and non-overlapping with existing subnets in an authorized
	   network.
	*/
	ReservedIpRange string `json:"reservedIpRange,omitempty" yaml:"reservedIpRange,omitempty"`

	/*
	   The service tier of the instance. Must be one of these values:
	   - BASIC: standalone instance
	   - STANDARD_HA: highly available primary/replica instances
	   Default value is `BASIC`.
	   Possible values are: `BASIC`, `STANDARD_HA`.
	*/
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	/*
	   Maintenance policy for an instance.
	   Structure is documented below.
	*/
	MaintenancePolicy types.Redis_InstanceMaintenancePolicy `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`

	/*
	   Redis memory size in GiB.


	   - - -
	*/
	MemorySizeGb int `json:"memorySizeGb,omitempty" yaml:"memorySizeGb,omitempty"`

	// The name of the Redis region of the instance.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Optional. Additional IP range for node placement. Required when enabling read replicas on
	   an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
	   "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
	   range associated with the private service access connection, or "auto".
	*/
	SecondaryIpRange string `json:"secondaryIpRange,omitempty" yaml:"secondaryIpRange,omitempty"`
}

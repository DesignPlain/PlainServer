package neptune

type ClusterInstance struct {
	// The neptune engine version. Currently configuring this argumnet has no effect.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// The identifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// Determines whether a final DB snapshot is created before the DB instance is deleted.
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" yaml:"skipFinalSnapshot,omitempty"`

	// The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// The port on which the DB accepts connections. Defaults to `8182`.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Bool to control if instance is publicly accessible. Default is `false`.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix string `json:"identifierPrefix,omitempty" yaml:"identifierPrefix,omitempty"`

	// The instance class to use.
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier int `json:"promotionTier,omitempty" yaml:"promotionTier,omitempty"`

	// A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The EC2 Availability Zone that the neptune instance is created in.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The identifier of the `aws.neptune.Cluster` in which to launch this instance.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The name of the neptune parameter group to associate with this instance.
	NeptuneParameterGroupName string `json:"neptuneParameterGroupName,omitempty" yaml:"neptuneParameterGroupName,omitempty"`

	// A subnet group to associate with this neptune instance. --NOTE:-- This must match the `neptune_subnet_group_name` of the attached `aws.neptune.Cluster`.
	NeptuneSubnetGroupName string `json:"neptuneSubnetGroupName,omitempty" yaml:"neptuneSubnetGroupName,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" yaml:"preferredBackupWindow,omitempty"`

	/*
	   The window to perform maintenance in.
	   Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	*/
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	/*
	   Specifies whether any instance modifications
	   are applied immediately, or during the next maintenance window. Default is`false`.
	*/
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`
}

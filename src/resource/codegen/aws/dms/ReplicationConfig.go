package dms

import types "libds/aws/types"

type ReplicationConfig struct {
	// Unique identifier that you want to use to create the config.
	ReplicationConfigIdentifier string `json:"replicationConfigIdentifier,omitempty" yaml:"replicationConfigIdentifier,omitempty"`

	// An escaped JSON string that are used to provision this replication configuration. For example, [Change processing tuning settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.ChangeProcessingTuning.html)
	ReplicationSettings string `json:"replicationSettings,omitempty" yaml:"replicationSettings,omitempty"`

	// Unique value or name that you set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource. For more information, see [Fine-grained access control using resource names and tags](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.FineGrainedAccess)
	ResourceIdentifier string `json:"resourceIdentifier,omitempty" yaml:"resourceIdentifier,omitempty"`

	// JSON settings for specifying supplemental data. For more information see [Specifying supplemental data for task settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.TaskData.html)
	SupplementalSettings string `json:"supplementalSettings,omitempty" yaml:"supplementalSettings,omitempty"`

	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings string `json:"tableMappings,omitempty" yaml:"tableMappings,omitempty"`

	// Configuration block for provisioning an DMS Serverless replication.
	ComputeConfig types.Dms_ReplicationConfigComputeConfig `json:"computeConfig,omitempty" yaml:"computeConfig,omitempty"`

	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	ReplicationType string `json:"replicationType,omitempty" yaml:"replicationType,omitempty"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn string `json:"sourceEndpointArn,omitempty" yaml:"sourceEndpointArn,omitempty"`

	// Whether to run or stop the serverless replication, default is false.
	StartReplication bool `json:"startReplication,omitempty" yaml:"startReplication,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn string `json:"targetEndpointArn,omitempty" yaml:"targetEndpointArn,omitempty"`
}

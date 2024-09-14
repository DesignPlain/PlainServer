package dms

type ReplicationTask struct {
	// Escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings string `json:"tableMappings,omitempty" yaml:"tableMappings,omitempty"`

	// ARN that uniquely identifies the target endpoint.
	TargetEndpointArn string `json:"targetEndpointArn,omitempty" yaml:"targetEndpointArn,omitempty"`

	// RFC3339 formatted date string or UNIX timestamp for the start of the Change Data Capture (CDC) operation.
	CdcStartTime string `json:"cdcStartTime,omitempty" yaml:"cdcStartTime,omitempty"`

	// Migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType string `json:"migrationType,omitempty" yaml:"migrationType,omitempty"`

	// ARN of the replication instance.
	ReplicationInstanceArn string `json:"replicationInstanceArn,omitempty" yaml:"replicationInstanceArn,omitempty"`

	// Replication task identifier which must contain from 1 to 255 alphanumeric characters or hyphens, first character must be a letter, cannot end with a hyphen, and cannot contain two consecutive hyphens.
	ReplicationTaskId string `json:"replicationTaskId,omitempty" yaml:"replicationTaskId,omitempty"`

	// Escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html). Note that `Logging.CloudWatchLogGroup` and `Logging.CloudWatchLogStream` are read only and should not be defined, even as `null`, in the configuration since AWS provides a value for these settings.
	ReplicationTaskSettings string `json:"replicationTaskSettings,omitempty" yaml:"replicationTaskSettings,omitempty"`

	// Indicates when you want a change data capture (CDC) operation to start. The value can be a RFC3339 formatted date, a checkpoint, or a LSN/SCN format depending on the source engine. For more information see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
	CdcStartPosition string `json:"cdcStartPosition,omitempty" yaml:"cdcStartPosition,omitempty"`

	// A friendly name for the resource identifier at the end of the EndpointArn response parameter that is returned in the created Endpoint object.
	ResourceIdentifier string `json:"resourceIdentifier,omitempty" yaml:"resourceIdentifier,omitempty"`

	// ARN that uniquely identifies the source endpoint.
	SourceEndpointArn string `json:"sourceEndpointArn,omitempty" yaml:"sourceEndpointArn,omitempty"`

	// Whether to run or stop the replication task.
	StartReplicationTask bool `json:"startReplicationTask,omitempty" yaml:"startReplicationTask,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

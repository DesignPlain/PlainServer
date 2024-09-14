package types

type Ecs_ClusterConfiguration struct {
	// Details of the execute command configuration. See `execute_command_configuration` Block for details.
	ExecuteCommandConfiguration Ecs_ClusterConfigurationExecuteCommandConfiguration `json:"executeCommandConfiguration,omitempty" yaml:"executeCommandConfiguration,omitempty"`

	// Details of the managed storage configuration. See `managed_storage_configuration` Block for details.
	ManagedStorageConfiguration Ecs_ClusterConfigurationManagedStorageConfiguration `json:"managedStorageConfiguration,omitempty" yaml:"managedStorageConfiguration,omitempty"`
}

package types

type Databasemigrationservice_ConnectionProfileAlloydbSettingsPrimaryInstanceSettings struct {
	/*
	   Configuration for the machines that host the underlying database engine.
	   Structure is documented below.
	*/
	MachineConfig Databasemigrationservice_ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfig `json:"machineConfig,omitempty" yaml:"machineConfig,omitempty"`

	/*
	   (Output)
	   Output only. The private IP address for the Instance. This is the connection endpoint for an end-user application.
	*/
	PrivateIp string `json:"privateIp,omitempty" yaml:"privateIp,omitempty"`

	// Database flags to pass to AlloyDB when DMS is creating the AlloyDB cluster and instances. See the AlloyDB documentation for how these can be used.
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty" yaml:"databaseFlags,omitempty"`

	// The database username.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Labels for the AlloyDB primary instance created by DMS.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

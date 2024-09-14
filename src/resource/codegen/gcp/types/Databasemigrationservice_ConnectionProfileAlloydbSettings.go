package types

type Databasemigrationservice_ConnectionProfileAlloydbSettings struct {
	// Labels for the AlloyDB cluster created by DMS.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Settings for the cluster's primary instance
	   Structure is documented below.
	*/
	PrimaryInstanceSettings Databasemigrationservice_ConnectionProfileAlloydbSettingsPrimaryInstanceSettings `json:"primaryInstanceSettings,omitempty" yaml:"primaryInstanceSettings,omitempty"`

	/*
	   Required. The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster.
	   It is specified in the form: 'projects/{project_number}/global/networks/{network_id}'. This is required to create a cluster.
	*/
	VpcNetwork string `json:"vpcNetwork,omitempty" yaml:"vpcNetwork,omitempty"`

	/*
	   Required. Input only. Initial user to setup during cluster creation.
	   Structure is documented below.
	*/
	InitialUser Databasemigrationservice_ConnectionProfileAlloydbSettingsInitialUser `json:"initialUser,omitempty" yaml:"initialUser,omitempty"`
}

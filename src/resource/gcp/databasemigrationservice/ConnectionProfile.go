package databasemigrationservice

import types "DesignSphere_Server/src/resource/gcp/types"

type ConnectionProfile struct {
	/*
	   Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination database instance.
	   Structure is documented below.
	*/
	Cloudsql types.Databasemigrationservice_ConnectionProfileCloudsql `json:"cloudsql,omitempty" yaml:"cloudsql,omitempty"`

	// The connection profile display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Specifies connection parameters required specifically for Oracle databases.
	   Structure is documented below.
	*/
	Oracle types.Databasemigrationservice_ConnectionProfileOracle `json:"oracle,omitempty" yaml:"oracle,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
	   Structure is documented below.
	*/
	Alloydb types.Databasemigrationservice_ConnectionProfileAlloydb `json:"alloydb,omitempty" yaml:"alloydb,omitempty"`

	/*
	   The ID of the connection profile.


	   - - -
	*/
	ConnectionProfileId string `json:"connectionProfileId,omitempty" yaml:"connectionProfileId,omitempty"`

	/*
	   The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location where the connection profile should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Specifies connection parameters required specifically for MySQL databases.
	   Structure is documented below.
	*/
	Mysql types.Databasemigrationservice_ConnectionProfileMysql `json:"mysql,omitempty" yaml:"mysql,omitempty"`

	/*
	   Specifies connection parameters required specifically for PostgreSQL databases.
	   Structure is documented below.
	*/
	Postgresql types.Databasemigrationservice_ConnectionProfilePostgresql `json:"postgresql,omitempty" yaml:"postgresql,omitempty"`
}

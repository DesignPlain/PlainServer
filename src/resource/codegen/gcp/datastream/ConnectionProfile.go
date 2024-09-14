package datastream

import types "libds/gcp/types"

type ConnectionProfile struct {
	/*
	   Cloud Storage bucket profile.
	   Structure is documented below.
	*/
	GcsProfile types.Datastream_ConnectionProfileGcsProfile `json:"gcsProfile,omitempty" yaml:"gcsProfile,omitempty"`

	/*
	   The name of the location this connection profile is located in.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   PostgreSQL database profile.
	   Structure is documented below.
	*/
	PostgresqlProfile types.Datastream_ConnectionProfilePostgresqlProfile `json:"postgresqlProfile,omitempty" yaml:"postgresqlProfile,omitempty"`

	// The connection profile identifier.
	ConnectionProfileId string `json:"connectionProfileId,omitempty" yaml:"connectionProfileId,omitempty"`

	/*
	   Forward SSH tunnel connectivity.
	   Structure is documented below.
	*/
	ForwardSshConnectivity types.Datastream_ConnectionProfileForwardSshConnectivity `json:"forwardSshConnectivity,omitempty" yaml:"forwardSshConnectivity,omitempty"`

	/*
	   Labels.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   MySQL database profile.
	   Structure is documented below.
	*/
	MysqlProfile types.Datastream_ConnectionProfileMysqlProfile `json:"mysqlProfile,omitempty" yaml:"mysqlProfile,omitempty"`

	/*
	   Oracle database profile.
	   Structure is documented below.
	*/
	OracleProfile types.Datastream_ConnectionProfileOracleProfile `json:"oracleProfile,omitempty" yaml:"oracleProfile,omitempty"`

	/*
	   Private connectivity.
	   Structure is documented below.
	*/
	PrivateConnectivity types.Datastream_ConnectionProfilePrivateConnectivity `json:"privateConnectivity,omitempty" yaml:"privateConnectivity,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// BigQuery warehouse profile.
	BigqueryProfile types.Datastream_ConnectionProfileBigqueryProfile `json:"bigqueryProfile,omitempty" yaml:"bigqueryProfile,omitempty"`

	// Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}

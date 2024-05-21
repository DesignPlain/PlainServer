package types

type Databasemigrationservice_ConnectionProfileCloudsql struct {
	/*
	   (Output)
	   Output only. The Cloud SQL instance ID that this connection profile is associated with.
	*/
	CloudSqlId string `json:"cloudSqlId,omitempty" yaml:"cloudSqlId,omitempty"`

	/*
	   (Output)
	   Output only. The Cloud SQL database instance's private IP.
	*/
	PrivateIp string `json:"privateIp,omitempty" yaml:"privateIp,omitempty"`

	/*
	   (Output)
	   Output only. The Cloud SQL database instance's public IP.
	*/
	PublicIp string `json:"publicIp,omitempty" yaml:"publicIp,omitempty"`

	/*
	   Immutable. Metadata used to create the destination Cloud SQL database.
	   Structure is documented below.
	*/
	Settings Databasemigrationservice_ConnectionProfileCloudsqlSettings `json:"settings,omitempty" yaml:"settings,omitempty"`
}

package types

type Databasemigrationservice_ConnectionProfileAlloydb struct {
	// Required. The AlloyDB cluster ID that this connection profile is associated with.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	/*
	   Immutable. Metadata used to create the destination AlloyDB cluster.
	   Structure is documented below.
	*/
	Settings Databasemigrationservice_ConnectionProfileAlloydbSettings `json:"settings,omitempty" yaml:"settings,omitempty"`
}

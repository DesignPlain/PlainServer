package types

type Composer_getEnvironmentConfigDatabaseConfig struct {
	// Optional. Cloud SQL machine type used by Airflow database. It has to be one of: db-n1-standard-2, db-n1-standard-4, db-n1-standard-8 or db-n1-standard-16. If not specified, db-n1-standard-2 will be used.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// Optional. Cloud SQL database preferred zone.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

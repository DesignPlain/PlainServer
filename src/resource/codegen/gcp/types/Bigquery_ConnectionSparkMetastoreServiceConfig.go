package types

type Bigquery_ConnectionSparkMetastoreServiceConfig struct {
	// Resource name of an existing Dataproc Metastore service in the form of projects/[projectId]/locations/[region]/services/[serviceId].
	MetastoreService string `json:"metastoreService,omitempty" yaml:"metastoreService,omitempty"`
}

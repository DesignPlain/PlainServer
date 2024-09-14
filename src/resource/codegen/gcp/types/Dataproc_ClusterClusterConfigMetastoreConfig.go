package types

type Dataproc_ClusterClusterConfigMetastoreConfig struct {
	/*
	   Resource name of an existing Dataproc Metastore service.

	   Only resource names including projectid and location (region) are valid. Examples:

	   `projects/[projectId]/locations/[dataproc_region]/services/[service-name]`
	*/
	DataprocMetastoreService string `json:"dataprocMetastoreService,omitempty" yaml:"dataprocMetastoreService,omitempty"`
}

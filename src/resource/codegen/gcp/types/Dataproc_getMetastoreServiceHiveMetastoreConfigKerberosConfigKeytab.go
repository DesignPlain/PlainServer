package types

type Dataproc_getMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab struct {
	/*
	   The relative resource name of a Secret Manager secret version, in the following form:

	   "projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}".
	*/
	CloudSecret string `json:"cloudSecret,omitempty" yaml:"cloudSecret,omitempty"`
}

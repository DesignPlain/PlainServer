package types

type Dataproc_MetastoreFederationBackendMetastore struct {
	/*
	   The type of the backend metastore.
	   Possible values are: `METASTORE_TYPE_UNSPECIFIED`, `DATAPROC_METASTORE`, `BIGQUERY`.

	   - - -
	*/
	MetastoreType string `json:"metastoreType,omitempty" yaml:"metastoreType,omitempty"`

	// The relative resource name of the metastore that is being federated. The formats of the relative resource names for the currently supported metastores are listed below: Dataplex: projects/{projectId}/locations/{location}/lakes/{lake_id} BigQuery: projects/{projectId} Dataproc Metastore: projects/{projectId}/locations/{location}/services/{serviceId}
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The identifier for this object. Format specified above.
	Rank string `json:"rank,omitempty" yaml:"rank,omitempty"`
}

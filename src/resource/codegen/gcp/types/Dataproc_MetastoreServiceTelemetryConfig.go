package types

type Dataproc_MetastoreServiceTelemetryConfig struct {
	/*
	   The output format of the Dataproc Metastore service's logs.
	   Default value is `JSON`.
	   Possible values are: `LEGACY`, `JSON`.
	*/
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`
}

package types

type Dataproc_getMetastoreServiceTelemetryConfig struct {
	// The output format of the Dataproc Metastore service's logs. Default value: "JSON" Possible values: ["LEGACY", "JSON"]
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`
}

package types

type Glue_CrawlerDeltaTarget struct {
	// Specifies whether to write the manifest files to the Delta table path.
	WriteManifest bool `json:"writeManifest,omitempty" yaml:"writeManifest,omitempty"`

	// The name of the connection to use to connect to the Delta table target.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.
	CreateNativeDeltaTable bool `json:"createNativeDeltaTable,omitempty" yaml:"createNativeDeltaTable,omitempty"`

	// A list of the Amazon S3 paths to the Delta tables.
	DeltaTables []string `json:"deltaTables,omitempty" yaml:"deltaTables,omitempty"`
}

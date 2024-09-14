package kinesisanalyticsv2

type ApplicationSnapshot struct {
	// The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
	ApplicationName string `json:"applicationName,omitempty" yaml:"applicationName,omitempty"`

	// The name of the application snapshot.
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`
}

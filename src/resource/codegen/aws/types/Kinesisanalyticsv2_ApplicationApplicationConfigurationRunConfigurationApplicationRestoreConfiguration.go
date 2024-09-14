package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfiguration struct {
	// Specifies how the application should be restored. Valid values: `RESTORE_FROM_CUSTOM_SNAPSHOT`, `RESTORE_FROM_LATEST_SNAPSHOT`, `SKIP_RESTORE_FROM_SNAPSHOT`.
	ApplicationRestoreType string `json:"applicationRestoreType,omitempty" yaml:"applicationRestoreType,omitempty"`

	// The identifier of an existing snapshot of application state to use to restart an application. The application uses this value if `RESTORE_FROM_CUSTOM_SNAPSHOT` is specified for `application_restore_type`.
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`
}

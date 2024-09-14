package types

type Cloudrunv2_getJobTemplateTemplateVolume struct {
	// For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
	CloudSqlInstances []Cloudrunv2_getJobTemplateTemplateVolumeCloudSqlInstance `json:"cloudSqlInstances,omitempty" yaml:"cloudSqlInstances,omitempty"`

	// Ephemeral storage used as a shared volume.
	EmptyDirs []Cloudrunv2_getJobTemplateTemplateVolumeEmptyDir `json:"emptyDirs,omitempty" yaml:"emptyDirs,omitempty"`

	// The name of the Cloud Run v2 Job.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	Secrets []Cloudrunv2_getJobTemplateTemplateVolumeSecret `json:"secrets,omitempty" yaml:"secrets,omitempty"`
}

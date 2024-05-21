package types

type Cloudrunv2_getServiceTemplateVolume struct {
	// Ephemeral storage used as a shared volume.
	EmptyDirs []Cloudrunv2_getServiceTemplateVolumeEmptyDir `json:"emptyDirs,omitempty" yaml:"emptyDirs,omitempty"`

	// Represents a GCS Bucket mounted as a volume.
	Gcs []Cloudrunv2_getServiceTemplateVolumeGc `json:"gcs,omitempty" yaml:"gcs,omitempty"`

	// The name of the Cloud Run v2 Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Represents an NFS mount.
	Nfs []Cloudrunv2_getServiceTemplateVolumeNf `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	// Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	Secrets []Cloudrunv2_getServiceTemplateVolumeSecret `json:"secrets,omitempty" yaml:"secrets,omitempty"`

	// For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
	CloudSqlInstances []Cloudrunv2_getServiceTemplateVolumeCloudSqlInstance `json:"cloudSqlInstances,omitempty" yaml:"cloudSqlInstances,omitempty"`
}

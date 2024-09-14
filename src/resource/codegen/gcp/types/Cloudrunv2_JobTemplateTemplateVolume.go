package types

type Cloudrunv2_JobTemplateTemplateVolume struct {
	// Volume's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	   Structure is documented below.
	*/
	Secret Cloudrunv2_JobTemplateTemplateVolumeSecret `json:"secret,omitempty" yaml:"secret,omitempty"`

	/*
	   For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
	   Structure is documented below.
	*/
	CloudSqlInstance Cloudrunv2_JobTemplateTemplateVolumeCloudSqlInstance `json:"cloudSqlInstance,omitempty" yaml:"cloudSqlInstance,omitempty"`

	/*
	   Ephemeral storage used as a shared volume.
	   Structure is documented below.
	*/
	EmptyDir Cloudrunv2_JobTemplateTemplateVolumeEmptyDir `json:"emptyDir,omitempty" yaml:"emptyDir,omitempty"`
}

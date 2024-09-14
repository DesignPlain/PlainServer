package types

type Cloudrunv2_ServiceTemplateVolume struct {
	// Volume's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Represents an NFS mount.
	   Structure is documented below.
	*/
	Nfs Cloudrunv2_ServiceTemplateVolumeNfs `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	/*
	   Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	   Structure is documented below.
	*/
	Secret Cloudrunv2_ServiceTemplateVolumeSecret `json:"secret,omitempty" yaml:"secret,omitempty"`

	/*
	   For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
	   Structure is documented below.
	*/
	CloudSqlInstance Cloudrunv2_ServiceTemplateVolumeCloudSqlInstance `json:"cloudSqlInstance,omitempty" yaml:"cloudSqlInstance,omitempty"`

	/*
	   Ephemeral storage used as a shared volume.
	   Structure is documented below.
	*/
	EmptyDir Cloudrunv2_ServiceTemplateVolumeEmptyDir `json:"emptyDir,omitempty" yaml:"emptyDir,omitempty"`

	/*
	   Represents a GCS Bucket mounted as a volume.
	   Structure is documented below.
	*/
	Gcs Cloudrunv2_ServiceTemplateVolumeGcs `json:"gcs,omitempty" yaml:"gcs,omitempty"`
}

package fsx

type Backup struct {
	// The ID of the file system to back up. Required if backing up Lustre or Windows file systems.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the volume to back up. Required if backing up a ONTAP Volume.
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`
}

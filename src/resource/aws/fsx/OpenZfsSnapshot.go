package fsx

type OpenZfsSnapshot struct {
	// The name of the Snapshot. You can use a maximum of 203 alphanumeric characters plus either _ or -  or : or . for the name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the volume to snapshot. This can be the root volume or a child volume.
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`
}

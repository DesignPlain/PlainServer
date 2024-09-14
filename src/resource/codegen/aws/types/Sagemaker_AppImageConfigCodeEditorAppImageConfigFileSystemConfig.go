package types

type Sagemaker_AppImageConfigCodeEditorAppImageConfigFileSystemConfig struct {
	// The default POSIX group ID (GID). If not specified, defaults to `100`. Valid values are `0` and `100`.
	DefaultGid int `json:"defaultGid,omitempty" yaml:"defaultGid,omitempty"`

	// The default POSIX user ID (UID). If not specified, defaults to `1000`. Valid values are `0` and `1000`.
	DefaultUid int `json:"defaultUid,omitempty" yaml:"defaultUid,omitempty"`

	/*
	   The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to `/home/sagemaker-user`.

	   > --Note:-- When specifying `default_gid` and `default_uid`, Valid value pairs are [`0`, `0`] and [`100`, `1000`].
	*/
	MountPath string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`
}

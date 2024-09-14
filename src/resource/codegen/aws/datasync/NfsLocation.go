package datasync

import types "libds/aws/types"

type NfsLocation struct {
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block containing mount options used by DataSync to access the NFS Server.
	MountOptions types.Datasync_NfsLocationMountOptions `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`

	// Configuration block containing information for connecting to the NFS File System.
	OnPremConfig types.Datasync_NfsLocationOnPremConfig `json:"onPremConfig,omitempty" yaml:"onPremConfig,omitempty"`

	// Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
	ServerHostname string `json:"serverHostname,omitempty" yaml:"serverHostname,omitempty"`
}

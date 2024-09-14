package datasync

import types "libds/aws/types"

type LocationSmb struct {
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User string `json:"user,omitempty" yaml:"user,omitempty"`

	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `json:"agentArns,omitempty" yaml:"agentArns,omitempty"`

	// The name of the Windows domain the SMB server belongs to.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions types.Datasync_LocationSmbMountOptions `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`

	// The password of the user who can mount the share and has file permissions in the SMB.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname string `json:"serverHostname,omitempty" yaml:"serverHostname,omitempty"`

	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

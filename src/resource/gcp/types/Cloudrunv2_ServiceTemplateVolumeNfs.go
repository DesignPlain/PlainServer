package types

type Cloudrunv2_ServiceTemplateVolumeNfs struct {
	// Path that is exported by the NFS server.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   If true, mount the NFS volume as read only

	   - - -
	*/
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`

	// Hostname or IP address of the NFS server
	Server string `json:"server,omitempty" yaml:"server,omitempty"`
}

package types

type Datasync_NfsLocationMountOptions struct {
	// The specific NFS version that you want DataSync to use for mounting your NFS share. Valid values: `AUTOMATIC`, `NFS3`, `NFS4_0` and `NFS4_1`. Default: `AUTOMATIC`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

package types

type Datasync_LocationSmbMountOptions struct {
	// The specific SMB version that you want DataSync to use for mounting your SMB share. Valid values: `AUTOMATIC`, `SMB2`, and `SMB3`. Default: `AUTOMATIC`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

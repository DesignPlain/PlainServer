package types

type Networkmanager_LinkBandwidth struct {
	// Download speed in Mbps.
	DownloadSpeed int `json:"downloadSpeed,omitempty" yaml:"downloadSpeed,omitempty"`

	// Upload speed in Mbps.
	UploadSpeed int `json:"uploadSpeed,omitempty" yaml:"uploadSpeed,omitempty"`
}

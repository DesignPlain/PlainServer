package types

type Ecr_RepositoryImageScanningConfiguration struct {
	// Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
	ScanOnPush bool `json:"scanOnPush,omitempty" yaml:"scanOnPush,omitempty"`
}

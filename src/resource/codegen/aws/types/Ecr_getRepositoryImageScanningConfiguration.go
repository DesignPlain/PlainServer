package types

type Ecr_getRepositoryImageScanningConfiguration struct {
	// Whether images are scanned after being pushed to the repository.
	ScanOnPush bool `json:"scanOnPush,omitempty" yaml:"scanOnPush,omitempty"`
}

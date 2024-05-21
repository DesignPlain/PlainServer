package types

type Compute_RegionCommitmentLicenseResource struct {
	// The number of licenses purchased.
	Amount string `json:"amount,omitempty" yaml:"amount,omitempty"`

	// Specifies the core range of the instance for which this license applies.
	CoresPerLicense string `json:"coresPerLicense,omitempty" yaml:"coresPerLicense,omitempty"`

	// Any applicable license URI.
	License string `json:"license,omitempty" yaml:"license,omitempty"`
}

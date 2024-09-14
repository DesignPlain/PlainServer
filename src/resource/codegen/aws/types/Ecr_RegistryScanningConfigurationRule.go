package types

type Ecr_RegistryScanningConfigurationRule struct {
	// One or more repository filter blocks, containing a `filter` (required string filtering repositories, see pattern regex [here](https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ScanningRepositoryFilter.html)) and a `filter_type` (required string, currently only `WILDCARD` is supported).
	RepositoryFilters []Ecr_RegistryScanningConfigurationRuleRepositoryFilter `json:"repositoryFilters,omitempty" yaml:"repositoryFilters,omitempty"`

	// The frequency that scans are performed at for a private registry. Can be `SCAN_ON_PUSH`, `CONTINUOUS_SCAN`, or `MANUAL`.
	ScanFrequency string `json:"scanFrequency,omitempty" yaml:"scanFrequency,omitempty"`
}

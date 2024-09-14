package types

type Cloudfront_ContinuousDeploymentPolicyStagingDistributionDnsNames struct {
	// A list of CloudFront domain names for the staging distribution.
	Items []string `json:"items,omitempty" yaml:"items,omitempty"`

	// Number of CloudFront domain names in the staging distribution.
	Quantity int `json:"quantity,omitempty" yaml:"quantity,omitempty"`
}

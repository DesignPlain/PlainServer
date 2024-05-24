package cloudfront

import types "DesignSphere_Server/src/resource/aws/types"

type ContinuousDeploymentPolicy struct {
	// Whether this continuous deployment policy is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// CloudFront domain name of the staging distribution. See `staging_distribution_dns_names`.
	StagingDistributionDnsNames types.Cloudfront_ContinuousDeploymentPolicyStagingDistributionDnsNames `json:"stagingDistributionDnsNames,omitempty" yaml:"stagingDistributionDnsNames,omitempty"`

	// Parameters for routing production traffic from primary to staging distributions. See `traffic_config`.
	TrafficConfig types.Cloudfront_ContinuousDeploymentPolicyTrafficConfig `json:"trafficConfig,omitempty" yaml:"trafficConfig,omitempty"`
}

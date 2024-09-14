package lightsail

import types "libds/aws/types"

type Distribution struct {
	// Bundle ID to use for the distribution.
	BundleId string `json:"bundleId,omitempty" yaml:"bundleId,omitempty"`

	// A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
	CacheBehaviors []types.Lightsail_DistributionCacheBehavior `json:"cacheBehaviors,omitempty" yaml:"cacheBehaviors,omitempty"`

	// Object that describes the default cache behavior of the distribution. Detailed below
	DefaultCacheBehavior types.Lightsail_DistributionDefaultCacheBehavior `json:"defaultCacheBehavior,omitempty" yaml:"defaultCacheBehavior,omitempty"`

	// The IP address type of the distribution. Default: `dualstack`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// Name of the distribution.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
	Origin types.Lightsail_DistributionOrigin `json:"origin,omitempty" yaml:"origin,omitempty"`

	/*
	   An object that describes the cache behavior settings of the distribution. Detailed below

	   The following arguments are optional:
	*/
	CacheBehaviorSettings types.Lightsail_DistributionCacheBehaviorSettings `json:"cacheBehaviorSettings,omitempty" yaml:"cacheBehaviorSettings,omitempty"`

	// The name of the SSL/TLS certificate attached to the distribution, if any.
	CertificateName string `json:"certificateName,omitempty" yaml:"certificateName,omitempty"`

	// Indicates whether the distribution is enabled. Default: `true`.
	IsEnabled bool `json:"isEnabled,omitempty" yaml:"isEnabled,omitempty"`

	// Map of tags for the Lightsail Distribution. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

package types

type Cloudfront_DistributionTrustedKeyGroupItem struct {
	// ID of the key group that contains the public keys.
	KeyGroupId string `json:"keyGroupId,omitempty" yaml:"keyGroupId,omitempty"`

	// Set of active CloudFront key pairs associated with the signer account
	KeyPairIds []string `json:"keyPairIds,omitempty" yaml:"keyPairIds,omitempty"`
}

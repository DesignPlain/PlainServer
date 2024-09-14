package types

type Cloudfront_DistributionTrustedSignerItem struct {
	// AWS account ID or `self`
	AwsAccountNumber string `json:"awsAccountNumber,omitempty" yaml:"awsAccountNumber,omitempty"`

	// Set of active CloudFront key pairs associated with the signer account
	KeyPairIds []string `json:"keyPairIds,omitempty" yaml:"keyPairIds,omitempty"`
}

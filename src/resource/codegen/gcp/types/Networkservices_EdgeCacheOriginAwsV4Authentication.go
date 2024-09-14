package types

type Networkservices_EdgeCacheOriginAwsV4Authentication struct {
	// The name of the AWS region that your origin is in.
	OriginRegion string `json:"originRegion,omitempty" yaml:"originRegion,omitempty"`

	/*
	   The Secret Manager secret version of the secret access key used by your origin.
	   This is the resource name of the secret version in the format `projects/-/secrets/-/versions/-` where the `-` values are replaced by the project, secret, and version you require.
	*/
	SecretAccessKeyVersion string `json:"secretAccessKeyVersion,omitempty" yaml:"secretAccessKeyVersion,omitempty"`

	// The access key ID your origin uses to identify the key.
	AccessKeyId string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`
}

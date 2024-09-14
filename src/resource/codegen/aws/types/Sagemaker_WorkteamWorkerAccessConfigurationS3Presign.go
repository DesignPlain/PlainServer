package types

type Sagemaker_WorkteamWorkerAccessConfigurationS3Presign struct {
	// Use this parameter to specify the allowed request source. Possible sources are either SourceIp or VpcSourceIp. see IAM Policy Constraints details below.
	IamPolicyConstraints Sagemaker_WorkteamWorkerAccessConfigurationS3PresignIamPolicyConstraints `json:"iamPolicyConstraints,omitempty" yaml:"iamPolicyConstraints,omitempty"`
}

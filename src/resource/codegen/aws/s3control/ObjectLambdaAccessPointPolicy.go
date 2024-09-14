package s3control

type ObjectLambdaAccessPointPolicy struct {
	// The name of the Object Lambda Access Point.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Object Lambda Access Point resource policy document.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The AWS account ID for the account that owns the Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}

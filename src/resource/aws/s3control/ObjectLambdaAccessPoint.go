package s3control

import types "DesignSphere_Server/src/resource/aws/types"

type ObjectLambdaAccessPoint struct {
	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration types.S3control_ObjectLambdaAccessPointConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// The name for this Object Lambda Access Point.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

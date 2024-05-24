package cfg

import types "DesignSphere_Server/src/resource/aws/types"

type OrganizationConformancePack struct {
	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody string `json:"templateBody,omitempty" yaml:"templateBody,omitempty"`

	// Location of file, e.g., `s3://bucketname/prefix`, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3Uri string `json:"templateS3Uri,omitempty" yaml:"templateS3Uri,omitempty"`

	// Amazon S3 bucket where AWS Config stores conformance pack templates. Delivery bucket must begin with `awsconfigconforms` prefix. Maximum length of 63.
	DeliveryS3Bucket string `json:"deliveryS3Bucket,omitempty" yaml:"deliveryS3Bucket,omitempty"`

	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix string `json:"deliveryS3KeyPrefix,omitempty" yaml:"deliveryS3KeyPrefix,omitempty"`

	// Set of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack. Maximum of 1000 accounts.
	ExcludedAccounts []string `json:"excludedAccounts,omitempty" yaml:"excludedAccounts,omitempty"`

	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the `template_body` or in the template stored in Amazon S3 if using `template_s3_uri`.
	InputParameters []types.Cfg_OrganizationConformancePackInputParameter `json:"inputParameters,omitempty" yaml:"inputParameters,omitempty"`

	// The name of the organization conformance pack. Must begin with a letter and contain from 1 to 128 alphanumeric characters and hyphens.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

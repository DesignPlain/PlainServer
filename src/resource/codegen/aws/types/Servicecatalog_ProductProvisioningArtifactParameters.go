package types

type Servicecatalog_ProductProvisioningArtifactParameters struct {
	// Type of provisioning artifact. See [AWS Docs](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisioningArtifactProperties.html) for valid list of values.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	DisableTemplateValidation bool `json:"disableTemplateValidation,omitempty" yaml:"disableTemplateValidation,omitempty"`

	// Name of the provisioning artifact (for example, `v1`, `v2beta`). No spaces are allowed.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as `arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID]`.
	TemplatePhysicalId string `json:"templatePhysicalId,omitempty" yaml:"templatePhysicalId,omitempty"`

	// Template source as URL of the CloudFormation template in Amazon S3.
	TemplateUrl string `json:"templateUrl,omitempty" yaml:"templateUrl,omitempty"`
}

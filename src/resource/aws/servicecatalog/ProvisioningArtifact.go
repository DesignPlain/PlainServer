package servicecatalog

type ProvisioningArtifact struct {
	// Identifier of the product.
	ProductId string `json:"productId,omitempty" yaml:"productId,omitempty"`

	// Type of provisioning artifact. See [AWS Docs](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisioningArtifactProperties.html) for valid list of values.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). The default value is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`

	// Name of the provisioning artifact (for example, `v1`, `v2beta`). No spaces are allowed.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	DisableTemplateValidation bool `json:"disableTemplateValidation,omitempty" yaml:"disableTemplateValidation,omitempty"`

	// Information set by the administrator to provide guidance to end users about which provisioning artifacts to use. Valid values are `DEFAULT` and `DEPRECATED`. The default is `DEFAULT`. Users are able to make updates to a provisioned product of a deprecated version but cannot launch new provisioned products using a deprecated version.
	Guidance string `json:"guidance,omitempty" yaml:"guidance,omitempty"`

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as `arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID]`.
	TemplatePhysicalId string `json:"templatePhysicalId,omitempty" yaml:"templatePhysicalId,omitempty"`

	/*
	   Template source as URL of the CloudFormation template in Amazon S3.

	   The following arguments are optional:
	*/
	TemplateUrl string `json:"templateUrl,omitempty" yaml:"templateUrl,omitempty"`

	// Whether the product version is active. Inactive provisioning artifacts are invisible to end users. End users cannot launch or update a provisioned product from an inactive provisioning artifact. Default is `true`.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

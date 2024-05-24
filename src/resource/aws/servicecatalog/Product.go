package servicecatalog

import types "DesignSphere_Server/src/resource/aws/types"

type Product struct {
	// Description of the product.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the product.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Owner of the product.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`

	// Contact URL for product support.
	SupportUrl string `json:"supportUrl,omitempty" yaml:"supportUrl,omitempty"`

	// Tags to apply to the product. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`

	// Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
	ProvisioningArtifactParameters types.Servicecatalog_ProductProvisioningArtifactParameters `json:"provisioningArtifactParameters,omitempty" yaml:"provisioningArtifactParameters,omitempty"`

	// Support information about the product.
	SupportDescription string `json:"supportDescription,omitempty" yaml:"supportDescription,omitempty"`

	// Contact email for product support.
	SupportEmail string `json:"supportEmail,omitempty" yaml:"supportEmail,omitempty"`

	/*
	   Type of product. See [AWS Docs](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_CreateProduct.html#API_CreateProduct_RequestSyntax) for valid list of values.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Distributor (i.e., vendor) of the product.
	Distributor string `json:"distributor,omitempty" yaml:"distributor,omitempty"`
}

package types

type Sagemaker_ProjectServiceCatalogProvisioningDetails struct {
	// The path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path.
	PathId string `json:"pathId,omitempty" yaml:"pathId,omitempty"`

	// The ID of the product to provision.
	ProductId string `json:"productId,omitempty" yaml:"productId,omitempty"`

	// The ID of the provisioning artifact.
	ProvisioningArtifactId string `json:"provisioningArtifactId,omitempty" yaml:"provisioningArtifactId,omitempty"`

	// A list of key value pairs that you specify when you provision a product. See Provisioning Parameter below.
	ProvisioningParameters []Sagemaker_ProjectServiceCatalogProvisioningDetailsProvisioningParameter `json:"provisioningParameters,omitempty" yaml:"provisioningParameters,omitempty"`
}

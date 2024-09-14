package sagemaker

type ServicecatalogPortfolioStatus struct {
	// Whether Service Catalog is enabled or disabled in SageMaker. Valid values are `Enabled` and `Disabled`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}

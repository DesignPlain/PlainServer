package types

type Gkebackup_RestorePlanRestoreConfigTransformationRuleResourceFilterGroupKind struct {
	/*
	   API Group string of a Kubernetes resource, e.g.
	   "apiextensions.k8s.io", "storage.k8s.io", etc.
	   Use empty string for core group.
	*/
	ResourceGroup string `json:"resourceGroup,omitempty" yaml:"resourceGroup,omitempty"`

	/*
	   Kind of a Kubernetes resource, e.g.
	   "CustomResourceDefinition", "StorageClass", etc.
	*/
	ResourceKind string `json:"resourceKind,omitempty" yaml:"resourceKind,omitempty"`
}

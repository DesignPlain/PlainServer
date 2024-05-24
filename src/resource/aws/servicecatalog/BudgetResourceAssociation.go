package servicecatalog

type BudgetResourceAssociation struct {
	// Budget name.
	BudgetName string `json:"budgetName,omitempty" yaml:"budgetName,omitempty"`

	// Resource identifier.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`
}

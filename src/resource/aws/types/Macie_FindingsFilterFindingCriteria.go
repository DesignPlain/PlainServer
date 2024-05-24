package types

type Macie_FindingsFilterFindingCriteria struct {
	// A condition that specifies the property, operator, and one or more values to use to filter the results.  (documented below)
	Criterions []Macie_FindingsFilterFindingCriteriaCriterion `json:"criterions,omitempty" yaml:"criterions,omitempty"`
}

package types

type Lakeformation_PermissionsLfTagPolicy struct {
	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	/*
	   A list of tag conditions that apply to the resource's tag policy. Configuration block for tag conditions that apply to the policy. See `expression` below.

	   The following argument is optional:
	*/
	Expressions []Lakeformation_PermissionsLfTagPolicyExpression `json:"expressions,omitempty" yaml:"expressions,omitempty"`

	// The resource type for which the tag policy applies. Valid values are `DATABASE` and `TABLE`.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}

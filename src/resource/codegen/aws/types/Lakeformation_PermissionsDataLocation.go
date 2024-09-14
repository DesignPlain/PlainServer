package types

type Lakeformation_PermissionsDataLocation struct {
	/*
	   Amazon Resource Name (ARN) that uniquely identifies the data location resource.

	   The following argument is optional:
	*/
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Identifier for the Data Catalog where the location is registered with Lake Formation. By default, it is the account ID of the caller.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`
}

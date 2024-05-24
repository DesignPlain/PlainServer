package inspector2

type Enabler struct {
	/*
	   Set of account IDs.
	   Can contain one of: the Organization's Administrator Account, or one or more Member Accounts.
	*/
	AccountIds []string `json:"accountIds,omitempty" yaml:"accountIds,omitempty"`

	/*
	   Type of resources to scan.
	   Valid values are `EC2`, `ECR`, `LAMBDA` and `LAMBDA_CODE`.
	   At least one item is required.
	*/
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`
}

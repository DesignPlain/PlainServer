package types

type Bigquery_ConnectionAws struct {
	/*
	   Authentication using Google owned service account to assume into customer's AWS IAM Role.
	   Structure is documented below.
	*/
	AccessRole Bigquery_ConnectionAwsAccessRole `json:"accessRole,omitempty" yaml:"accessRole,omitempty"`
}

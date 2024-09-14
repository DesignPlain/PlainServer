package types

type Bigquery_ConnectionAwsAccessRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.
	IamRoleId string `json:"iamRoleId,omitempty" yaml:"iamRoleId,omitempty"`

	/*
	   (Output)
	   A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's AWS IAM Role.
	*/
	Identity string `json:"identity,omitempty" yaml:"identity,omitempty"`
}

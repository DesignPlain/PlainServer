package types

type Bigquery_ConnectionCloudSqlCredential struct {
	/*
	   Password for database.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Username for database.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

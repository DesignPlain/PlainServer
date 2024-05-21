package types

type Bigquery_AppProfileStandardIsolation struct {
	/*
	   The priority of requests sent using this app profile.
	   Possible values are: `PRIORITY_LOW`, `PRIORITY_MEDIUM`, `PRIORITY_HIGH`.
	*/
	Priority string `json:"priority,omitempty" yaml:"priority,omitempty"`
}

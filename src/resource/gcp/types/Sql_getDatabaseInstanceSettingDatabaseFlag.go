package types

type Sql_getDatabaseInstanceSettingDatabaseFlag struct {
	// The name of the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the flag.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

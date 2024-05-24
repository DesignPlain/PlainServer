package types

type Backup_getFrameworkControlInputParameter struct {
	// Backup framework name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of parameter, for example, hourly.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

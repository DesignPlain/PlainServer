package types

type Backup_FrameworkControlInputParameter struct {
	// The name of a parameter, for example, BackupPlanFrequency.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of parameter, for example, hourly.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

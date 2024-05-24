package types

type Backup_getFrameworkControl struct {
	// One or more input parameter blocks. An example of a control with two parameters is: "backup plan frequency is at least daily and the retention period is at least 1 year". The first parameter is daily. The second parameter is 1 year. Detailed below.
	InputParameters []Backup_getFrameworkControlInputParameter `json:"inputParameters,omitempty" yaml:"inputParameters,omitempty"`

	// Backup framework name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. Detailed below.
	Scopes []Backup_getFrameworkControlScope `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}

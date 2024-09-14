package types

type Backup_FrameworkControl struct {
	// One or more input parameter blocks. An example of a control with two parameters is: "backup plan frequency is at least daily and the retention period is at least 1 year". The first parameter is daily. The second parameter is 1 year. Detailed below.
	InputParameters []Backup_FrameworkControlInputParameter `json:"inputParameters,omitempty" yaml:"inputParameters,omitempty"`

	// The name of a control. This name is between 1 and 256 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. Detailed below.
	Scope Backup_FrameworkControlScope `json:"scope,omitempty" yaml:"scope,omitempty"`
}

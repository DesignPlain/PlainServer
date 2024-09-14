package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceExecEnforce struct {
	/*
	   Optional arguments to pass to the source during
	   execution.
	*/
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	/*
	   A remote or local file. Structure is
	   documented below.
	*/
	File Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceExecEnforceFile `json:"file,omitempty" yaml:"file,omitempty"`

	/*
	   The script interpreter to use. Possible values
	   are: `INTERPRETER_UNSPECIFIED`, `NONE`, `SHELL`, `POWERSHELL`.
	*/
	Interpreter string `json:"interpreter,omitempty" yaml:"interpreter,omitempty"`

	/*
	   Only recorded for enforce Exec. Path to an
	   output file (that is created by this Exec) whose content will be recorded in
	   OSPolicyResourceCompliance after a successful run. Absence or failure to
	   read this file will result in this ExecResource being non-compliant. Output
	   file size is limited to 100K bytes.
	*/
	OutputFilePath string `json:"outputFilePath,omitempty" yaml:"outputFilePath,omitempty"`

	/*
	   An inline script. The size of the script is limited to
	   1024 characters.
	*/
	Script string `json:"script,omitempty" yaml:"script,omitempty"`
}

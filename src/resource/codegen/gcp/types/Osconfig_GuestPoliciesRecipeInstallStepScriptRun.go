package types

type Osconfig_GuestPoliciesRecipeInstallStepScriptRun struct {
	// Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0]
	AllowedExitCodes []int `json:"allowedExitCodes,omitempty" yaml:"allowedExitCodes,omitempty"`

	/*
	   The script interpreter to use to run the script. If no interpreter is specified the script is executed directly,
	   which likely only succeed for scripts with shebang lines.
	   Possible values are: `SHELL`, `POWERSHELL`.
	*/
	Interpreter string `json:"interpreter,omitempty" yaml:"interpreter,omitempty"`

	// The shell script to be executed.
	Script string `json:"script,omitempty" yaml:"script,omitempty"`
}

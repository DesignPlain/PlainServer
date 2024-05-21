package types

type Osconfig_PatchDeploymentPatchConfigPreStepWindowsExecStepConfig struct {
	/*
	   A Cloud Storage object containing the executable.
	   Structure is documented below.
	*/
	GcsObject Osconfig_PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject `json:"gcsObject,omitempty" yaml:"gcsObject,omitempty"`

	/*
	   The script interpreter to use to run the script. If no interpreter is specified the script will
	   be executed directly, which will likely only succeed for scripts with shebang lines.
	   Possible values are: `SHELL`, `POWERSHELL`.
	*/
	Interpreter string `json:"interpreter,omitempty" yaml:"interpreter,omitempty"`

	// An absolute path to the executable on the VM.
	LocalPath string `json:"localPath,omitempty" yaml:"localPath,omitempty"`

	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	AllowedSuccessCodes []int `json:"allowedSuccessCodes,omitempty" yaml:"allowedSuccessCodes,omitempty"`
}

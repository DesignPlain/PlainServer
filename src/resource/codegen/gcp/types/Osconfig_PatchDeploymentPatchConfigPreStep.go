package types

type Osconfig_PatchDeploymentPatchConfigPreStep struct {
	/*
	   The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	   Structure is documented below.
	*/
	LinuxExecStepConfig Osconfig_PatchDeploymentPatchConfigPreStepLinuxExecStepConfig `json:"linuxExecStepConfig,omitempty" yaml:"linuxExecStepConfig,omitempty"`

	/*
	   The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	   Structure is documented below.
	*/
	WindowsExecStepConfig Osconfig_PatchDeploymentPatchConfigPreStepWindowsExecStepConfig `json:"windowsExecStepConfig,omitempty" yaml:"windowsExecStepConfig,omitempty"`
}

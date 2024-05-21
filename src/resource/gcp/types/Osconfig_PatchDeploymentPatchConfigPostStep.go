package types

type Osconfig_PatchDeploymentPatchConfigPostStep struct {
	/*
	   The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	   Structure is documented below.
	*/
	LinuxExecStepConfig Osconfig_PatchDeploymentPatchConfigPostStepLinuxExecStepConfig `json:"linuxExecStepConfig,omitempty" yaml:"linuxExecStepConfig,omitempty"`

	/*
	   The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	   Structure is documented below.
	*/
	WindowsExecStepConfig Osconfig_PatchDeploymentPatchConfigPostStepWindowsExecStepConfig `json:"windowsExecStepConfig,omitempty" yaml:"windowsExecStepConfig,omitempty"`
}

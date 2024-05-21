package types

type Osconfig_PatchDeploymentPatchConfig struct {
	/*
	   The ExecStep to run after the patch update.
	   Structure is documented below.
	*/
	PostStep Osconfig_PatchDeploymentPatchConfigPostStep `json:"postStep,omitempty" yaml:"postStep,omitempty"`

	/*
	   The ExecStep to run before the patch update.
	   Structure is documented below.
	*/
	PreStep Osconfig_PatchDeploymentPatchConfigPreStep `json:"preStep,omitempty" yaml:"preStep,omitempty"`

	/*
	   Yum update settings. Use this setting to override the default yum patch rules.
	   Structure is documented below.
	*/
	Yum Osconfig_PatchDeploymentPatchConfigYum `json:"yum,omitempty" yaml:"yum,omitempty"`

	/*
	   zypper update settings. Use this setting to override the default zypper patch rules.
	   Structure is documented below.
	*/
	Zypper Osconfig_PatchDeploymentPatchConfigZypper `json:"zypper,omitempty" yaml:"zypper,omitempty"`

	/*
	   Apt update settings. Use this setting to override the default apt patch rules.
	   Structure is documented below.
	*/
	Apt Osconfig_PatchDeploymentPatchConfigApt `json:"apt,omitempty" yaml:"apt,omitempty"`

	/*
	   goo update settings. Use this setting to override the default goo patch rules.
	   Structure is documented below.
	*/
	Goo Osconfig_PatchDeploymentPatchConfigGoo `json:"goo,omitempty" yaml:"goo,omitempty"`

	// Allows the patch job to run on Managed instance groups (MIGs).
	MigInstancesAllowed bool `json:"migInstancesAllowed,omitempty" yaml:"migInstancesAllowed,omitempty"`

	/*
	   Post-patch reboot settings.
	   Possible values are: `DEFAULT`, `ALWAYS`, `NEVER`.
	*/
	RebootConfig string `json:"rebootConfig,omitempty" yaml:"rebootConfig,omitempty"`

	/*
	   Windows update settings. Use this setting to override the default Windows patch rules.
	   Structure is documented below.
	*/
	WindowsUpdate Osconfig_PatchDeploymentPatchConfigWindowsUpdate `json:"windowsUpdate,omitempty" yaml:"windowsUpdate,omitempty"`
}

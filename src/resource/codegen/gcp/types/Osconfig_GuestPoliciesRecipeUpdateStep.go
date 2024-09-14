package types

type Osconfig_GuestPoliciesRecipeUpdateStep struct {
	/*
	   Copies a file onto the instance.
	   Structure is documented below.
	*/
	FileCopy Osconfig_GuestPoliciesRecipeUpdateStepFileCopy `json:"fileCopy,omitempty" yaml:"fileCopy,omitempty"`

	/*
	   Executes an artifact or local file.
	   Structure is documented below.
	*/
	FileExec Osconfig_GuestPoliciesRecipeUpdateStepFileExec `json:"fileExec,omitempty" yaml:"fileExec,omitempty"`

	/*
	   Installs an MSI file.
	   Structure is documented below.
	*/
	MsiInstallation Osconfig_GuestPoliciesRecipeUpdateStepMsiInstallation `json:"msiInstallation,omitempty" yaml:"msiInstallation,omitempty"`

	/*
	   Installs an rpm file via the rpm utility.
	   Structure is documented below.
	*/
	RpmInstallation Osconfig_GuestPoliciesRecipeUpdateStepRpmInstallation `json:"rpmInstallation,omitempty" yaml:"rpmInstallation,omitempty"`

	/*
	   Runs commands in a shell.
	   Structure is documented below.
	*/
	ScriptRun Osconfig_GuestPoliciesRecipeUpdateStepScriptRun `json:"scriptRun,omitempty" yaml:"scriptRun,omitempty"`

	/*
	   Extracts an archive into the specified directory.
	   Structure is documented below.
	*/
	ArchiveExtraction Osconfig_GuestPoliciesRecipeUpdateStepArchiveExtraction `json:"archiveExtraction,omitempty" yaml:"archiveExtraction,omitempty"`

	/*
	   Installs a deb file via dpkg.
	   Structure is documented below.
	*/
	DpkgInstallation Osconfig_GuestPoliciesRecipeUpdateStepDpkgInstallation `json:"dpkgInstallation,omitempty" yaml:"dpkgInstallation,omitempty"`
}

package types

type Osconfig_GuestPoliciesRecipeInstallStep struct {
	/*
	   Installs an MSI file.
	   Structure is documented below.
	*/
	MsiInstallation Osconfig_GuestPoliciesRecipeInstallStepMsiInstallation `json:"msiInstallation,omitempty" yaml:"msiInstallation,omitempty"`

	/*
	   Installs an rpm file via the rpm utility.
	   Structure is documented below.
	*/
	RpmInstallation Osconfig_GuestPoliciesRecipeInstallStepRpmInstallation `json:"rpmInstallation,omitempty" yaml:"rpmInstallation,omitempty"`

	/*
	   Runs commands in a shell.
	   Structure is documented below.
	*/
	ScriptRun Osconfig_GuestPoliciesRecipeInstallStepScriptRun `json:"scriptRun,omitempty" yaml:"scriptRun,omitempty"`

	/*
	   Extracts an archive into the specified directory.
	   Structure is documented below.
	*/
	ArchiveExtraction Osconfig_GuestPoliciesRecipeInstallStepArchiveExtraction `json:"archiveExtraction,omitempty" yaml:"archiveExtraction,omitempty"`

	/*
	   Installs a deb file via dpkg.
	   Structure is documented below.
	*/
	DpkgInstallation Osconfig_GuestPoliciesRecipeInstallStepDpkgInstallation `json:"dpkgInstallation,omitempty" yaml:"dpkgInstallation,omitempty"`

	/*
	   Copies a file onto the instance.
	   Structure is documented below.
	*/
	FileCopy Osconfig_GuestPoliciesRecipeInstallStepFileCopy `json:"fileCopy,omitempty" yaml:"fileCopy,omitempty"`

	/*
	   Executes an artifact or local file.
	   Structure is documented below.
	*/
	FileExec Osconfig_GuestPoliciesRecipeInstallStepFileExec `json:"fileExec,omitempty" yaml:"fileExec,omitempty"`
}

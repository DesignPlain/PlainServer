package types

type Deploymentmanager_DeploymentTarget struct {
	/*
	   The root configuration file to use for this deployment.
	   Structure is documented below.
	*/
	Config Deploymentmanager_DeploymentTargetConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   Specifies import files for this configuration. This can be
	   used to import templates or other files. For example, you might
	   import a text file in order to use the file in a template.
	   Structure is documented below.
	*/
	Imports []Deploymentmanager_DeploymentTargetImport `json:"imports,omitempty" yaml:"imports,omitempty"`
}

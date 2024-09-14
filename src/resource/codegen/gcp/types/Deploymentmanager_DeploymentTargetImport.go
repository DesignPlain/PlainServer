package types

type Deploymentmanager_DeploymentTargetImport struct {
	// The full contents of the template that you want to import.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	/*
	   The name of the template to import, as declared in the YAML
	   configuration.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

package types

type Cloudrun_ServiceTemplateSpecContainerEnvValueFromSecretKeyRef struct {
	/*
	   A Cloud Secret Manager secret version. Must be 'latest' for the latest
	   version or an integer for a specific version.
	*/
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   The name of the secret in Cloud Secret Manager. By default, the secret is assumed to be in the same project.
	   If the secret is in another project, you must define an alias.
	   An alias definition has the form: :projects/{project-id|project-number}/secrets/.
	   If multiple alias definitions are needed, they must be separated by commas.
	   The alias definitions must be set on the run.googleapis.com/secrets annotation.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

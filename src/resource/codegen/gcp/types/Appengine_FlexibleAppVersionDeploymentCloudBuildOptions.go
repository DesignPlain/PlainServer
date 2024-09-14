package types

type Appengine_FlexibleAppVersionDeploymentCloudBuildOptions struct {
	// Path to the yaml file used in deployment, used to determine runtime configuration details.
	AppYamlPath string `json:"appYamlPath,omitempty" yaml:"appYamlPath,omitempty"`

	/*
	   The Cloud Build timeout used as part of any dependent builds performed by version creation. Defaults to 10 minutes.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	CloudBuildTimeout string `json:"cloudBuildTimeout,omitempty" yaml:"cloudBuildTimeout,omitempty"`
}

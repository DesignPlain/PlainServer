package amplify

type BackendEnvironment struct {
	// AWS CloudFormation stack name of a backend environment.
	StackName string `json:"stackName,omitempty" yaml:"stackName,omitempty"`

	// Unique ID for an Amplify app.
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	// Name of deployment artifacts.
	DeploymentArtifacts string `json:"deploymentArtifacts,omitempty" yaml:"deploymentArtifacts,omitempty"`

	// Name for the backend environment.
	EnvironmentName string `json:"environmentName,omitempty" yaml:"environmentName,omitempty"`
}

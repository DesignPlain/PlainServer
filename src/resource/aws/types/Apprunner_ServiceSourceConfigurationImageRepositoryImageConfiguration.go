package types

type Apprunner_ServiceSourceConfigurationImageRepositoryImageConfiguration struct {
	// Port that your application listens to in the container. Defaults to `"8080"`.
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	// Secrets and parameters available to your service as environment variables. A map of key/value pairs, where the key is the desired name of the Secret in the environment (i.e. it does not have to match the name of the secret in Secrets Manager or SSM Parameter Store), and the value is the ARN of the secret from AWS Secrets Manager or the ARN of the parameter in AWS SSM Parameter Store.
	RuntimeEnvironmentSecrets map[string]string `json:"runtimeEnvironmentSecrets,omitempty" yaml:"runtimeEnvironmentSecrets,omitempty"`

	// Environment variables available to your running App Runner service. A map of key/value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables map[string]string `json:"runtimeEnvironmentVariables,omitempty" yaml:"runtimeEnvironmentVariables,omitempty"`

	// Command App Runner runs to start the application in the source image. If specified, this command overrides the Docker image’s default start command.
	StartCommand string `json:"startCommand,omitempty" yaml:"startCommand,omitempty"`
}

package types

type Apprunner_ServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues struct {
	// Command App Runner runs to start your application.
	StartCommand string `json:"startCommand,omitempty" yaml:"startCommand,omitempty"`

	// Command App Runner runs to build your application.
	BuildCommand string `json:"buildCommand,omitempty" yaml:"buildCommand,omitempty"`

	// Port that your application listens to in the container. Defaults to `"8080"`.
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	// Runtime environment type for building and running an App Runner service. Represents a programming language runtime. Valid values: `PYTHON_3`, `NODEJS_12`, `NODEJS_14`, `NODEJS_16`, `CORRETTO_8`, `CORRETTO_11`, `GO_1`, `DOTNET_6`, `PHP_81`, `RUBY_31`.
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	// Secrets and parameters available to your service as environment variables. A map of key/value pairs, where the key is the desired name of the Secret in the environment (i.e. it does not have to match the name of the secret in Secrets Manager or SSM Parameter Store), and the value is the ARN of the secret from AWS Secrets Manager or the ARN of the parameter in AWS SSM Parameter Store.
	RuntimeEnvironmentSecrets map[string]string `json:"runtimeEnvironmentSecrets,omitempty" yaml:"runtimeEnvironmentSecrets,omitempty"`

	// Environment variables available to your running App Runner service. A map of key/value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables map[string]string `json:"runtimeEnvironmentVariables,omitempty" yaml:"runtimeEnvironmentVariables,omitempty"`
}

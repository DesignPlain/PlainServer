package types

type Cloudbuild_TriggerBuildAvailableSecretsSecretManager struct {
	/*
	   Environment variable name to associate with the secret. Secret environment
	   variables must be unique across all of a build's secrets, and must be used
	   by at least one build step.
	*/
	Env string `json:"env,omitempty" yaml:"env,omitempty"`

	// Resource name of the SecretVersion. In format: projects/-/secrets/-/versions/-
	VersionName string `json:"versionName,omitempty" yaml:"versionName,omitempty"`
}

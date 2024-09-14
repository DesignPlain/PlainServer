package types

type Cloudbuild_TriggerBuildAvailableSecrets struct {
	/*
	   Pairs a secret environment variable with a SecretVersion in Secret Manager.
	   Structure is documented below.
	*/
	SecretManagers []Cloudbuild_TriggerBuildAvailableSecretsSecretManager `json:"secretManagers,omitempty" yaml:"secretManagers,omitempty"`
}

package types

type Cloudbuild_getTriggerBuildAvailableSecret struct {
	// Pairs a secret environment variable with a SecretVersion in Secret Manager.
	SecretManagers []Cloudbuild_getTriggerBuildAvailableSecretSecretManager `json:"secretManagers,omitempty" yaml:"secretManagers,omitempty"`
}

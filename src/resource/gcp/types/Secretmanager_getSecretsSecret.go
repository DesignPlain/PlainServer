package types

type Secretmanager_getSecretsSecret struct {
	// This must be unique within the project.
	SecretId string `json:"secretId,omitempty" yaml:"secretId,omitempty"`

	/*
	   A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
	   Structure is documented below.
	*/
	Topics []Secretmanager_getSecretsSecretTopic `json:"topics,omitempty" yaml:"topics,omitempty"`

	/*
	   The TTL for the Secret.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	   Only one of 'ttl' or 'expire_time' can be provided.
	*/
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	//
	EffectiveLabels map[string]string `json:"effectiveLabels,omitempty" yaml:"effectiveLabels,omitempty"`

	// The resource name of the Pub/Sub topic that will be published to.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	EffectiveAnnotations map[string]string `json:"effectiveAnnotations,omitempty" yaml:"effectiveAnnotations,omitempty"`

	// Custom metadata about the secret.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// The time at which the Secret was created.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// The ID of the project.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The combination of labels configured directly on the resource
	    and default labels configured on the provider.
	*/
	PulumiLabels map[string]string `json:"pulumiLabels,omitempty" yaml:"pulumiLabels,omitempty"`

	/*
	   The replication policy of the secret data attached to the Secret.
	   Structure is documented below.
	*/
	Replications []Secretmanager_getSecretsSecretReplication `json:"replications,omitempty" yaml:"replications,omitempty"`

	/*
	   The rotation time and period for a Secret.
	   Structure is documented below.
	*/
	Rotations []Secretmanager_getSecretsSecretRotation `json:"rotations,omitempty" yaml:"rotations,omitempty"`

	// Mapping from version alias to version name.
	VersionAliases map[string]string `json:"versionAliases,omitempty" yaml:"versionAliases,omitempty"`

	// Timestamp in UTC when the Secret is scheduled to expire.
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`

	// The labels assigned to this Secret.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

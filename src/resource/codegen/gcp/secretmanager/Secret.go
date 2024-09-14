package secretmanager

import types "libds/gcp/types"

type Secret struct {
	// This must be unique within the project.
	SecretId string `json:"secretId,omitempty" yaml:"secretId,omitempty"`

	/*
	   The TTL for the Secret.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	   Only one of `ttl` or `expire_time` can be provided.
	*/
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	/*
	   Custom metadata about the secret.
	   Annotations are distinct from various forms of labels. Annotations exist to allow
	   client tools to store their own state information without requiring a database.
	   Annotation keys must be between 1 and 63 characters long, have a UTF-8 encoding of
	   maximum 128 bytes, begin and end with an alphanumeric character ([a-z0-9A-Z]), and
	   may have dashes (-), underscores (_), dots (.), and alphanumerics in between these
	   symbols.
	   The total size of annotation keys and values must be less than 16KiB.
	   An object containing a list of "key": value pairs. Example:
	   { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	   Only one of `expire_time` or `ttl` can be provided.
	*/
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`

	/*
	   The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
	   Structure is documented below.
	*/
	Rotation types.Secretmanager_SecretRotation `json:"rotation,omitempty" yaml:"rotation,omitempty"`

	/*
	   A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
	   Structure is documented below.
	*/
	Topics []types.Secretmanager_SecretTopic `json:"topics,omitempty" yaml:"topics,omitempty"`

	/*
	   Mapping from version alias to version name.
	   A version alias is a string with a maximum length of 63 characters and can contain
	   uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_')
	   characters. An alias string must start with a letter and cannot be the string
	   'latest' or 'NEW'. No more than 50 aliases can be assigned to a given secret.
	   An object containing a list of "key": value pairs. Example:
	   { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	*/
	VersionAliases map[string]string `json:"versionAliases,omitempty" yaml:"versionAliases,omitempty"`

	/*
	   The labels assigned to this Secret.
	   Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	   and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	   Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	   and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	   No more than 64 labels can be assigned to a given resource.
	   An object containing a list of "key": value pairs. Example:
	   { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The replication policy of the secret data attached to the Secret. It cannot be changed
	   after the Secret has been created.
	   Structure is documented below.
	*/
	Replication types.Secretmanager_SecretReplication `json:"replication,omitempty" yaml:"replication,omitempty"`
}

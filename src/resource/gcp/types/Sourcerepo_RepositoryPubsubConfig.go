package types

type Sourcerepo_RepositoryPubsubConfig struct {
	/*
	   The format of the Cloud Pub/Sub messages.
	   - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
	   - JSON: The message payload is a JSON string of SourceRepoEvent.
	   Possible values are: `PROTOBUF`, `JSON`.
	*/
	MessageFormat string `json:"messageFormat,omitempty" yaml:"messageFormat,omitempty"`

	/*
	   Email address of the service account used for publishing Cloud Pub/Sub messages.
	   This service account needs to be in the same project as the PubsubConfig. When added,
	   the caller needs to have iam.serviceAccounts.actAs permission on this service account.
	   If unspecified, it defaults to the compute engine default service account.
	*/
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	// The identifier for this object. Format specified above.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
}

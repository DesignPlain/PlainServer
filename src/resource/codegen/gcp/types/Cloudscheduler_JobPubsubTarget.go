package types

type Cloudscheduler_JobPubsubTarget struct {
	/*
	   Attributes for PubsubMessage.
	   Pubsub message must contain either non-empty data, or at least one attribute.
	*/
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	/*
	   The message payload for PubsubMessage.
	   Pubsub message must contain either non-empty data, or at least one attribute.
	   A base64-encoded string.
	*/
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	/*
	   The full resource name for the Cloud Pub/Sub topic to which
	   messages will be published when a job is delivered. ~>--NOTE:--
	   The topic name must be in the same format as required by PubSub's
	   PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
	*/
	TopicName string `json:"topicName,omitempty" yaml:"topicName,omitempty"`
}

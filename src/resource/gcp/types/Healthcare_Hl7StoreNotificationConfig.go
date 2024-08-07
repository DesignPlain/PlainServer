package types

type Healthcare_Hl7StoreNotificationConfig struct {
	/*
	   The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
	   PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
	   It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
	   was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
	   project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
	   Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
	*/
	PubsubTopic string `json:"pubsubTopic,omitempty" yaml:"pubsubTopic,omitempty"`
}

package types

type Sesv2_ContactListTopic struct {
	/*
	   Name of the topic.

	   The following arguments are optional:
	*/
	TopicName string `json:"topicName,omitempty" yaml:"topicName,omitempty"`

	// Default subscription status to be applied to a contact if the contact has not noted their preference for subscribing to a topic.
	DefaultSubscriptionStatus string `json:"defaultSubscriptionStatus,omitempty" yaml:"defaultSubscriptionStatus,omitempty"`

	// Description of what the topic is about, which the contact will see.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the topic the contact will see.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}

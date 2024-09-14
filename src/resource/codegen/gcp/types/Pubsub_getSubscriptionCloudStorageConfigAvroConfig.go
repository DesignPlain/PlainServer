package types

type Pubsub_getSubscriptionCloudStorageConfigAvroConfig struct {
	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey as additional fields in the output.
	WriteMetadata bool `json:"writeMetadata,omitempty" yaml:"writeMetadata,omitempty"`
}

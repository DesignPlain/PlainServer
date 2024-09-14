package types

type Pubsub_LiteSubscriptionDeliveryConfig struct {
	/*
	   When this subscription should send messages to subscribers relative to messages persistence in storage.
	   Possible values are: `DELIVER_IMMEDIATELY`, `DELIVER_AFTER_STORED`, `DELIVERY_REQUIREMENT_UNSPECIFIED`.
	*/
	DeliveryRequirement string `json:"deliveryRequirement,omitempty" yaml:"deliveryRequirement,omitempty"`
}

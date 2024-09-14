package types

type Costexplorer_AnomalySubscriptionSubscriber struct {
	// The address of the subscriber. If type is `SNS`, this will be the arn of the sns topic. If type is `EMAIL`, this will be the destination email address.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The type of subscription. Valid Values: `SNS` | `EMAIL`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

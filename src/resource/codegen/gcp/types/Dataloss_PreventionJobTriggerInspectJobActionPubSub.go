package types

type Dataloss_PreventionJobTriggerInspectJobActionPubSub struct {
	// Cloud Pub/Sub topic to send notifications to.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
}

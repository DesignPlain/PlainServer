package types

type Eventarc_TriggerTransport struct {
	// The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.
	Pubsub Eventarc_TriggerTransportPubsub `json:"pubsub,omitempty" yaml:"pubsub,omitempty"`
}

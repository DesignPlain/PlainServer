package types

type Cloudasset_ProjectFeedFeedOutputConfig struct {
	/*
	   Destination on Cloud Pubsub.
	   Structure is documented below.
	*/
	PubsubDestination Cloudasset_ProjectFeedFeedOutputConfigPubsubDestination `json:"pubsubDestination,omitempty" yaml:"pubsubDestination,omitempty"`
}

package types

type Cloudasset_OrganizationFeedFeedOutputConfig struct {
	/*
	   Destination on Cloud Pubsub.
	   Structure is documented below.
	*/
	PubsubDestination Cloudasset_OrganizationFeedFeedOutputConfigPubsubDestination `json:"pubsubDestination,omitempty" yaml:"pubsubDestination,omitempty"`
}

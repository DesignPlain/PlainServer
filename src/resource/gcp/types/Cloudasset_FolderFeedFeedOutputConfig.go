package types

type Cloudasset_FolderFeedFeedOutputConfig struct {
	/*
	   Destination on Cloud Pubsub.
	   Structure is documented below.
	*/
	PubsubDestination Cloudasset_FolderFeedFeedOutputConfigPubsubDestination `json:"pubsubDestination,omitempty" yaml:"pubsubDestination,omitempty"`
}

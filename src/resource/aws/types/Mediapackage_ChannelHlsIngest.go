package types

type Mediapackage_ChannelHlsIngest struct {
	// A list of the ingest endpoints
	IngestEndpoints []Mediapackage_ChannelHlsIngestIngestEndpoint `json:"ingestEndpoints,omitempty" yaml:"ingestEndpoints,omitempty"`
}

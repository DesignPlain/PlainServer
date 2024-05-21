package types

type Bigqueryanalyticshub_ListingPublisher struct {
	// Name of the listing publisher.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Email or URL of the listing publisher.
	PrimaryContact string `json:"primaryContact,omitempty" yaml:"primaryContact,omitempty"`
}

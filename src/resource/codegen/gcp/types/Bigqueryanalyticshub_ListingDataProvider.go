package types

type Bigqueryanalyticshub_ListingDataProvider struct {
	// Email or URL of the data provider.
	PrimaryContact string `json:"primaryContact,omitempty" yaml:"primaryContact,omitempty"`

	// Name of the data provider.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

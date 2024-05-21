package types

type Bigqueryanalyticshub_ListingDataProvider struct {
	// Name of the data provider.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Email or URL of the data provider.
	PrimaryContact string `json:"primaryContact,omitempty" yaml:"primaryContact,omitempty"`
}

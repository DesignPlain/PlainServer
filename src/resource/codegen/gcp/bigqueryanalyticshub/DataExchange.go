package bigqueryanalyticshub

type DataExchange struct {
	// Documentation describing the data exchange.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Base64 encoded image representing the data exchange.
	Icon string `json:"icon,omitempty" yaml:"icon,omitempty"`

	// The name of the location this data exchange.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact string `json:"primaryContact,omitempty" yaml:"primaryContact,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId string `json:"dataExchangeId,omitempty" yaml:"dataExchangeId,omitempty"`

	// Description of the data exchange.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}

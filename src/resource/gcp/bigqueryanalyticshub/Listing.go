package bigqueryanalyticshub

import types "DesignSphere_Server/src/resource/gcp/types"

type Listing struct {
	// Email or URL of the primary point of contact of the listing.
	PrimaryContact string `json:"primaryContact,omitempty" yaml:"primaryContact,omitempty"`

	/*
	   Shared dataset i.e. BigQuery dataset source.
	   Structure is documented below.
	*/
	BigqueryDataset types.Bigqueryanalyticshub_ListingBigqueryDataset `json:"bigqueryDataset,omitempty" yaml:"bigqueryDataset,omitempty"`

	// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
	RequestAccess string `json:"requestAccess,omitempty" yaml:"requestAccess,omitempty"`

	// Documentation describing the listing.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	ListingId string `json:"listingId,omitempty" yaml:"listingId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Categories of the listing. Up to two categories are allowed.
	Categories []string `json:"categories,omitempty" yaml:"categories,omitempty"`

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId string `json:"dataExchangeId,omitempty" yaml:"dataExchangeId,omitempty"`

	// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Details of the publisher who owns the listing and who can share the source data.
	   Structure is documented below.
	*/
	Publisher types.Bigqueryanalyticshub_ListingPublisher `json:"publisher,omitempty" yaml:"publisher,omitempty"`

	/*
	   If set, restricted export configuration will be propagated and enforced on the linked dataset.
	   Structure is documented below.
	*/
	RestrictedExportConfig types.Bigqueryanalyticshub_ListingRestrictedExportConfig `json:"restrictedExportConfig,omitempty" yaml:"restrictedExportConfig,omitempty"`

	/*
	   Details of the data provider who owns the source data.
	   Structure is documented below.
	*/
	DataProvider types.Bigqueryanalyticshub_ListingDataProvider `json:"dataProvider,omitempty" yaml:"dataProvider,omitempty"`

	// Base64 encoded image representing the listing.
	Icon string `json:"icon,omitempty" yaml:"icon,omitempty"`

	// The name of the location this data exchange listing.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}

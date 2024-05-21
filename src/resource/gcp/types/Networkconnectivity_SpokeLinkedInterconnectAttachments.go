package types

type Networkconnectivity_SpokeLinkedInterconnectAttachments struct {
	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer bool `json:"siteToSiteDataTransfer,omitempty" yaml:"siteToSiteDataTransfer,omitempty"`

	// The URIs of linked interconnect attachment resources
	Uris []string `json:"uris,omitempty" yaml:"uris,omitempty"`
}

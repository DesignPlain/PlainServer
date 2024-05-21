package types

type Networkconnectivity_SpokeLinkedRouterApplianceInstances struct {
	// The list of router appliance instances
	Instances []Networkconnectivity_SpokeLinkedRouterApplianceInstancesInstance `json:"instances,omitempty" yaml:"instances,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer bool `json:"siteToSiteDataTransfer,omitempty" yaml:"siteToSiteDataTransfer,omitempty"`
}

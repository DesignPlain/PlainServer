package types

type Ec2_getNetworkInsightsAnalysisReturnPathComponentInboundHeader struct {
	//
	DestinationAddresses []string `json:"destinationAddresses,omitempty" yaml:"destinationAddresses,omitempty"`

	//
	DestinationPortRanges []Ec2_getNetworkInsightsAnalysisReturnPathComponentInboundHeaderDestinationPortRange `json:"destinationPortRanges,omitempty" yaml:"destinationPortRanges,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	SourceAddresses []string `json:"sourceAddresses,omitempty" yaml:"sourceAddresses,omitempty"`

	//
	SourcePortRanges []Ec2_getNetworkInsightsAnalysisReturnPathComponentInboundHeaderSourcePortRange `json:"sourcePortRanges,omitempty" yaml:"sourcePortRanges,omitempty"`
}

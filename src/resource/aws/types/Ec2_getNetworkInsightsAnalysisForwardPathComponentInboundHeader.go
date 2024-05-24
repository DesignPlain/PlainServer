package types

type Ec2_getNetworkInsightsAnalysisForwardPathComponentInboundHeader struct {
	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	SourceAddresses []string `json:"sourceAddresses,omitempty" yaml:"sourceAddresses,omitempty"`

	//
	SourcePortRanges []Ec2_getNetworkInsightsAnalysisForwardPathComponentInboundHeaderSourcePortRange `json:"sourcePortRanges,omitempty" yaml:"sourcePortRanges,omitempty"`

	//
	DestinationAddresses []string `json:"destinationAddresses,omitempty" yaml:"destinationAddresses,omitempty"`

	//
	DestinationPortRanges []Ec2_getNetworkInsightsAnalysisForwardPathComponentInboundHeaderDestinationPortRange `json:"destinationPortRanges,omitempty" yaml:"destinationPortRanges,omitempty"`
}

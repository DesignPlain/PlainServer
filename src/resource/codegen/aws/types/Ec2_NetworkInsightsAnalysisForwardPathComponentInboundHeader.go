package types

type Ec2_NetworkInsightsAnalysisForwardPathComponentInboundHeader struct {
	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	SourceAddresses []string `json:"sourceAddresses,omitempty" yaml:"sourceAddresses,omitempty"`

	//
	SourcePortRanges []Ec2_NetworkInsightsAnalysisForwardPathComponentInboundHeaderSourcePortRange `json:"sourcePortRanges,omitempty" yaml:"sourcePortRanges,omitempty"`

	//
	DestinationAddresses []string `json:"destinationAddresses,omitempty" yaml:"destinationAddresses,omitempty"`

	//
	DestinationPortRanges []Ec2_NetworkInsightsAnalysisForwardPathComponentInboundHeaderDestinationPortRange `json:"destinationPortRanges,omitempty" yaml:"destinationPortRanges,omitempty"`
}

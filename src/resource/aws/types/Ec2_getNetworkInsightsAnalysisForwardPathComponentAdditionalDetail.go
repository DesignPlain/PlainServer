package types

type Ec2_getNetworkInsightsAnalysisForwardPathComponentAdditionalDetail struct {
	//
	AdditionalDetailType string `json:"additionalDetailType,omitempty" yaml:"additionalDetailType,omitempty"`

	//
	Components []Ec2_getNetworkInsightsAnalysisForwardPathComponentAdditionalDetailComponent `json:"components,omitempty" yaml:"components,omitempty"`
}

package types

type Ec2_getNetworkInsightsAnalysisReturnPathComponentAdditionalDetail struct {
	//
	AdditionalDetailType string `json:"additionalDetailType,omitempty" yaml:"additionalDetailType,omitempty"`

	//
	Components []Ec2_getNetworkInsightsAnalysisReturnPathComponentAdditionalDetailComponent `json:"components,omitempty" yaml:"components,omitempty"`
}

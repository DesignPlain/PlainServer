package types

type Ec2_NetworkInsightsAnalysisForwardPathComponentAdditionalDetail struct {
	//
	AdditionalDetailType string `json:"additionalDetailType,omitempty" yaml:"additionalDetailType,omitempty"`

	//
	Components []Ec2_NetworkInsightsAnalysisForwardPathComponentAdditionalDetailComponent `json:"components,omitempty" yaml:"components,omitempty"`
}

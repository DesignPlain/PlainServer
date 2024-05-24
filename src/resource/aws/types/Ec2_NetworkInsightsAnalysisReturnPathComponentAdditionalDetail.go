package types

type Ec2_NetworkInsightsAnalysisReturnPathComponentAdditionalDetail struct {
	//
	AdditionalDetailType string `json:"additionalDetailType,omitempty" yaml:"additionalDetailType,omitempty"`

	//
	Components []Ec2_NetworkInsightsAnalysisReturnPathComponentAdditionalDetailComponent `json:"components,omitempty" yaml:"components,omitempty"`
}

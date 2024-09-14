package securitycenter

type SourceIamPolicy struct {
	/*
	   The organization whose Cloud Security Command Center the Source
	   lives in.


	   - - -
	*/
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	//
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`

	//
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}

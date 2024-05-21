package types

type Diagflow_CxTestCaseTestConfig struct {
	/*
	   Flow name to start the test case with.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	   Only one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.
	*/
	Flow string `json:"flow,omitempty" yaml:"flow,omitempty"`

	/*
	   The page to start the test case with.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	   Only one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.
	*/
	Page string `json:"page,omitempty" yaml:"page,omitempty"`

	// Session parameters to be compared when calculating differences.
	TrackingParameters []string `json:"trackingParameters,omitempty" yaml:"trackingParameters,omitempty"`
}

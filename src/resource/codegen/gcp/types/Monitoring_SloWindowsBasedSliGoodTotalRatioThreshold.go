package types

type Monitoring_SloWindowsBasedSliGoodTotalRatioThreshold struct {
	/*
	   Basic SLI to evaluate to judge window quality.
	   Structure is documented below.
	*/
	BasicSliPerformance Monitoring_SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance `json:"basicSliPerformance,omitempty" yaml:"basicSliPerformance,omitempty"`

	/*
	   Request-based SLI to evaluate to judge window quality.
	   Structure is documented below.
	*/
	Performance Monitoring_SloWindowsBasedSliGoodTotalRatioThresholdPerformance `json:"performance,omitempty" yaml:"performance,omitempty"`

	/*
	   If window performance >= threshold, the window is counted
	   as good.
	*/
	Threshold float64 `json:"threshold,omitempty" yaml:"threshold,omitempty"`
}

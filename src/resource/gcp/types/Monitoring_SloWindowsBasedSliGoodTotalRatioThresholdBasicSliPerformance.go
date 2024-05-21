package types

type Monitoring_SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance struct {
	/*
	   Availability based SLI, dervied from count of requests made to this service that return successfully.
	   Structure is documented below.
	*/
	Availability Monitoring_SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability `json:"availability,omitempty" yaml:"availability,omitempty"`

	/*
	   Parameters for a latency threshold SLI.
	   Structure is documented below.
	*/
	Latency Monitoring_SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency `json:"latency,omitempty" yaml:"latency,omitempty"`

	/*
	   An optional set of locations to which this SLI is relevant.
	   Telemetry from other locations will not be used to calculate
	   performance for this SLI. If omitted, this SLI applies to all
	   locations in which the Service has activity. For service types
	   that don't support breaking down by location, setting this
	   field will result in an error.
	*/
	Locations []string `json:"locations,omitempty" yaml:"locations,omitempty"`

	/*
	   An optional set of RPCs to which this SLI is relevant.
	   Telemetry from other methods will not be used to calculate
	   performance for this SLI. If omitted, this SLI applies to all
	   the Service's methods. For service types that don't support
	   breaking down by method, setting this field will result in an
	   error.
	*/
	Methods []string `json:"methods,omitempty" yaml:"methods,omitempty"`

	/*
	   The set of API versions to which this SLI is relevant.
	   Telemetry from other API versions will not be used to
	   calculate performance for this SLI. If omitted,
	   this SLI applies to all API versions. For service types
	   that don't support breaking down by version, setting this
	   field will result in an error.
	*/
	Versions []string `json:"versions,omitempty" yaml:"versions,omitempty"`
}

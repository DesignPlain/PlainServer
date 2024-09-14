package cloudwatch

import types "libds/aws/types"

type InternetMonitor struct {
	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The percentage of the internet-facing traffic for your application that you want to monitor with this monitor.
	TrafficPercentageToMonitor int `json:"trafficPercentageToMonitor,omitempty" yaml:"trafficPercentageToMonitor,omitempty"`

	// Health event thresholds. A health event threshold percentage, for performance and availability, determines when Internet Monitor creates a health event when there's an internet issue that affects your application end users. See Health Events Config below.
	HealthEventsConfig types.Cloudwatch_InternetMonitorHealthEventsConfig `json:"healthEventsConfig,omitempty" yaml:"healthEventsConfig,omitempty"`

	// Publish internet measurements for Internet Monitor to an Amazon S3 bucket in addition to CloudWatch Logs.
	InternetMeasurementsLogDelivery types.Cloudwatch_InternetMonitorInternetMeasurementsLogDelivery `json:"internetMeasurementsLogDelivery,omitempty" yaml:"internetMeasurementsLogDelivery,omitempty"`

	// The maximum number of city-networks to monitor for your resources. A city-network is the location (city) where clients access your application resources from and the network or ASN, such as an internet service provider (ISP), that clients access the resources through. This limit helps control billing costs.
	MaxCityNetworksToMonitor int `json:"maxCityNetworksToMonitor,omitempty" yaml:"maxCityNetworksToMonitor,omitempty"`

	/*
	   The name of the monitor.

	   The following arguments are optional:
	*/
	MonitorName string `json:"monitorName,omitempty" yaml:"monitorName,omitempty"`

	// The resources to include in a monitor, which you provide as a set of Amazon Resource Names (ARNs).
	Resources []string `json:"resources,omitempty" yaml:"resources,omitempty"`

	// The status for a monitor. The accepted values for Status with the UpdateMonitor API call are the following: `ACTIVE` and `INACTIVE`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}

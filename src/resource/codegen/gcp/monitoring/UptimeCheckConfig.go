package monitoring

import types "libds/gcp/types"

type UptimeCheckConfig struct {
	/*
	   A Synthetic Monitor deployed to a Cloud Functions V2 instance.
	   Structure is documented below.
	*/
	SyntheticMonitor types.Monitoring_UptimeCheckConfigSyntheticMonitor `json:"syntheticMonitor,omitempty" yaml:"syntheticMonitor,omitempty"`

	/*
	   The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). See the accepted formats


	   - - -
	*/
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// User-supplied key/value data to be used for organizing and identifying the `UptimeCheckConfig` objects. The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `json:"userLabels,omitempty" yaml:"userLabels,omitempty"`

	/*
	   The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	   Structure is documented below.
	*/
	ContentMatchers []types.Monitoring_UptimeCheckConfigContentMatcher `json:"contentMatchers,omitempty" yaml:"contentMatchers,omitempty"`

	/*
	   The [monitored resource]
	   (https://cloud.google.com/monitoring/api/resources) associated with the
	   configuration. The following monitored resource types are supported for
	   uptime checks:
	*/
	MonitoredResource types.Monitoring_UptimeCheckConfigMonitoredResource `json:"monitoredResource,omitempty" yaml:"monitoredResource,omitempty"`

	/*
	   The group resource associated with the configuration.
	   Structure is documented below.
	*/
	ResourceGroup types.Monitoring_UptimeCheckConfigResourceGroup `json:"resourceGroup,omitempty" yaml:"resourceGroup,omitempty"`

	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period string `json:"period,omitempty" yaml:"period,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions []string `json:"selectedRegions,omitempty" yaml:"selectedRegions,omitempty"`

	/*
	   Contains information needed to make a TCP check.
	   Structure is documented below.
	*/
	TcpCheck types.Monitoring_UptimeCheckConfigTcpCheck `json:"tcpCheck,omitempty" yaml:"tcpCheck,omitempty"`

	/*
	   The checker type to use for the check. If the monitored resource type is `servicedirectory_service`, `checker_type` must be set to `VPC_CHECKERS`.
	   Possible values are: `STATIC_IP_CHECKERS`, `VPC_CHECKERS`.
	*/
	CheckerType string `json:"checkerType,omitempty" yaml:"checkerType,omitempty"`

	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Contains information needed to make an HTTP or HTTPS check.
	   Structure is documented below.
	*/
	HttpCheck types.Monitoring_UptimeCheckConfigHttpCheck `json:"httpCheck,omitempty" yaml:"httpCheck,omitempty"`
}

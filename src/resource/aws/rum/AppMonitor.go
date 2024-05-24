package rum

import types "DesignSphere_Server/src/resource/aws/types"

type AppMonitor struct {
	// configuration data for the app monitor. See app_monitor_configuration below.
	AppMonitorConfiguration types.Rum_AppMonitorAppMonitorConfiguration `json:"appMonitorConfiguration,omitempty" yaml:"appMonitorConfiguration,omitempty"`

	// Specifies whether this app monitor allows the web client to define and send custom events. If you omit this parameter, custom events are `DISABLED`. See custom_events below.
	CustomEvents types.Rum_AppMonitorCustomEvents `json:"customEvents,omitempty" yaml:"customEvents,omitempty"`

	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter  specifies whether RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges. Default value is `false`.
	CwLogEnabled bool `json:"cwLogEnabled,omitempty" yaml:"cwLogEnabled,omitempty"`

	// The top-level internet domain name for which your application has administrative authority.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// The name of the log stream.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

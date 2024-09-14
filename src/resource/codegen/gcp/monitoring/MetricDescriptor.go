package monitoring

import types "libds/gcp/types"

type MetricDescriptor struct {
	/*
	   Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might not be supported.
	   Possible values are: `BOOL`, `INT64`, `DOUBLE`, `STRING`, `DISTRIBUTION`.
	*/
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`

	/*
	   A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The launch stage of the metric definition.
	   Possible values are: `LAUNCH_STAGE_UNSPECIFIED`, `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, `DEPRECATED`.
	*/
	LaunchStage string `json:"launchStage,omitempty" yaml:"launchStage,omitempty"`

	/*
	   Metadata which can be used to guide usage of the metric.
	   Structure is documented below.
	*/
	Metadata types.Monitoring_MetricDescriptorMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	   Possible values are: `METRIC_KIND_UNSPECIFIED`, `GAUGE`, `DELTA`, `CUMULATIVE`.
	*/
	MetricKind string `json:"metricKind,omitempty" yaml:"metricKind,omitempty"`

	/*
	   The units in which the metric value is reported. It is only applicable if the
	   valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	   the stored metric values.
	   Different systems may scale the values to be more easily displayed (so a value of
	   0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	   3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	   thousands of bytes, no matter how it may be displayed.
	   If you want a custom metric to record the exact number of CPU-seconds used by a job,
	   you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	   1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	   12005.
	   Alternatively, if you want a custom metric to record data in a more granular way, you
	   can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	   12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	   The supported units are a subset of The Unified Code for Units of Measure standard.
	   More info can be found in the API documentation
	   (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	*/
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// A detailed description of the metric, which can be used in documentation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	   Structure is documented below.
	*/
	Labels []types.Monitoring_MetricDescriptorLabel `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

package logging

import types "libds/gcp/types"

type Metric struct {
	/*
	   The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	   describes the bucket boundaries used to create a histogram of the extracted values.
	   Structure is documented below.
	*/
	BucketOptions types.Logging_MetricBucketOptions `json:"bucketOptions,omitempty" yaml:"bucketOptions,omitempty"`

	/*
	   A description of this metric, which is used in documentation. The maximum length of the
	   description is 8000 characters.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	   is used to match log entries.


	   - - -
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	/*
	   A map from a label key string to an extractor expression which is used to extract data from a log
	   entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	   have an associated extractor expression in this map. The syntax of the extractor expression is
	   the same as for the valueExtractor field.
	*/
	LabelExtractors map[string]string `json:"labelExtractors,omitempty" yaml:"labelExtractors,omitempty"`

	/*
	   The optional metric descriptor associated with the logs-based metric.
	   If unspecified, it uses a default metric descriptor with a DELTA metric kind,
	   INT64 value type, with no labels and a unit of "1". Such a metric counts the
	   number of log entries matching the filter expression.
	   Structure is documented below.
	*/
	MetricDescriptor types.Logging_MetricMetricDescriptor `json:"metricDescriptor,omitempty" yaml:"metricDescriptor,omitempty"`

	/*
	   The client-assigned metric identifier. Examples - "error_count", "nginx/requests".
	   Metric identifiers are limited to 100 characters and can include only the following
	   characters A-Z, a-z, 0-9, and the special characters _-.,+!-',()%!!(MISSING)/(MISSING). The forward-slash
	   character (/) denotes a hierarchy of name pieces, and it cannot be the first character
	   of the name.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects
	   are supported. The bucket has to be in the same project as the metric.
	*/
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	/*
	   A valueExtractor is required when using a distribution logs-based metric to extract the values to
	   record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	   REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	   the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	   (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	   log entry field. The value of the field is converted to a string before applying the regex. It is an
	   error to specify a regex that does not include exactly one capture group.
	*/
	ValueExtractor string `json:"valueExtractor,omitempty" yaml:"valueExtractor,omitempty"`

	// If set to True, then this metric is disabled and it does not generate any points.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}

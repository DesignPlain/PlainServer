package types

type Monitoring_AlertPolicyCondition struct {
	/*
	   A Monitoring Query Language query that outputs a boolean stream
	   Structure is documented below.
	*/
	ConditionMonitoringQueryLanguage Monitoring_AlertPolicyConditionConditionMonitoringQueryLanguage `json:"conditionMonitoringQueryLanguage,omitempty" yaml:"conditionMonitoringQueryLanguage,omitempty"`

	/*
	   A condition type that allows alert policies to be defined using
	   Prometheus Query Language (PromQL).
	   The PrometheusQueryLanguageCondition message contains information
	   from a Prometheus alerting rule and its associated rule group.
	   Structure is documented below.
	*/
	ConditionPrometheusQueryLanguage Monitoring_AlertPolicyConditionConditionPrometheusQueryLanguage `json:"conditionPrometheusQueryLanguage,omitempty" yaml:"conditionPrometheusQueryLanguage,omitempty"`

	/*
	   A condition that compares a time series against a
	   threshold.
	   Structure is documented below.
	*/
	ConditionThreshold Monitoring_AlertPolicyConditionConditionThreshold `json:"conditionThreshold,omitempty" yaml:"conditionThreshold,omitempty"`

	/*
	   A short name or phrase used to identify the
	   condition in dashboards, notifications, and
	   incidents. To avoid confusion, don't use the same
	   display name for multiple conditions in the same
	   policy.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   (Output)
	   The unique resource name for this condition.
	   Its syntax is:
	   projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	   [CONDITION_ID] is assigned by Stackdriver Monitoring when
	   the condition is created as part of a new or updated alerting
	   policy.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A condition that checks that a time series
	   continues to receive new data points.
	   Structure is documented below.
	*/
	ConditionAbsent Monitoring_AlertPolicyConditionConditionAbsent `json:"conditionAbsent,omitempty" yaml:"conditionAbsent,omitempty"`

	/*
	   A condition that checks for log messages matching given constraints.
	   If set, no other conditions can be present.
	   Structure is documented below.
	*/
	ConditionMatchedLog Monitoring_AlertPolicyConditionConditionMatchedLog `json:"conditionMatchedLog,omitempty" yaml:"conditionMatchedLog,omitempty"`
}

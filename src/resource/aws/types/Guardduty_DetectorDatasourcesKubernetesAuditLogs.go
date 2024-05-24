package types

type Guardduty_DetectorDatasourcesKubernetesAuditLogs struct {
	/*
	   If true, enables Kubernetes audit logs as a data source for [Kubernetes protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html).
	   Defaults to `true`.
	*/
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`
}

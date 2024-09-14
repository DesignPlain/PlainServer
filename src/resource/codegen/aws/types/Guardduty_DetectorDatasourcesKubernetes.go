package types

type Guardduty_DetectorDatasourcesKubernetes struct {
	/*
	   Configures Kubernetes audit logs as a data source for [Kubernetes protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html).
	   See Kubernetes Audit Logs below for more details.
	*/
	AuditLogs Guardduty_DetectorDatasourcesKubernetesAuditLogs `json:"auditLogs,omitempty" yaml:"auditLogs,omitempty"`
}

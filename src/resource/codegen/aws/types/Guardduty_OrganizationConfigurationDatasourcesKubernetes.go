package types

type Guardduty_OrganizationConfigurationDatasourcesKubernetes struct {
	/*
	   Enable Kubernetes Audit Logs Monitoring automatically for new member accounts. [Kubernetes protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html).
	   See Kubernetes Audit Logs below for more details.
	*/
	AuditLogs Guardduty_OrganizationConfigurationDatasourcesKubernetesAuditLogs `json:"auditLogs,omitempty" yaml:"auditLogs,omitempty"`
}

package types

type Container_getClusterProtectConfigWorkloadConfig struct {
	// Sets which mode of auditing should be used for the cluster's workloads. Accepted values are DISABLED, BASIC.
	AuditMode string `json:"auditMode,omitempty" yaml:"auditMode,omitempty"`
}

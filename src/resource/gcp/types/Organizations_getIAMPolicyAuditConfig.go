package types

type Organizations_getIAMPolicyAuditConfig struct {
	// A nested block that defines the operations you'd like to log.
	AuditLogConfigs []Organizations_getIAMPolicyAuditConfigAuditLogConfig `json:"auditLogConfigs,omitempty" yaml:"auditLogConfigs,omitempty"`

	// Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

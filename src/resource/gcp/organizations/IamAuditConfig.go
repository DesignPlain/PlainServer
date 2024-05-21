package organizations

import types "DesignSphere_Server/src/resource/gcp/types"

type IamAuditConfig struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []types.Organizations_IamAuditConfigAuditLogConfig `json:"auditLogConfigs,omitempty" yaml:"auditLogConfigs,omitempty"`

	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

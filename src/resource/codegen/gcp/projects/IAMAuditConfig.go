package projects

import types "libds/gcp/types"

type IAMAuditConfig struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []types.Projects_IAMAuditConfigAuditLogConfig `json:"auditLogConfigs,omitempty" yaml:"auditLogConfigs,omitempty"`

	/*
	   The project id of the target project. This is not
	   inferred from the provider.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

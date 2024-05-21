package folder

import types "DesignSphere_Server/src/resource/gcp/types"

type IamAuditConfig struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []types.Folder_IamAuditConfigAuditLogConfig `json:"auditLogConfigs,omitempty" yaml:"auditLogConfigs,omitempty"`

	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_folder\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

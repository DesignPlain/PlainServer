package types

type Organizations_getIAMPolicyAuditConfigAuditLogConfig struct {
	// Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
	ExemptedMembers []string `json:"exemptedMembers,omitempty" yaml:"exemptedMembers,omitempty"`

	// Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`
}

package types

type Fsx_getWindowsFileSystemAuditLogConfiguration struct {
	//
	FileAccessAuditLogLevel string `json:"fileAccessAuditLogLevel,omitempty" yaml:"fileAccessAuditLogLevel,omitempty"`

	//
	FileShareAccessAuditLogLevel string `json:"fileShareAccessAuditLogLevel,omitempty" yaml:"fileShareAccessAuditLogLevel,omitempty"`

	//
	AuditLogDestination string `json:"auditLogDestination,omitempty" yaml:"auditLogDestination,omitempty"`
}

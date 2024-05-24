package auditmanager

type AccountRegistration struct {
	// Identifier for the delegated administrator account.
	DelegatedAdminAccount string `json:"delegatedAdminAccount,omitempty" yaml:"delegatedAdminAccount,omitempty"`

	// Flag to deregister AuditManager in the account upon destruction. Defaults to `false` (ie. AuditManager will remain active in the account, even if this resource is removed).
	DeregisterOnDestroy bool `json:"deregisterOnDestroy,omitempty" yaml:"deregisterOnDestroy,omitempty"`

	// KMS key identifier.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}

package types

type Accesscontextmanager_AccessLevelConditionDevicePolicy struct {
	/*
	   A list of allowed OS versions.
	   An empty list allows all types and all versions.
	   Structure is documented below.
	*/
	OsConstraints []Accesscontextmanager_AccessLevelConditionDevicePolicyOsConstraint `json:"osConstraints,omitempty" yaml:"osConstraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	RequireAdminApproval bool `json:"requireAdminApproval,omitempty" yaml:"requireAdminApproval,omitempty"`

	// Whether the device needs to be corp owned.
	RequireCorpOwned bool `json:"requireCorpOwned,omitempty" yaml:"requireCorpOwned,omitempty"`

	/*
	   Whether or not screenlock is required for the DevicePolicy
	   to be true. Defaults to false.
	*/
	RequireScreenLock bool `json:"requireScreenLock,omitempty" yaml:"requireScreenLock,omitempty"`

	/*
	   A list of allowed device management levels.
	   An empty list allows all management levels.
	   Each value may be one of: `MANAGEMENT_UNSPECIFIED`, `NONE`, `BASIC`, `COMPLETE`.
	*/
	AllowedDeviceManagementLevels []string `json:"allowedDeviceManagementLevels,omitempty" yaml:"allowedDeviceManagementLevels,omitempty"`

	/*
	   A list of allowed encryptions statuses.
	   An empty list allows all statuses.
	   Each value may be one of: `ENCRYPTION_UNSPECIFIED`, `ENCRYPTION_UNSUPPORTED`, `UNENCRYPTED`, `ENCRYPTED`.
	*/
	AllowedEncryptionStatuses []string `json:"allowedEncryptionStatuses,omitempty" yaml:"allowedEncryptionStatuses,omitempty"`
}

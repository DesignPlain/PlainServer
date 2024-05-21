package projects

type DefaultServiceAccounts struct {
	// The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// The project ID where service accounts are created.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The action to be performed in the default service accounts on the resource destroy.
	   Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	   If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
	   If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
	*/
	RestorePolicy string `json:"restorePolicy,omitempty" yaml:"restorePolicy,omitempty"`
}

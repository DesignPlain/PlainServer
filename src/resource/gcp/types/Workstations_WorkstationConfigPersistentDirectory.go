package types

type Workstations_WorkstationConfigPersistentDirectory struct {
	/*
	   A directory to persist across workstation sessions, backed by a Compute Engine regional persistent disk. Can only be updated if not empty during creation.
	   Structure is documented below.
	*/
	GcePd Workstations_WorkstationConfigPersistentDirectoryGcePd `json:"gcePd,omitempty" yaml:"gcePd,omitempty"`

	// Location of this directory in the running workstation.
	MountPath string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`
}

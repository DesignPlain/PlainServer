package types

type Workstations_WorkstationConfigHost struct {
	/*
	   A runtime using a Compute Engine instance.
	   Structure is documented below.
	*/
	GceInstance Workstations_WorkstationConfigHostGceInstance `json:"gceInstance,omitempty" yaml:"gceInstance,omitempty"`
}

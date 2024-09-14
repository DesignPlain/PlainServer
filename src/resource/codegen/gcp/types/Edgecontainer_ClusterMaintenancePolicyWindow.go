package types

type Edgecontainer_ClusterMaintenancePolicyWindow struct {
	/*
	   Represents an arbitrary window of time that recurs.
	   Structure is documented below.
	*/
	RecurringWindow Edgecontainer_ClusterMaintenancePolicyWindowRecurringWindow `json:"recurringWindow,omitempty" yaml:"recurringWindow,omitempty"`
}

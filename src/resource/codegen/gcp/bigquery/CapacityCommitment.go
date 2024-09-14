package bigquery

type CapacityCommitment struct {
	// Number of slots in this commitment.
	SlotCount int `json:"slotCount,omitempty" yaml:"slotCount,omitempty"`

	/*
	   The optional capacity commitment ID. Capacity commitment name will be generated automatically if this field is
	   empty. This field must only contain lower case alphanumeric characters or dashes. The first and last character
	   cannot be a dash. Max length is 64 characters. NOTE: this ID won't be kept if the capacity commitment is split
	   or merged.
	*/
	CapacityCommitmentId string `json:"capacityCommitmentId,omitempty" yaml:"capacityCommitmentId,omitempty"`

	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`

	// If true, fail the request if another project in the organization has a capacity commitment.
	EnforceSingleAdminProjectPerOrg string `json:"enforceSingleAdminProjectPerOrg,omitempty" yaml:"enforceSingleAdminProjectPerOrg,omitempty"`

	/*
	   The geographic location where the transfer config should reside.
	   Examples: US, EU, asia-northeast1. The default value is US.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Capacity commitment plan. Valid values are at https://cloud.google.com/bigquery/docs/reference/reservations/rpc/google.cloud.bigquery.reservation.v1#commitmentplan


	   - - -
	*/
	Plan string `json:"plan,omitempty" yaml:"plan,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The plan this capacity commitment is converted to after commitmentEndTime passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for some commitment plans.
	RenewalPlan string `json:"renewalPlan,omitempty" yaml:"renewalPlan,omitempty"`
}

package types

type Container_ClusterMaintenancePolicyMaintenanceExclusionExclusionOptions struct {
	/*
	   The scope of automatic upgrades to restrict in the exclusion window. One of: --NO_UPGRADES | NO_MINOR_UPGRADES | NO_MINOR_OR_NODE_UPGRADES--

	   Specify `start_time` and `end_time` in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) "Zulu" date format.  The start time's date is
	   the initial date that the window starts, and the end time is used for calculating duration.Specify `recurrence` in
	   [RFC5545](https://tools.ietf.org/html/rfc5545#section-3.8.5.3) RRULE format, to specify when this recurs.
	   Note that GKE may accept other formats, but will return values in UTC, causing a permanent diff.

	   Examples:
	*/
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}

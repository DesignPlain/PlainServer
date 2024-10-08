package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigHybridOptionsTableOptions struct {
	/*
	   The columns that are the primary keys for table objects included in ContentItem. A copy of this
	   cell's value will stored alongside alongside each finding so that the finding can be traced to
	   the specific row it came from. No more than 3 may be provided.
	   Structure is documented below.
	*/
	IdentifyingFields []Dataloss_PreventionJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingField `json:"identifyingFields,omitempty" yaml:"identifyingFields,omitempty"`
}

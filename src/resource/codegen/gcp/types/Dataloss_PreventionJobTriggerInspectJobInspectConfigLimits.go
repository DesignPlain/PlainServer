package types

type Dataloss_PreventionJobTriggerInspectJobInspectConfigLimits struct {
	// Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
	MaxFindingsPerItem int `json:"maxFindingsPerItem,omitempty" yaml:"maxFindingsPerItem,omitempty"`

	// Max number of findings that will be returned per request/job. The maximum returned is 2000.
	MaxFindingsPerRequest int `json:"maxFindingsPerRequest,omitempty" yaml:"maxFindingsPerRequest,omitempty"`

	/*
	   Configuration of findings limit given for specified infoTypes.
	   Structure is documented below.
	*/
	MaxFindingsPerInfoTypes []Dataloss_PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType `json:"maxFindingsPerInfoTypes,omitempty" yaml:"maxFindingsPerInfoTypes,omitempty"`
}

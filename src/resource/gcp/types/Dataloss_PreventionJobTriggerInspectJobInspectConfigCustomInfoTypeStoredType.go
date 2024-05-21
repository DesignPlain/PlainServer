package types

type Dataloss_PreventionJobTriggerInspectJobInspectConfigCustomInfoTypeStoredType struct {
	/*
	   (Output)
	   The creation timestamp of an inspectTemplate. Set by the server.
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	/*
	   Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
	   or `projects/project-id/storedInfoTypes/432452342`.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

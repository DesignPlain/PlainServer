package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleDictionaryCloudStoragePath struct {
	// A url representing a file or path (no wildcards) in Cloud Storage. Example: `gs://[BUCKET_NAME]/dictionary.txt`
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}

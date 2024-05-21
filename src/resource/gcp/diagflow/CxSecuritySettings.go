package diagflow

import types "DesignSphere_Server/src/resource/gcp/types"

type CxSecuritySettings struct {
	/*
	   Controls conversation exporting settings to Insights after conversation is completed.
	   If retentionStrategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
	   Structure is documented below.
	*/
	InsightsExportSettings types.Diagflow_CxSecuritySettingsInsightsExportSettings `json:"insightsExportSettings,omitempty" yaml:"insightsExportSettings,omitempty"`

	/*
	   The location these settings are located in. Settings can only be applied to an agent in the same location.
	   See [Available Regions](https://cloud.google.com/dialogflow/cx/docs/concept/region#avail) for a list of supported locations.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   List of types of data to remove when retention settings triggers purge.
	   Each value may be one of: `DIALOGFLOW_HISTORY`.
	*/
	PurgeDataTypes []string `json:"purgeDataTypes,omitempty" yaml:"purgeDataTypes,omitempty"`

	/*
	   Defines what types of data to redact. If not set, defaults to not redacting any kind of data.
	   - REDACT_DISK_STORAGE: On data to be written to disk or similar devices that are capable of holding data even if power is disconnected. This includes data that are temporarily saved on disk.
	   Possible values are: `REDACT_DISK_STORAGE`.
	*/
	RedactionScope string `json:"redactionScope,omitempty" yaml:"redactionScope,omitempty"`

	/*
	   Defines how we redact data. If not set, defaults to not redacting.
	   - REDACT_WITH_SERVICE: Call redaction service to clean up the data to be persisted.
	   Possible values are: `REDACT_WITH_SERVICE`.
	*/
	RedactionStrategy string `json:"redactionStrategy,omitempty" yaml:"redactionStrategy,omitempty"`

	/*
	   Defines how long we retain persisted data that contains sensitive info. Only one of `retention_window_days` and `retention_strategy` may be set.
	   - REMOVE_AFTER_CONVERSATION: Removes data when the conversation ends. If there is no conversation explicitly established, a default conversation ends when the corresponding Dialogflow session ends.
	   Possible values are: `REMOVE_AFTER_CONVERSATION`.
	*/
	RetentionStrategy string `json:"retentionStrategy,omitempty" yaml:"retentionStrategy,omitempty"`

	/*
	   [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. If empty, Dialogflow replaces sensitive info with [redacted] text.
	   Note: deidentifyTemplate must be located in the same region as the SecuritySettings.
	   Format: projects/<Project ID>/locations/<Location ID>/deidentifyTemplates/<Template ID> OR organizations/<Organization ID>/locations/<Location ID>/deidentifyTemplates/<Template ID>
	*/
	DeidentifyTemplate string `json:"deidentifyTemplate,omitempty" yaml:"deidentifyTemplate,omitempty"`

	/*
	   The human-readable name of the security settings, unique within the location.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config.
	   Note: inspectTemplate must be located in the same region as the SecuritySettings.
	   Format: projects/<Project ID>/locations/<Location ID>/inspectTemplates/<Template ID> OR organizations/<Organization ID>/locations/<Location ID>/inspectTemplates/<Template ID>
	*/
	InspectTemplate string `json:"inspectTemplate,omitempty" yaml:"inspectTemplate,omitempty"`

	/*
	   Retains the data for the specified number of days. User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.
	   Only one of `retention_window_days` and `retention_strategy` may be set.
	*/
	RetentionWindowDays int `json:"retentionWindowDays,omitempty" yaml:"retentionWindowDays,omitempty"`

	/*
	   Controls audio export settings for post-conversation analytics when ingesting audio to conversations.
	   If retention_strategy is set to REMOVE_AFTER_CONVERSATION or gcs_bucket is empty, audio export is disabled.
	   If audio export is enabled, audio is recorded and saved to gcs_bucket, subject to retention policy of gcs_bucket.
	   This setting won't effect audio input for implicit sessions via [Sessions.DetectIntent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.sessions/detectIntent#google.cloud.dialogflow.cx.v3.Sessions.DetectIntent).
	   Structure is documented below.
	*/
	AudioExportSettings types.Diagflow_CxSecuritySettingsAudioExportSettings `json:"audioExportSettings,omitempty" yaml:"audioExportSettings,omitempty"`
}

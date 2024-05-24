package connect

import types "DesignSphere_Server/src/resource/aws/types"

type InstanceStorageConfig struct {
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// A valid resource type. Valid Values: `AGENT_EVENTS` | `ATTACHMENTS` | `CALL_RECORDINGS` | `CHAT_TRANSCRIPTS` | `CONTACT_EVALUATIONS` | `CONTACT_TRACE_RECORDS` | `MEDIA_STREAMS` | `REAL_TIME_CONTACT_ANALYSIS_SEGMENTS` | `SCHEDULED_REPORTS` | `SCREEN_RECORDINGS`.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Specifies the storage configuration options for the Connect Instance. Documented below.
	StorageConfig types.Connect_InstanceStorageConfigStorageConfig `json:"storageConfig,omitempty" yaml:"storageConfig,omitempty"`
}

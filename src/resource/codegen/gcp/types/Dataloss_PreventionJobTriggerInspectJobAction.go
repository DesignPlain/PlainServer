package types

type Dataloss_PreventionJobTriggerInspectJobAction struct {
	// Publish findings of a DlpJob to Data Catalog.
	PublishFindingsToCloudDataCatalog Dataloss_PreventionJobTriggerInspectJobActionPublishFindingsToCloudDataCatalog `json:"publishFindingsToCloudDataCatalog,omitempty" yaml:"publishFindingsToCloudDataCatalog,omitempty"`

	// Publish the result summary of a DlpJob to the Cloud Security Command Center.
	PublishSummaryToCscc Dataloss_PreventionJobTriggerInspectJobActionPublishSummaryToCscc `json:"publishSummaryToCscc,omitempty" yaml:"publishSummaryToCscc,omitempty"`

	// Enable Stackdriver metric dlp.googleapis.com/findingCount.
	PublishToStackdriver Dataloss_PreventionJobTriggerInspectJobActionPublishToStackdriver `json:"publishToStackdriver,omitempty" yaml:"publishToStackdriver,omitempty"`

	/*
	   If set, the detailed findings will be persisted to the specified OutputStorageConfig. Only a single instance of this action can be specified. Compatible with: Inspect, Risk
	   Structure is documented below.
	*/
	SaveFindings Dataloss_PreventionJobTriggerInspectJobActionSaveFindings `json:"saveFindings,omitempty" yaml:"saveFindings,omitempty"`

	/*
	   Create a de-identified copy of the requested table or files.
	   Structure is documented below.
	*/
	Deidentify Dataloss_PreventionJobTriggerInspectJobActionDeidentify `json:"deidentify,omitempty" yaml:"deidentify,omitempty"`

	// Sends an email when the job completes. The email goes to IAM project owners and technical Essential Contacts.
	JobNotificationEmails Dataloss_PreventionJobTriggerInspectJobActionJobNotificationEmails `json:"jobNotificationEmails,omitempty" yaml:"jobNotificationEmails,omitempty"`

	/*
	   Publish a message into a given Pub/Sub topic when the job completes.
	   Structure is documented below.
	*/
	PubSub Dataloss_PreventionJobTriggerInspectJobActionPubSub `json:"pubSub,omitempty" yaml:"pubSub,omitempty"`
}

package bigquery

import types "DesignSphere_Server/src/resource/gcp/types"

type DataTransferConfig struct {
	/*
	   Email notifications will be sent according to these preferences to the
	   email address of the user who owns this transfer config.
	   Structure is documented below.
	*/
	EmailPreferences types.Bigquery_DataTransferConfigEmailPreferences `json:"emailPreferences,omitempty" yaml:"emailPreferences,omitempty"`

	/*
	   The geographic location where the transfer config should reside.
	   Examples: US, EU, asia-northeast1. The default value is US.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Data transfer schedule. If the data source does not support a custom
	   schedule, this should be empty. If it is empty, the default value for
	   the data source will be used. The specified times are in UTC. Examples
	   of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
	   jun 13:15, and first sunday of quarter 00:00. See more explanation
	   about the format here:
	   https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	   NOTE: the granularity should be at least 8 hours, or less frequent.
	*/
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	/*
	   Service account email. If this field is set, transfer config will
	   be created with this service account credentials. It requires that
	   requesting user calling this API has permissions to act as this service account.
	*/
	ServiceAccountName string `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`

	/*
	   The number of days to look back to automatically refresh the data.
	   For example, if dataRefreshWindowDays = 10, then every day BigQuery
	   reingests data for [today-10, today-1], rather than ingesting data for
	   just [today-1]. Only valid if the data source supports the feature.
	   Set the value to 0 to use the default value.
	*/
	DataRefreshWindowDays int `json:"dataRefreshWindowDays,omitempty" yaml:"dataRefreshWindowDays,omitempty"`

	// The BigQuery target dataset id.
	DestinationDatasetId string `json:"destinationDatasetId,omitempty" yaml:"destinationDatasetId,omitempty"`

	// The user specified display name for the transfer config.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Pub/Sub topic where notifications will be sent after transfer runs
	   associated with this transfer config finish.
	*/
	NotificationPubsubTopic string `json:"notificationPubsubTopic,omitempty" yaml:"notificationPubsubTopic,omitempty"`

	/*
	   Different parameters are configured primarily using the the `params` field on this
	   resource. This block contains the parameters which contain secrets or passwords so that they can be marked
	   sensitive and hidden from plan output. The name of the field, eg: secret_access_key, will be the key
	   in the `params` map in the api request.
	   Credentials may not be specified in both locations and will cause an error. Changing from one location
	   to a different credential configuration in the config will require an apply to update state.
	   Structure is documented below.
	*/
	SensitiveParams types.Bigquery_DataTransferConfigSensitiveParams `json:"sensitiveParams,omitempty" yaml:"sensitiveParams,omitempty"`

	// When set to true, no runs are scheduled for a given transfer.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer'
	   section for each data source. For example the parameters for Cloud Storage transfers are listed here:
	   https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	   --NOTE-- : If you are attempting to update a parameter that cannot be updated (due to api limitations) please force recreation of the resource.


	   - - -
	*/
	Params map[string]string `json:"params,omitempty" yaml:"params,omitempty"`

	/*
	   Options customizing the data transfer schedule.
	   Structure is documented below.
	*/
	ScheduleOptions types.Bigquery_DataTransferConfigScheduleOptions `json:"scheduleOptions,omitempty" yaml:"scheduleOptions,omitempty"`

	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceId string `json:"dataSourceId,omitempty" yaml:"dataSourceId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}

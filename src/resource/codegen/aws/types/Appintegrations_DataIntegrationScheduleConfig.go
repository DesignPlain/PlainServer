package types

type Appintegrations_DataIntegrationScheduleConfig struct {
	// The start date for objects to import in the first flow run as an Unix/epoch timestamp in milliseconds or in ISO-8601 format. This needs to be a time in the past, meaning that the data created or updated before this given date will not be downloaded.
	FirstExecutionFrom string `json:"firstExecutionFrom,omitempty" yaml:"firstExecutionFrom,omitempty"`

	// The name of the object to pull from the data source. Examples of objects in Salesforce include `Case`, `Account`, or `Lead`.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	// How often the data should be pulled from data source. Examples include `rate(1 hour)`, `rate(3 hours)`, `rate(1 day)`.
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`
}

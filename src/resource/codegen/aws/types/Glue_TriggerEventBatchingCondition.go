package types

type Glue_TriggerEventBatchingCondition struct {
	// Number of events that must be received from Amazon EventBridge before EventBridge  event trigger fires.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	// Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received. Default value is `900`.
	BatchWindow int `json:"batchWindow,omitempty" yaml:"batchWindow,omitempty"`
}

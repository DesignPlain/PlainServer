package types

type Mskconnect_ConnectorLogDelivery struct {
	// The workers can send worker logs to different destination  This configuration specifies the details of these destinations. See below.
	WorkerLogDelivery Mskconnect_ConnectorLogDeliveryWorkerLogDelivery `json:"workerLogDelivery,omitempty" yaml:"workerLogDelivery,omitempty"`
}

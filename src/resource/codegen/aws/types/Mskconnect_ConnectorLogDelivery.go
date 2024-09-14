package types

type Mskconnect_ConnectorLogDelivery struct {
	// The workers can send worker logs to different destination types. This configuration specifies the details of these destinations. See `worker_log_delivery` Block for details.
	WorkerLogDelivery Mskconnect_ConnectorLogDeliveryWorkerLogDelivery `json:"workerLogDelivery,omitempty" yaml:"workerLogDelivery,omitempty"`
}

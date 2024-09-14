package types

type Tpu_V2VmSymptom struct {
	/*
	   (Output)
	   Detailed information of the current Symptom.
	*/
	Details string `json:"details,omitempty" yaml:"details,omitempty"`

	/*
	   (Output)
	   Type of the Symptom.
	*/
	SymptomType string `json:"symptomType,omitempty" yaml:"symptomType,omitempty"`

	/*
	   (Output)
	   A string used to uniquely distinguish a worker within a TPU node.
	*/
	WorkerId string `json:"workerId,omitempty" yaml:"workerId,omitempty"`

	/*
	   (Output)
	   Timestamp when the Symptom is created.
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`
}

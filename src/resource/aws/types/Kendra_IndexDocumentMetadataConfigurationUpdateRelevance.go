package types

type Kendra_IndexDocumentMetadataConfigurationUpdateRelevance struct {
	// Determines how values should be interpreted. For more information, refer to [RankOrder](https://docs.aws.amazon.com/kendra/latest/dg/API_Relevance.html#Kendra-Type-Relevance-RankOrder).
	RankOrder string `json:"rankOrder,omitempty" yaml:"rankOrder,omitempty"`

	// A list of values that should be given a different boost when they appear in the result list. For more information, refer to [ValueImportanceMap](https://docs.aws.amazon.com/kendra/latest/dg/API_Relevance.html#Kendra-Type-Relevance-ValueImportanceMap).
	ValuesImportanceMap map[string]int `json:"valuesImportanceMap,omitempty" yaml:"valuesImportanceMap,omitempty"`

	// Specifies the time period that the boost applies to. For more information, refer to [Duration](https://docs.aws.amazon.com/kendra/latest/dg/API_Relevance.html#Kendra-Type-Relevance-Duration).
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	// Indicates that this field determines how "fresh" a document is. For more information, refer to [Freshness](https://docs.aws.amazon.com/kendra/latest/dg/API_Relevance.html#Kendra-Type-Relevance-Freshness).
	Freshness bool `json:"freshness,omitempty" yaml:"freshness,omitempty"`

	// The relative importance of the field in the search. Larger numbers provide more of a boost than smaller numbers. Minimum value of 1. Maximum value of 10.
	Importance int `json:"importance,omitempty" yaml:"importance,omitempty"`
}

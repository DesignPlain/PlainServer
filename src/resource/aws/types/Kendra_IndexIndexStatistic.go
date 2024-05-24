package types

type Kendra_IndexIndexStatistic struct {
	// A block that specifies the number of text documents indexed. Detailed below.
	TextDocumentStatistics []Kendra_IndexIndexStatisticTextDocumentStatistic `json:"textDocumentStatistics,omitempty" yaml:"textDocumentStatistics,omitempty"`

	// A block that specifies the number of question and answer topics in the index. Detailed below.
	FaqStatistics []Kendra_IndexIndexStatisticFaqStatistic `json:"faqStatistics,omitempty" yaml:"faqStatistics,omitempty"`
}

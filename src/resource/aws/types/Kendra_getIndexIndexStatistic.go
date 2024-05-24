package types

type Kendra_getIndexIndexStatistic struct {
	// Block that specifies the number of question and answer topics in the index. Documented below.
	FaqStatistics []Kendra_getIndexIndexStatisticFaqStatistic `json:"faqStatistics,omitempty" yaml:"faqStatistics,omitempty"`

	// A block that specifies the number of text documents indexed.
	TextDocumentStatistics []Kendra_getIndexIndexStatisticTextDocumentStatistic `json:"textDocumentStatistics,omitempty" yaml:"textDocumentStatistics,omitempty"`
}

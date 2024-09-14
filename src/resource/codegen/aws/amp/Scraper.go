package amp

import types "libds/aws/types"

type Scraper struct {
	//
	Timeouts types.Amp_ScraperTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// a name to associate with the managed scraper. This is for your use, and does not need to be unique.
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// Configuration block for the managed scraper to send metrics to. See `destination`.
	Destination types.Amp_ScraperDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The configuration file to use in the new scraper. For more information, see [Scraper configuration](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html#AMP-collector-configuration).
	ScrapeConfiguration string `json:"scrapeConfiguration,omitempty" yaml:"scrapeConfiguration,omitempty"`

	/*
	   Configuration block to specify where the managed scraper will collect metrics from. See `source`.

	   The following arguments are optional:
	*/
	Source types.Amp_ScraperSource `json:"source,omitempty" yaml:"source,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

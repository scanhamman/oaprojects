package jsonfile

type Projects []Project

type Project struct {
	Acronym                          *string  `json:"acronym"`
	Callidentifier                   *string  `json:"callidentifier"`
	Code                             *string  `json:"code"`
	Enddate                          *string  `json:"enddate"`
	Funding                          []Fund   `json:"funding"`
	Granted                          Grant    `json:"granted,omitempty"`
	H2020Programme                   []H2020  `json:"h2020programme"`
	ID                               *string  `json:"id"`
	Openaccessmandatefordataset      *bool    `json:"openaccessmandatefordataset"`
	Openaccessmandateforpublications *bool    `json:"openaccessmandateforpublications"`
	Startdate                        *string  `json:"startdate"`
	Subject                          []string `json:"subject"`
	Summary                          *string  `json:"summary"`
	Title                            *string  `json:"title"`
	Keywords                         *string  `json:"keywords,omitempty"`
	Websiteurl                       *string  `json:"websiteurl"`
}

type Fund struct {
	FundingStream struct {
		Description *string `json:"description"`
		ID          *string `json:"id"`
	} `json:"funding_stream"`
	Jurisdiction *string `json:"jurisdiction"`
	Name         *string `json:"name"`
	ShortName    *string `json:"shortName"`
}

type Grant struct {
	Currency     *string  `json:"currency"`
	Fundedamount *float64 `json:"fundedamount"`
	Totalcost    *float64 `json:"totalcost"`
}

type H2020 struct {
	Code        *string `json:"code"`
	Description *string `json:"description"`
}

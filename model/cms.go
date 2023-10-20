package model

// IncommingData - batch incomming data
type IncommingData struct {
	Ruleset     string `json:"ruleset"`
	Version     string `json:"version"`
	ResultTopic string `json:"resultTopic"`
	Data        struct {
		Security struct {
			Isin       string `json:"isin"`
			Type       string `json:"type"`
			SecurityID int    `json:"securityId"`
			Product    struct {
				Product string `json:"product"`
			} `json:"product"`
			Listing struct {
				Isin string `json:"isin"`
			} `json:"listing"`
			PrevDay struct {
				Isin string `json:"isin"`
				Type string `json:"type"`
			} `json:"prevDay"`
		} `json:"security"`
	} `json:"data"`
}

type SecurityRule struct {
	Service string `json:"service"`
	RuleSet string `json:"ruleset"`
	Version string `json:"version"`
	Data    Data   `json:"data"`
}

type Data struct {
	Security    SecurityData   `json:"security"`
	PrevDaySec  prevDaySecData `json:"prevDaySec"`
	Product     ProductData    `json:"product"`
	Listing     ListingData    `json:"listing"`
	ValidValues ValidValueData `json:"validValues"`
}

type SecurityData struct {
	ISIN                  string `json:"isincode"`
	WKN                   string `json:"wkn"`
	Marketplace           string `json:"marketplace"`
	Market                string `json:"market"`
	MarketType            string `json:"marketType"`
	SecurityId            string `json:"securityId"`
	Name                  string `json:"name"`
	Mnemonic              string `json:"mnemonic"`
	SecType               string `json:"secType"`
	SubType               string `json:"subType"`
	HomeMarketIdentCode   string `json:"homeMarketIdentCode"`
	LegalEntityIdentifier string `json:"legalEntityIdentifier"`
	IssuerName            string `json:"issuerName"`
	LiquidForRegmmo       string `json:"liquidForRegmmo"`
	CreationTimestamp     string `json:"creationTimestamp"`
	FirstTradingDay       string `json:"firstTradingDay"`
}

type prevDaySecData struct {
	ISIN                  string `json:"isincode"`
	WKN                   string `json:"wkn"`
	Marketplace           string `json:"marketplace"`
	Market                string `json:"market"`
	MarketType            string `json:"marketType"`
	SecurityId            string `json:"securityId"`
	Name                  string `json:"name"`
	Mnemonic              string `json:"mnemonic"`
	SecType               string `json:"secType"`
	SubType               string `json:"subType"`
	HomeMarketIdentCode   string `json:"homeMarketIdentCode"`
	LegalEntityIdentifier string `json:"legalEntityIdentifier"`
	IssuerName            string `json:"issuerName"`
	LiquidForRegmmo       string `json:"liquidForRegmmo"`
	CreationTimestamp     string `json:"creationTimestamp"`
}

type ProductData struct {
	ISIN                   string `json:"isincode"`
	Marketplace            string `json:"marketplace"`
	Market                 string `json:"market"`
	SecurityId             string `json:"securityId"`
	ProductAssignmentGroup string `json:"productAssignmentGroup"`
}

type ListingData struct {
	ISIN            string `json:"isincode"`
	Marketplace     string `json:"marketplace"`
	Market          string `json:"market"`
	SecurityId      string `json:"securityId"`
	FirstTradingDay string `json:"firstTradingDay"`
}

type ValidValueData struct {
	AttributeName  string `json:"attributeName"`
	AttributeValue string `json:"attributeValue"`
}

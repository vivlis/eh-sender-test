package model

import "time"

// SecurityData - data structure
type SecurityData struct {
	ISINCode            string     `json:"ISINCode"`
	WKN                 string     `json:"WKN"`
	MarketPlace         string     `json:"marketplace"`
	Market              string     `json:"market "`
	SecurityID          string     `json:"securityID"`
	Name                string     `json:"name"`
	Mnemonic            string     `json:"mnemonic"`
	SecType             string     `json:"secType"`
	CashSecuritySubType string     `json:"cashSecuritySubType"`
	WarrantType         string     `json:"warrantType"`
	HomeMarket          string     `json:"homeMarket"`
	IdentifierCode      string     `json:"IdentifierCode"`
	IssuerLegalEntityID string     `json:"issuerLegalEntityId"`
	StrikePrice         string     `json:"strikePrice"`
	Issuer              string     `json:"issuer"`
	Underlying          string     `json:"underlying"`
	IssueDate           string     `json:"issueDate"`
	MaturityDate        string     `json:"maturityDate"`
	Author              string     `json:"author"`
	CreationTimestamp   *time.Time `json:"creationTimestamp"`
	ActionCode          string     `json:"actionCode"`
	WSSMasterID         int        `json:"WSSMasterID"`
	CoverIndicator      string     `json:"coverIndicator"`
	ValidationCodes     []string   `json:"validationCodes"`
}

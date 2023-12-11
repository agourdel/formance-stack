// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/formancehq/formance-sdk-go/pkg/utils"
)

type ConnectorConfigType string

const (
	ConnectorConfigTypeStripeConfig        ConnectorConfigType = "StripeConfig"
	ConnectorConfigTypeDummyPayConfig      ConnectorConfigType = "DummyPayConfig"
	ConnectorConfigTypeWiseConfig          ConnectorConfigType = "WiseConfig"
	ConnectorConfigTypeModulrConfig        ConnectorConfigType = "ModulrConfig"
	ConnectorConfigTypeCurrencyCloudConfig ConnectorConfigType = "CurrencyCloudConfig"
	ConnectorConfigTypeBankingCircleConfig ConnectorConfigType = "BankingCircleConfig"
	ConnectorConfigTypeMangoPayConfig      ConnectorConfigType = "MangoPayConfig"
	ConnectorConfigTypeMoneycorpConfig     ConnectorConfigType = "MoneycorpConfig"
	ConnectorConfigTypeAtlarConfig         ConnectorConfigType = "AtlarConfig"
)

type ConnectorConfig struct {
	StripeConfig        *StripeConfig
	DummyPayConfig      *DummyPayConfig
	WiseConfig          *WiseConfig
	ModulrConfig        *ModulrConfig
	CurrencyCloudConfig *CurrencyCloudConfig
	BankingCircleConfig *BankingCircleConfig
	MangoPayConfig      *MangoPayConfig
	MoneycorpConfig     *MoneycorpConfig
	AtlarConfig         *AtlarConfig

	Type ConnectorConfigType
}

func CreateConnectorConfigStripeConfig(stripeConfig StripeConfig) ConnectorConfig {
	typ := ConnectorConfigTypeStripeConfig

	return ConnectorConfig{
		StripeConfig: &stripeConfig,
		Type:         typ,
	}
}

func CreateConnectorConfigDummyPayConfig(dummyPayConfig DummyPayConfig) ConnectorConfig {
	typ := ConnectorConfigTypeDummyPayConfig

	return ConnectorConfig{
		DummyPayConfig: &dummyPayConfig,
		Type:           typ,
	}
}

func CreateConnectorConfigWiseConfig(wiseConfig WiseConfig) ConnectorConfig {
	typ := ConnectorConfigTypeWiseConfig

	return ConnectorConfig{
		WiseConfig: &wiseConfig,
		Type:       typ,
	}
}

func CreateConnectorConfigModulrConfig(modulrConfig ModulrConfig) ConnectorConfig {
	typ := ConnectorConfigTypeModulrConfig

	return ConnectorConfig{
		ModulrConfig: &modulrConfig,
		Type:         typ,
	}
}

func CreateConnectorConfigCurrencyCloudConfig(currencyCloudConfig CurrencyCloudConfig) ConnectorConfig {
	typ := ConnectorConfigTypeCurrencyCloudConfig

	return ConnectorConfig{
		CurrencyCloudConfig: &currencyCloudConfig,
		Type:                typ,
	}
}

func CreateConnectorConfigBankingCircleConfig(bankingCircleConfig BankingCircleConfig) ConnectorConfig {
	typ := ConnectorConfigTypeBankingCircleConfig

	return ConnectorConfig{
		BankingCircleConfig: &bankingCircleConfig,
		Type:                typ,
	}
}

func CreateConnectorConfigMangoPayConfig(mangoPayConfig MangoPayConfig) ConnectorConfig {
	typ := ConnectorConfigTypeMangoPayConfig

	return ConnectorConfig{
		MangoPayConfig: &mangoPayConfig,
		Type:           typ,
	}
}

func CreateConnectorConfigMoneycorpConfig(moneycorpConfig MoneycorpConfig) ConnectorConfig {
	typ := ConnectorConfigTypeMoneycorpConfig

	return ConnectorConfig{
		MoneycorpConfig: &moneycorpConfig,
		Type:            typ,
	}
}

func CreateConnectorConfigAtlarConfig(atlarConfig AtlarConfig) ConnectorConfig {
	typ := ConnectorConfigTypeAtlarConfig

	return ConnectorConfig{
		AtlarConfig: &atlarConfig,
		Type:        typ,
	}
}

func (u *ConnectorConfig) UnmarshalJSON(data []byte) error {

	wiseConfig := WiseConfig{}
	if err := utils.UnmarshalJSON(data, &wiseConfig, "", true, true); err == nil {
		u.WiseConfig = &wiseConfig
		u.Type = ConnectorConfigTypeWiseConfig
		return nil
	}

	stripeConfig := StripeConfig{}
	if err := utils.UnmarshalJSON(data, &stripeConfig, "", true, true); err == nil {
		u.StripeConfig = &stripeConfig
		u.Type = ConnectorConfigTypeStripeConfig
		return nil
	}

	dummyPayConfig := DummyPayConfig{}
	if err := utils.UnmarshalJSON(data, &dummyPayConfig, "", true, true); err == nil {
		u.DummyPayConfig = &dummyPayConfig
		u.Type = ConnectorConfigTypeDummyPayConfig
		return nil
	}

	modulrConfig := ModulrConfig{}
	if err := utils.UnmarshalJSON(data, &modulrConfig, "", true, true); err == nil {
		u.ModulrConfig = &modulrConfig
		u.Type = ConnectorConfigTypeModulrConfig
		return nil
	}

	currencyCloudConfig := CurrencyCloudConfig{}
	if err := utils.UnmarshalJSON(data, &currencyCloudConfig, "", true, true); err == nil {
		u.CurrencyCloudConfig = &currencyCloudConfig
		u.Type = ConnectorConfigTypeCurrencyCloudConfig
		return nil
	}

	mangoPayConfig := MangoPayConfig{}
	if err := utils.UnmarshalJSON(data, &mangoPayConfig, "", true, true); err == nil {
		u.MangoPayConfig = &mangoPayConfig
		u.Type = ConnectorConfigTypeMangoPayConfig
		return nil
	}

	moneycorpConfig := MoneycorpConfig{}
	if err := utils.UnmarshalJSON(data, &moneycorpConfig, "", true, true); err == nil {
		u.MoneycorpConfig = &moneycorpConfig
		u.Type = ConnectorConfigTypeMoneycorpConfig
		return nil
	}

	atlarConfig := AtlarConfig{}
	if err := utils.UnmarshalJSON(data, &atlarConfig, "", true, true); err == nil {
		u.AtlarConfig = &atlarConfig
		u.Type = ConnectorConfigTypeAtlarConfig
		return nil
	}

	bankingCircleConfig := BankingCircleConfig{}
	if err := utils.UnmarshalJSON(data, &bankingCircleConfig, "", true, true); err == nil {
		u.BankingCircleConfig = &bankingCircleConfig
		u.Type = ConnectorConfigTypeBankingCircleConfig
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConnectorConfig) MarshalJSON() ([]byte, error) {
	if u.StripeConfig != nil {
		return utils.MarshalJSON(u.StripeConfig, "", true)
	}

	if u.DummyPayConfig != nil {
		return utils.MarshalJSON(u.DummyPayConfig, "", true)
	}

	if u.WiseConfig != nil {
		return utils.MarshalJSON(u.WiseConfig, "", true)
	}

	if u.ModulrConfig != nil {
		return utils.MarshalJSON(u.ModulrConfig, "", true)
	}

	if u.CurrencyCloudConfig != nil {
		return utils.MarshalJSON(u.CurrencyCloudConfig, "", true)
	}

	if u.BankingCircleConfig != nil {
		return utils.MarshalJSON(u.BankingCircleConfig, "", true)
	}

	if u.MangoPayConfig != nil {
		return utils.MarshalJSON(u.MangoPayConfig, "", true)
	}

	if u.MoneycorpConfig != nil {
		return utils.MarshalJSON(u.MoneycorpConfig, "", true)
	}

	if u.AtlarConfig != nil {
		return utils.MarshalJSON(u.AtlarConfig, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

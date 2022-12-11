package types

type MerchantAccountInformation struct {
	GloballyUniqueIdentifier EmvTag
	PixKey                   EmvTag
	AditionalData            EmvTag
}

type AditionDataFieldTemplate struct {
	ReferenceLabel EmvTag
}

type PixStaticEmv struct {
	PayloadFormatIndicator     EmvTag
	PointOfInitiationMethod    EmvTag
	MerchantAccountInformation MerchantAccountInformation
	MerchantCategoryCode       EmvTag
	TransactionCurrency        EmvTag
	TransactionAmount          EmvTag
	CountryCode                EmvTag
	MerchantName               EmvTag
	MerchantCity               EmvTag
	AdditionDataFieldTemplate  AditionDataFieldTemplate
	CRC                        EmvTag
}

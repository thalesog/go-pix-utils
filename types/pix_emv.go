package types

var BcbGui = "br.gov.bcb.pix"

var BcbGuiTag = EmvTag{
	Tag:   "00",
	Size:  14,
	Value: BcbGui,
}

type MerchantAccountInformation struct {
	GUI           EmvTag // EL 00
	PixKey        EmvTag // EL 01
	AditionalData EmvTag // EL 02
	Fss           EmvTag // EL 03
	Url           EmvTag // EL 25
}

type AditionDataFieldTemplate struct {
	ReferenceLabel EmvTag
}

type PixEmv struct {
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

type ParsedPixEmv struct {
	Raw      string
	Type     string
	Elements PixEmv
}

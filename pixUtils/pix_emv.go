package pixUtils

var BcbGui = "br.gov.bcb.pix"

var BcbGuiTag = EmvTag{
	Tag:   "00",
	Size:  14,
	Value: BcbGui,
}

type MaiStructure struct {
	GUI         EmvTag // EL 00
	PixKey      EmvTag // EL 01
	Description EmvTag // EL 02
	Fss         EmvTag // EL 03
	Url         EmvTag // EL 25
}

type AdditionalDataStructure struct {
	ReferenceLabel EmvTag
}

type PixEmv struct {
	PayloadFormatIndicator     EmvTag
	MerchantAccountInformation MaiStructure
	MerchantCategoryCode       EmvTag
	TransactionCurrency        EmvTag
	TransactionAmount          EmvTag
	CountryCode                EmvTag
	MerchantName               EmvTag
	MerchantCity               EmvTag
	AdditionDataFieldTemplate  AdditionalDataStructure
	CRC                        EmvTag
}

type PixObject struct {
	Raw      string
	Type     string
	Elements PixEmv
}

var basePixEmv = PixEmv{
	PayloadFormatIndicator: EmvTag{
		Tag:   TagPayloadFormatIndicator,
		Size:  2,
		Value: "01",
	},
	MerchantAccountInformation: MaiStructure{
		GUI: BcbGuiTag,
	},
	MerchantCategoryCode: EmvTag{
		Tag:   TagMerchantCategoryCode,
		Size:  4,
		Value: "0000",
	},
	TransactionCurrency: EmvTag{
		Tag:   TagTransactionCurrency,
		Size:  3,
		Value: "986",
	},
	CountryCode: EmvTag{
		Tag:   TagCountryCode,
		Size:  2,
		Value: "BR",
	},
	AdditionDataFieldTemplate: AdditionalDataStructure{
		ReferenceLabel: EmvTag{
			Tag:   TagReferenceLabel,
			Size:  3,
			Value: "***",
		},
	},
	CRC: EmvTag{
		Tag:   TagCRC,
		Size:  4,
		Value: "",
	},
}

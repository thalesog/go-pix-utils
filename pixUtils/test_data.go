package pixUtils

const staticEmv = "00020126510014br.gov.bcb.pix0115thalesog@me.com0210Pedido 123520400005303986540510.005802BR5914THALES OGLIARI6013SAO MIGUEL DO62070503***63046CDA"

const staticEmvWrongCrc = "00020126510014br.gov.bcb.pix0115thalesog@me.com0210Pedido 123520400005303986540510.005802BR5914THALES OGLIARI6013SAO MIGUEL DO62070503***63046CDX"

const dynamicEmv = "00020126740014br.gov.bcb.pix2552payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e875204000053039865802BR5914THALES OGLIARI6013SAO MIGUEL DO62070503***63047E8D"

var dynamicElements = PixEmv{
	PayloadFormatIndicator: EmvTag{
		Size:  2,
		Value: "01",
		Tag:   "00",
	},
	MerchantAccountInformation: MaiStructure{
		GUI: EmvTag{
			Size:  14,
			Value: BcbGui,
			Tag:   "00",
		},
		Url: EmvTag{
			Size:  52,
			Value: "payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e87",
			Tag:   "25",
		},
	},
	MerchantCategoryCode: EmvTag{
		Size:  4,
		Value: "0000",
		Tag:   "52",
	},
	TransactionCurrency: EmvTag{
		Size:  3,
		Value: "986",
		Tag:   "53",
	},
	CountryCode: EmvTag{
		Size:  2,
		Value: "BR",
		Tag:   "58",
	},
	MerchantName: EmvTag{
		Size:  14,
		Value: "THALES OGLIARI",
		Tag:   "59",
	},
	MerchantCity: EmvTag{
		Size:  13,
		Value: "SAO MIGUEL DO",
		Tag:   "60",
	},
	AdditionDataFieldTemplate: AdditionalDataStructure{
		ReferenceLabel: EmvTag{
			Size:  3,
			Value: "***",
			Tag:   "05",
		},
	},
	CRC: EmvTag{
		Size:  4,
		Value: "7E8D",
		Tag:   "63",
	},
}

var staticElements = PixEmv{
	PayloadFormatIndicator: EmvTag{
		Size:  2,
		Value: "01",
		Tag:   TagPayloadFormatIndicator,
	},
	MerchantAccountInformation: MaiStructure{
		GUI: EmvTag{
			Size:  14,
			Value: BcbGui,
			Tag:   TagGUI,
		},
		PixKey: EmvTag{
			Size:  15,
			Value: "thalesog@me.com",
			Tag:   TagPixKey,
		},
		Description: EmvTag{
			Size:  10,
			Value: "Pedido 123",
			Tag:   TagDescription,
		},
	},
	MerchantCategoryCode: EmvTag{
		Size:  4,
		Value: "0000",
		Tag:   TagMerchantCategoryCode,
	},
	TransactionCurrency: EmvTag{
		Size:  3,
		Value: "986",
		Tag:   TagTransactionCurrency,
	},
	TransactionAmount: EmvTag{
		Size:  5,
		Value: "10.00",
		Tag:   TagTransactionAmount,
	},
	CountryCode: EmvTag{
		Size:  2,
		Value: "BR",
		Tag:   TagCountryCode,
	},
	MerchantName: EmvTag{
		Size:  14,
		Value: "THALES OGLIARI",
		Tag:   TagMerchantName,
	},
	MerchantCity: EmvTag{
		Size:  13,
		Value: "SAO MIGUEL DO",
		Tag:   TagMerchantCity,
	},

	AdditionDataFieldTemplate: AdditionalDataStructure{
		ReferenceLabel: EmvTag{
			Size:  3,
			Value: "***",
			Tag:   TagReferenceLabel,
		},
	},
	CRC: EmvTag{
		Size:  4,
		Value: "6CDA",
		Tag:   TagCRC,
	},
}

package main

import "github.com/thalesog/go-pix-utils/types"

const staticEmv = "00020126510014br.gov.bcb.pix0115thalesog@me.com0210Pedido 123520400005303986540510.005802BR5914THALES OGLIARI6015SAO MIGUEL DO O62070503***63040DC6"

const staticEmvWrongCrc = "00020126510014br.gov.bcb.pix0115thalesog@me.com0210Pedido 123520400005303986540510.005802BR5914THALES OGLIARI6015SAO MIGUEL DO O62070503***63040DCX"

const dynamicEmv = "00020126740014br.gov.bcb.pix2552payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e875204000053039865802BR5914THALES OGLIARI6015SAO MIGUEL DO O62070503***630442C5"

var dynamicElements = types.PixEmv{
	PayloadFormatIndicator: types.EmvTag{
		Size:  2,
		Value: "01",
		Tag:   "00",
	},
	MerchantAccountInformation: types.MerchantAccountInformation{
		GUI: types.EmvTag{
			Size:  14,
			Value: types.BcbGui,
			Tag:   "00",
		},
		Url: types.EmvTag{
			Size:  52,
			Value: "payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e87",
			Tag:   "25",
		},
	},
	MerchantCategoryCode: types.EmvTag{
		Size:  4,
		Value: "0000",
		Tag:   "52",
	},
	TransactionCurrency: types.EmvTag{
		Size:  3,
		Value: "986",
		Tag:   "53",
	},
	CountryCode: types.EmvTag{
		Size:  2,
		Value: "BR",
		Tag:   "58",
	},
	MerchantName: types.EmvTag{
		Size:  14,
		Value: "THALES OGLIARI",
		Tag:   "59",
	},
	MerchantCity: types.EmvTag{
		Size:  15,
		Value: "SAO MIGUEL DO O",
		Tag:   "60",
	},
	AdditionDataFieldTemplate: types.AditionDataFieldTemplate{
		ReferenceLabel: types.EmvTag{
			Size:  3,
			Value: "***",
			Tag:   "05",
		},
	},
	CRC: types.EmvTag{
		Size:  4,
		Value: "42C5",
		Tag:   "63",
	},
}

var staticElements = types.PixEmv{
	PayloadFormatIndicator: types.EmvTag{
		Size:  2,
		Value: "01",
		Tag:   "00",
	},
	MerchantAccountInformation: types.MerchantAccountInformation{
		GUI: types.EmvTag{
			Size:  14,
			Value: types.BcbGui,
			Tag:   "00",
		},
		PixKey: types.EmvTag{
			Size:  15,
			Value: "thalesog@me.com",
			Tag:   "01",
		},
		AditionalData: types.EmvTag{
			Size:  10,
			Value: "Pedido 123",
			Tag:   "02",
		},
	},
	MerchantCategoryCode: types.EmvTag{
		Size:  4,
		Value: "0000",
		Tag:   "52",
	},
	TransactionCurrency: types.EmvTag{
		Size:  3,
		Value: "986",
		Tag:   "53",
	},
	TransactionAmount: types.EmvTag{
		Size:  5,
		Value: "10.00",
		Tag:   "54",
	},
	CountryCode: types.EmvTag{
		Size:  2,
		Value: "BR",
		Tag:   "58",
	},
	MerchantName: types.EmvTag{
		Size:  14,
		Value: "THALES OGLIARI",
		Tag:   "59",
	},
	MerchantCity: types.EmvTag{
		Size:  15,
		Value: "SAO MIGUEL DO O",
		Tag:   "60",
	},

	AdditionDataFieldTemplate: types.AditionDataFieldTemplate{
		ReferenceLabel: types.EmvTag{
			Size:  3,
			Value: "***",
			Tag:   "05",
		},
	},
	CRC: types.EmvTag{
		Size:  4,
		Value: "0DC6",
		Tag:   "63",
	},
}

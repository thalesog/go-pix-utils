package emv

import "github.com/thalesog/go-pix-utils/types"

func fillPixEmv(parsedTag types.EmvTag, parsedEmv types.PixEmv) types.PixEmv {

	switch parsedTag.Tag {
	case "00":
		parsedEmv.PayloadFormatIndicator = parsedTag
	case "26":
		parsedEmv.MerchantAccountInformation = parseMAI(parsedTag, types.MerchantAccountInformation{})
	case "52":
		parsedEmv.MerchantCategoryCode = parsedTag
	case "53":
		parsedEmv.TransactionCurrency = parsedTag
	case "54":
		parsedEmv.TransactionAmount = parsedTag
	case "58":
		parsedEmv.CountryCode = parsedTag
	case "59":
		parsedEmv.MerchantName = parsedTag
	case "60":
		parsedEmv.MerchantCity = parsedTag
	case "62":
		parsedEmv.AdditionDataFieldTemplate = parseTxId(parsedTag, types.AditionDataFieldTemplate{})
	case "63":
		parsedEmv.CRC = parsedTag
	}

	return parsedEmv
}

func fillMAI(parsedTag types.EmvTag, parsedMai types.MerchantAccountInformation) types.MerchantAccountInformation {

	switch parsedTag.Tag {
	case "00":
		parsedMai.GUI = parsedTag
	case "01":
		parsedMai.PixKey = parsedTag
	case "02":
		parsedMai.AditionalData = parsedTag
	case "03":
		parsedMai.Fss = parsedTag
	case "25":
		parsedMai.Url = parsedTag
	}

	return parsedMai
}

func fillTxId(parsedTag types.EmvTag, parsedTxId types.AditionDataFieldTemplate) types.AditionDataFieldTemplate {

	switch parsedTag.Tag {
	case "05":
		parsedTxId.ReferenceLabel = parsedTag
	}

	return parsedTxId
}

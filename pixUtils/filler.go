package pixUtils

func fillPixEmv(parsedTag EmvTag, parsedEmv PixEmv) PixEmv {

	switch parsedTag.Tag {
	case TagPayloadFormatIndicator:
		parsedEmv.PayloadFormatIndicator = parsedTag
	case TagMerchantAccountInformation:
		parsedEmv.MerchantAccountInformation = parseMAI(parsedTag, MaiStructure{})
	case TagMerchantCategoryCode:
		parsedEmv.MerchantCategoryCode = parsedTag
	case TagTransactionCurrency:
		parsedEmv.TransactionCurrency = parsedTag
	case TagTransactionAmount:
		parsedEmv.TransactionAmount = parsedTag
	case TagCountryCode:
		parsedEmv.CountryCode = parsedTag
	case TagMerchantName:
		parsedEmv.MerchantName = parsedTag
	case TagMerchantCity:
		parsedEmv.MerchantCity = parsedTag
	case TagAdditionDataFieldTemplate:
		parsedEmv.AdditionDataFieldTemplate = parseTxId(parsedTag, AdditionalDataStructure{})
	case TagCRC:
		parsedEmv.CRC = parsedTag
	}

	return parsedEmv
}

func fillMAI(parsedTag EmvTag, parsedMai MaiStructure) MaiStructure {

	switch parsedTag.Tag {
	case TagGUI:
		parsedMai.GUI = parsedTag
	case TagPixKey:
		parsedMai.PixKey = parsedTag
	case TagDescription:
		parsedMai.Description = parsedTag
	case TagFss:
		parsedMai.Fss = parsedTag
	case TagUrl:
		parsedMai.Url = parsedTag
	}

	return parsedMai
}

func fillTxId(parsedTag EmvTag, parsedTxId AdditionalDataStructure) AdditionalDataStructure {

	switch parsedTag.Tag {
	case TagReferenceLabel:
		parsedTxId.ReferenceLabel = parsedTag
	}

	return parsedTxId
}

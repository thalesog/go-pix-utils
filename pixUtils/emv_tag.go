package pixUtils

type EmvTag struct {
	Size  int
	Value string
	Tag   string
}

const (
	TagPayloadFormatIndicator     = "00"
	TagMerchantAccountInformation = "26"
	TagMerchantCategoryCode       = "52"
	TagTransactionCurrency        = "53"
	TagTransactionAmount          = "54"
	TagCountryCode                = "58"
	TagMerchantName               = "59"
	TagMerchantCity               = "60"
	TagAdditionDataFieldTemplate  = "62"
	TagCRC                        = "63"

	/*
	   Merchant Account Information Tags
	*/
	TagGUI         = "00"
	TagPixKey      = "01"
	TagDescription = "02"
	TagFss         = "03"
	TagUrl         = "25"

	/*
	   Adittional Data Field Template Tags
	*/
	TagReferenceLabel = "05"
)

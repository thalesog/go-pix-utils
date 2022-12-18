package emv

import "github.com/thalesog/go-pix-utils/types"

func getPixType(parsedEmv types.PixEmv) string {
	if parsedEmv.MerchantAccountInformation.GUI.Value != types.BcbGui {
		return "invalid"
	}

	if parsedEmv.MerchantAccountInformation.PixKey.Size > 0 &&
		parsedEmv.MerchantAccountInformation.PixKey.Size < 77 &&
		parsedEmv.MerchantAccountInformation.Url.Size == 0 {
		return "static"
	}

	if parsedEmv.MerchantAccountInformation.Url.Size > 0 &&
		parsedEmv.MerchantAccountInformation.Url.Size < 77 &&
		parsedEmv.MerchantAccountInformation.PixKey.Size == 0 {
		return "dynamic"
	}

	return "invalid"
}

package emv

import (
	"log"
	"strconv"

	"github.com/thalesog/go-pix-utils/types"
)

func ParseEmvTag(emv string) (parsedTag types.EmvTag, err error) {
	tag := emv[0:2]
	size := emv[2:4]
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		log.Printf("Error parsing size: %s", err)
		return parsedTag, err
	}

	value := emv[4 : 4+sizeInt]

	return types.EmvTag{
		Size:  sizeInt,
		Value: value,
		Tag:   tag,
	}, nil
}

func ParseMaiEmv(maiTag types.EmvTag) types.MerchantAccountInformation {
	var parsedMai types.MerchantAccountInformation

	for true {
		parsedTag, err := ParseEmvTag(maiTag.Value)
		if err != nil {
			log.Printf("Error parsing tag: %s", err)
			break
		}
		log.Printf("Tag: %s, Size: %d, Value: %s", parsedTag.Tag, parsedTag.Size, parsedTag.Value)
		maiTag.Value = maiTag.Value[parsedTag.Size+4:]

		switch parsedTag.Tag {
		case "00":
			parsedMai.GloballyUniqueIdentifier = parsedTag
		case "01":
			parsedMai.PixKey = parsedTag
		case "02":
			parsedMai.AditionalData = parsedTag
		}

		if len(maiTag.Value) == 0 {
			break
		}

	}

	return parsedMai
}

func ParseTxId(emv string) (parsedEmv types.AditionDataFieldTemplate) {
	for true {
		parsedTag, err := ParseEmvTag(emv)
		if err != nil {
			log.Printf("Error parsing tag: %s", err)
			break
		}
		log.Printf("Tag: %s, Size: %d, Value: %s", parsedTag.Tag, parsedTag.Size, parsedTag.Value)
		emv = emv[parsedTag.Size+4:]

		switch parsedTag.Tag {
		case "05":
			parsedEmv.ReferenceLabel = parsedTag
		}

		if len(emv) == 0 {
			break
		}

	}

	return parsedEmv
}

func ParsePixStatic(emv string) (parsedEmv types.PixStaticEmv) {
	for true {
		parsedTag, err := ParseEmvTag(emv)
		if err != nil {
			log.Printf("Error parsing tag: %s", err)
			break
		}
		log.Printf("Tag: %s, Size: %d, Value: %s", parsedTag.Tag, parsedTag.Size, parsedTag.Value)
		emv = emv[parsedTag.Size+4:]

		switch parsedTag.Tag {
		case "00":
			parsedEmv.PayloadFormatIndicator = parsedTag
		case "26":
			parsedEmv.MerchantAccountInformation = ParseMaiEmv(parsedTag)
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
			parsedEmv.AdditionDataFieldTemplate = ParseTxId(parsedTag.Value)
		case "63":
			parsedEmv.CRC = parsedTag
		default:
			log.Printf("Tag not found: %s", parsedTag.Tag)
			break
		}

		if len(emv) == 0 {
			log.Println("No more tags")
			break
		}

	}

	return parsedEmv
}

package main

import (
	"errors"
	"github.com/thalesog/go-pix-utils/types"
	"log"
	"strconv"
)

func parseEmvTag(emv string) (parsedTag types.EmvTag, err error) {
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

func parsePixEmv(emv string, parsedCode types.PixEmv) types.PixEmv {
	parsedTag, err := parseEmvTag(emv)
	parsedCode = fillPixEmv(parsedTag, parsedCode)

	if err != nil {
		log.Printf("Error parsing tag: %s", err)
		return parsedCode
	}

	if len(emv)-parsedTag.Size-4 > 0 {
		emv = emv[parsedTag.Size+4:]
		return parsePixEmv(emv, parsedCode)
	}

	return parsedCode
}

func parseMAI(maiTag types.EmvTag, parsedMAI types.MerchantAccountInformation) types.MerchantAccountInformation {

	parsedTag, err := parseEmvTag(maiTag.Value)
	parsedMAI = fillMAI(parsedTag, parsedMAI)

	if err != nil {
		log.Printf("Error parsing tag: %s", err)
		return parsedMAI
	}

	if len(maiTag.Value)-parsedTag.Size-4 > 0 {
		maiTag.Value = maiTag.Value[parsedTag.Size+4:]
		return parseMAI(maiTag, parsedMAI)
	}

	return parsedMAI
}

func parseTxId(maiTag types.EmvTag, parsedTxId types.AditionDataFieldTemplate) types.AditionDataFieldTemplate {

	parsedTag, err := parseEmvTag(maiTag.Value)
	parsedTxId = fillTxId(parsedTag, parsedTxId)

	if err != nil {
		log.Printf("Error parsing tag: %s", err)
		return parsedTxId
	}

	if len(maiTag.Value)-parsedTag.Size-4 > 0 {
		maiTag.Value = maiTag.Value[parsedTag.Size+4:]
		return parseTxId(maiTag, parsedTxId)
	}

	return parsedTxId
}

func ParsePix(emvString string) (parsedPix types.ParsedPixEmv, err error) {
	parsedPix.Raw = emvString
	parsedPix.Elements = parsePixEmv(emvString, types.PixEmv{})
	parsedPix.Type = getPixType(parsedPix.Elements)

	if parsedPix.Type == "invalid" {
		err = errors.New("invalid pix")
	}

	if validateCRC(parsedPix.Raw) == false {
		err = errors.New("invalid crc")
	}

	return
}

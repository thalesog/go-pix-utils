package pixUtils

import (
	"fmt"
	"strings"
)

func createEmvTag(tag string, value string) string {
	if tag == TagCRC && value == "" {
		return tag + "04"
	}

	valueLength := len(value)

	if valueLength == 0 || valueLength > 99 {
		return ""
	}

	length := fmt.Sprintf("%02d", valueLength)

	return strings.Join([]string{tag, length, value}, "")
}

func createEmv(elements PixEmv) string {
	emptySeparator := ""

	maiTag := strings.Join([]string{
		createEmvTag(TagGUI, elements.MerchantAccountInformation.GUI.Value),
		createEmvTag(TagPixKey, elements.MerchantAccountInformation.PixKey.Value),
		createEmvTag(TagUrl, elements.MerchantAccountInformation.Url.Value),
		createEmvTag(TagDescription, elements.MerchantAccountInformation.Description.Value),
		createEmvTag(TagFss, elements.MerchantAccountInformation.Fss.Value),
	}, emptySeparator)

	additionalDataField := strings.Join([]string{
		createEmvTag(TagReferenceLabel, elements.AdditionDataFieldTemplate.ReferenceLabel.Value),
	}, emptySeparator)

	generatedEmv := strings.Join([]string{
		createEmvTag(TagPayloadFormatIndicator, elements.PayloadFormatIndicator.Value),
		createEmvTag(TagMerchantAccountInformation, maiTag),
		createEmvTag(TagMerchantCategoryCode, elements.MerchantCategoryCode.Value),
		createEmvTag(TagTransactionCurrency, elements.TransactionCurrency.Value),
		createEmvTag(TagTransactionAmount, elements.TransactionAmount.Value),
		createEmvTag(TagCountryCode, elements.CountryCode.Value),
		createEmvTag(TagMerchantName, elements.MerchantName.Value),
		createEmvTag(TagMerchantCity, elements.MerchantCity.Value),
		createEmvTag(TagAdditionDataFieldTemplate, additionalDataField),
		createEmvTag(TagCRC, elements.CRC.Value),
	}, emptySeparator)

	return generatedEmv
}

type CreateStaticPixParams struct {
	PixKey            string
	MerchantName      string
	MerchantCity      string
	TransactionAmount float64
	ReferenceLabel    string
	AditionalData     string
}

func CreateStaticPix(params CreateStaticPixParams) PixObject {

	elements := basePixEmv

	elements.MerchantAccountInformation.PixKey = EmvTag{
		Tag:   TagPixKey,
		Value: params.PixKey,
		Size:  len(params.PixKey),
	}

	elements.MerchantAccountInformation.Description = EmvTag{
		Tag:   TagDescription,
		Value: params.AditionalData,
		Size:  len(params.AditionalData),
	}

	nameLength := len(params.MerchantName)
	if nameLength > 25 {
		nameLength = 25
	}
	normalizedName, _ := RemoveAccents(params.MerchantName[:nameLength])
	normalizedName = strings.ToUpper(normalizedName)
	elements.MerchantName = EmvTag{
		Tag:   TagMerchantName,
		Value: normalizedName,
		Size:  len(normalizedName),
	}

	cityLength := len(params.MerchantName)
	if cityLength > 15 {
		cityLength = 15
	}
	normalizedCity, _ := RemoveAccents(params.MerchantCity[:cityLength])
	normalizedCity = strings.ToUpper(normalizedCity)
	elements.MerchantCity = EmvTag{
		Tag:   TagMerchantCity,
		Value: normalizedCity,
		Size:  len(normalizedCity),
	}

	elements.TransactionAmount = EmvTag{
		Tag:   TagTransactionAmount,
		Value: fmt.Sprintf("%.2f", params.TransactionAmount),
		Size:  len(fmt.Sprintf("%.2f", params.TransactionAmount)),
	}

	if params.ReferenceLabel != "" {
		elements.AdditionDataFieldTemplate.ReferenceLabel = EmvTag{
			Tag:   TagReferenceLabel,
			Value: params.ReferenceLabel,
			Size:  len(params.ReferenceLabel),
		}
	}

	generatedEmv := createEmv(elements)
	elements.CRC.Value = calculateCRC(generatedEmv)

	return PixObject{
		Raw:      generatedEmv + elements.CRC.Value,
		Elements: elements,
		Type:     "static",
	}
}

type CreateDynamicPixParams struct {
	MerchantName string
	MerchantCity string
	Url          string
}

func CreateDynamicPix(params CreateDynamicPixParams) PixObject {

	elements := basePixEmv

	elements.MerchantAccountInformation.Url = EmvTag{
		Tag:   TagUrl,
		Value: params.Url,
		Size:  len(params.Url),
	}

	nameLength := len(params.MerchantName)
	if nameLength > 25 {
		nameLength = 25
	}
	normalizedName, _ := RemoveAccents(params.MerchantName[:nameLength])
	normalizedName = strings.ToUpper(normalizedName)
	elements.MerchantName = EmvTag{
		Tag:   TagMerchantName,
		Value: normalizedName,
		Size:  len(normalizedName),
	}

	cityLength := len(params.MerchantCity)
	if cityLength > 15 {
		cityLength = 15
	}
	normalizedCity, _ := RemoveAccents(params.MerchantCity[:cityLength])
	normalizedCity = strings.ToUpper(normalizedCity)
	elements.MerchantCity = EmvTag{
		Tag:   TagMerchantCity,
		Value: normalizedCity,
		Size:  len(normalizedCity),
	}

	generatedEmv := createEmv(elements)
	elements.CRC.Value = calculateCRC(generatedEmv)

	return PixObject{
		Raw:      generatedEmv + elements.CRC.Value,
		Elements: elements,
		Type:     "dynamic",
	}
}

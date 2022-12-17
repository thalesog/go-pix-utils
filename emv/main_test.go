package emv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalesog/go-pix-utils/emv"
	"github.com/thalesog/go-pix-utils/types"
)

const StaticEmv = "00020126540014br.gov.bcb.pix0119nubank@thalesog.com0209Descricao520400005303986540115802BR5914Thales Ogliari6006Cidade62070503***630441E8"

const StaticEmvWithExtraObjects = "00020126540014br.gov.bcb.pix0119nubank@thalesog.com0209Descricao520400005303986540115802BR5914Thales Ogliari6006Cidade62070503***630441E8640441E8650441E8"

func TestParseEmvTag(t *testing.T) {
	var emvTag types.EmvTag

	emvTag, err := emv.ParseEmvTag("000201")
	assert.Nil(t, err)
	assert.Equal(t, emvTag, types.EmvTag{
		Size:  2,
		Value: "01",
		Tag:   "00",
	})

	emvTag, err = emv.ParseEmvTag("0119nubank@thalesog.com")
	assert.Nil(t, err)
	assert.Equal(t, emvTag, types.EmvTag{
		Size:  19,
		Value: "nubank@thalesog.com",
		Tag:   "01",
	})
}

func TestParsePixStatic(t *testing.T) {
	testingEmv := emv.ParsePixStatic(StaticEmv)
	resultEmv := types.PixStaticEmv{
		PayloadFormatIndicator: types.EmvTag{
			Size:  2,
			Value: "01",
			Tag:   "00",
		},
		MerchantAccountInformation: types.MerchantAccountInformation{
			GloballyUniqueIdentifier: types.EmvTag{
				Size:  14,
				Value: "br.gov.bcb.pix",
				Tag:   "00",
			},
			PixKey: types.EmvTag{
				Size:  19,
				Value: "nubank@thalesog.com",
				Tag:   "01",
			},
			AditionalData: types.EmvTag{
				Size:  9,
				Value: "Descricao",
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
			Size:  1,
			Value: "1",
			Tag:   "54",
		},
		CountryCode: types.EmvTag{
			Size:  2,
			Value: "BR",
			Tag:   "58",
		},
		MerchantName: types.EmvTag{
			Size:  14,
			Value: "Thales Ogliari",
			Tag:   "59",
		},
		MerchantCity: types.EmvTag{
			Size:  6,
			Value: "Cidade",
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
			Value: "41E8",
			Tag:   "63",
		},
	}

	assert.Equal(t, testingEmv, resultEmv)
}

func TestParsePixStaticWithExtraObjects(t *testing.T) {
	testingEmv := emv.ParsePixStatic(StaticEmvWithExtraObjects)
	resultEmv := types.PixStaticEmv{
		PayloadFormatIndicator: types.EmvTag{
			Size:  2,
			Value: "01",
			Tag:   "00",
		},
		MerchantAccountInformation: types.MerchantAccountInformation{
			GloballyUniqueIdentifier: types.EmvTag{
				Size:  14,
				Value: "br.gov.bcb.pix",
				Tag:   "00",
			},
			PixKey: types.EmvTag{
				Size:  19,
				Value: "nubank@thalesog.com",
				Tag:   "01",
			},
			AditionalData: types.EmvTag{
				Size:  9,
				Value: "Descricao",
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
			Size:  1,
			Value: "1",
			Tag:   "54",
		},
		CountryCode: types.EmvTag{
			Size:  2,
			Value: "BR",
			Tag:   "58",
		},
		MerchantName: types.EmvTag{
			Size:  14,
			Value: "Thales Ogliari",
			Tag:   "59",
		},
		MerchantCity: types.EmvTag{
			Size:  6,
			Value: "Cidade",
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
			Value: "41E8",
			Tag:   "63",
		},
	}

	assert.Equal(t, testingEmv, resultEmv)
}

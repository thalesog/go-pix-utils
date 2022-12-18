package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/thalesog/go-pix-utils/types"
	"testing"
)

func TestParsePixStatic(t *testing.T) {
	testingEmv, err := ParsePix(staticEmv)

	assert.Nil(t, err)
	assert.Equal(t, staticElements, testingEmv.Elements)
	assert.Equal(t, staticEmv, testingEmv.Raw)
	assert.Equal(t, "static", testingEmv.Type)
}

func TestParsePixDynamic(t *testing.T) {
	testingEmv, err := ParsePix(dynamicEmv)

	assert.Nil(t, err)
	assert.Equal(t, dynamicElements, testingEmv.Elements)
	assert.Equal(t, dynamicEmv, testingEmv.Raw)
	assert.Equal(t, "dynamic", testingEmv.Type)
}

func TestDetectInvalidCrc(t *testing.T) {
	_, err := ParsePix(staticEmvWrongCrc)
	assert.Errorf(t, err, "invalid crc")
}

func TestParseEmvTag(t *testing.T) {
	var emvTag types.EmvTag

	emvTag, err := parseEmvTag("000201")
	assert.Nil(t, err)
	assert.Equal(t, emvTag, types.EmvTag{
		Size:  2,
		Value: "01",
		Tag:   "00",
	})

	emvTag, err = parseEmvTag("0115thalesog@me.com")
	assert.Nil(t, err)
	assert.Equal(t, emvTag, types.EmvTag{
		Size:  15,
		Value: "thalesog@me.com",
		Tag:   "01",
	})
}

func TestParseMaiStatic(t *testing.T) {
	parsedMai := parseMAI(types.EmvTag{
		Size:  37,
		Value: "0014br.gov.bcb.pix0115thalesog@me.com",
		Tag:   "26",
	}, types.MerchantAccountInformation{})

	assert.Equal(t, parsedMai, types.MerchantAccountInformation{
		GUI: types.BcbGuiTag,
		PixKey: types.EmvTag{
			Size:  15,
			Value: "thalesog@me.com",
			Tag:   "01",
		},
	})
}

func TestParseMaiDynamic(t *testing.T) {

	parsedMai := parseMAI(types.EmvTag{
		Size:  37,
		Value: "0014br.gov.bcb.pix2552payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e87",
		Tag:   "26",
	}, types.MerchantAccountInformation{})

	assert.Equal(t, parsedMai, types.MerchantAccountInformation{
		GUI: types.BcbGuiTag,
		Url: types.EmvTag{
			Size:  52,
			Value: "payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e87",
			Tag:   "25",
		},
	})
}

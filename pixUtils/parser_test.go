package pixUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	var emvTag EmvTag

	emvTag, err := parseEmvTag("000201")
	assert.Nil(t, err)
	assert.Equal(t, emvTag, EmvTag{
		Size:  2,
		Value: "01",
		Tag:   "00",
	})

	emvTag, err = parseEmvTag("0115thalesog@me.com")
	assert.Nil(t, err)
	assert.Equal(t, emvTag, EmvTag{
		Size:  15,
		Value: "thalesog@me.com",
		Tag:   TagPixKey,
	})
}

func TestParseMaiStatic(t *testing.T) {
	parsedMai := parseMAI(EmvTag{
		Size:  37,
		Value: "0014br.gov.bcb.pix0115thalesog@me.com",
		Tag:   TagMerchantAccountInformation,
	}, MaiStructure{})

	assert.Equal(t, parsedMai, MaiStructure{
		GUI: BcbGuiTag,
		PixKey: EmvTag{
			Size:  15,
			Value: "thalesog@me.com",
			Tag:   TagPixKey,
		},
	})
}

func TestParseMaiDynamic(t *testing.T) {

	parsedMai := parseMAI(EmvTag{
		Size:  37,
		Value: "0014br.gov.bcb.pix2552payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e87",
		Tag:   TagMerchantAccountInformation,
	}, MaiStructure{})

	assert.Equal(t, parsedMai, MaiStructure{
		GUI: BcbGuiTag,
		Url: EmvTag{
			Size:  52,
			Value: "payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e87",
			Tag:   TagUrl,
		},
	})
}

package pixUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmvTag(t *testing.T) {
	const tag = TagPixKey
	const value = "thalesog@me.com"

	emvTag := createEmvTag(tag, value)

	assert.Equal(t, "0115thalesog@me.com", emvTag)
}

func TestCreateStaticPixEmv(t *testing.T) {
	pixEmv := createEmv(staticElements)

	assert.Equal(t, staticEmv, pixEmv)
}

func TestCreateDynamicPixEmv(t *testing.T) {
	pixEmv := createEmv(dynamicElements)

	assert.Equal(t, dynamicEmv, pixEmv)
}

func TestCreateStaticPix(t *testing.T) {
	createdPix := CreateStaticPix(CreateStaticPixParams{
		MerchantName:      "Thales Ogliari",
		MerchantCity:      "São Miguel do Oeste",
		PixKey:            "thalesog@me.com",
		TransactionAmount: 10.00,
		AditionalData:     "Pedido 123",
	})
	assert.Equal(t, PixObject{
		Type:     "static",
		Elements: staticElements,
		Raw:      staticEmv,
	}, createdPix)
}

func TestCreateDynamicPix(t *testing.T) {
	createdPix := CreateDynamicPix(CreateDynamicPixParams{
		MerchantName: "Thales Ogliari",
		MerchantCity: "São Miguel do Oeste",
		Url:          "payload.psp.com/3ec9d2f9-5f03-4e0e-820d-63a81e769e87",
	})

	assert.Equal(t, PixObject{
		Type:     "dynamic",
		Elements: dynamicElements,
		Raw:      dynamicEmv,
	}, createdPix)
}

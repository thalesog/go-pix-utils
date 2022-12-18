package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCRC(t *testing.T) {
	const emvWithoutCRC = "00020126370014br.gov.bcb.pix0115thalesog@me.com5204000053039865402105802BR5914THALES OGLIARI6015SAO MIGUEL DO O62070503***6304"
	const emvCrc = "4C2E"

	crc := calculateCRC(emvWithoutCRC)

	assert.Equal(t, emvCrc, crc)
}

func TestValidateWrongCRC(t *testing.T) {
	const staticEmvWrongCrc = "00020126510014br.gov.bcb.pix0115thalesog@me.com0210Pedido 123520400005303986540510.005802BR5914THALES OGLIARI6015SAO MIGUEL DO O62070503***63040DCX"

	assert.False(t, validateCRC(staticEmvWrongCrc))
}

func TestValidateCorrectCRC(t *testing.T) {
	const staticEmv = "00020126510014br.gov.bcb.pix0115thalesog@me.com0210Pedido 123520400005303986540510.005802BR5914THALES OGLIARI6015SAO MIGUEL DO O62070503***63040DC6"

	assert.True(t, validateCRC(staticEmv))
}

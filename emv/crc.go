package emv

import (
	"fmt"
	"github.com/sigurn/crc16"
	"regexp"
	"strings"
)

var crcTable = crc16.MakeTable(crc16.CRC16_CCITT_FALSE)

func CalculateCRC(value string) string {
	crcInt := crc16.Checksum([]byte(value), crcTable)

	return strings.ToUpper(fmt.Sprintf("%04X", crcInt))
}

func ValidateCRC(value string) bool {
	match, err := regexp.MatchString(`6304\w{4}$`, value)

	if err != nil || match == false {
		return false
	}

	crc := value[len(value)-4:]
	value = value[:len(value)-4]
	calculatedCRC := CalculateCRC(value)
	return crc == calculatedCRC

}

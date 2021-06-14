package tbcd

import (
	"fmt"
	"testing"

)

func TestEncodeToTBCD(t *testing.T) {
	encmsg, err1 := EncodeToTBCD("12345")
	dedmsg, err2 := DecodeToTBCD(encmsg)

	if err1 == nil && err2 == nil {
		fmt.Println("Encoded Message", encmsg)
		fmt.Println("Decoded Message", dedmsg)
	} else { t.Error("Encoded and Decoded message not matched")}
}

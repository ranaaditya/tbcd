# tbcd
Basic Implementation of Telephony BCD in pure GoLang

## Example
```
package main

import tbcd "github.com/ranaaditya/tbcd"

func main() {

var message string = "0123456789*#abc"
encodedmessage, err := tbcd.EncodeToTBCD(message)

if err == nil {
  fmt.Println("Encoded Message in TBCD format is :", encodedmessage)
} else {  fmt.Println(err) }

originalmessage, err1 := tbcd.DecodeToTBCD(encodedmessage)

if err1 == nil {
  fmt.Println("Decoded message is :", originalmessage)
} else { fmt.Println(err1) }

}
```

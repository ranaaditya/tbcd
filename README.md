# tbcd
Basic Implementation of Telephony BCD in pure GoLang

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/open-source.svg)](https://forthebadge.com)

### Building from source
Requires Go to be [installed](https://golang.org/doc/install) and [configured](https://golang.org/doc/install#testing)
```
go get github.com/ranaaditya/tbcd
```
### Usage
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

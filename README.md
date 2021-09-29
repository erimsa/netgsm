[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-soft/netgsm/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-soft/netgsm)](https://pkg.go.dev/github.com/ozgur-soft/netgsm/src)

# Netgsm
An easy-to-use netgsm.com.tr API with golang

# Installation
```bash
go get github.com/ozgur-soft/netgsm // For original repo
go get github.com/erimsa/netgsm // For this repo
```

# Usage
```go
/*
	TO Send SMS
*/
package main

import (
	"fmt"

	netgsm "github.com/ozgur-soft/netgsm/src"
)

func main() {
	config := netgsm.Config{SmsCompany: "NETGSM", SmsMsgHeader: "", SmsUserCode: "", SmsPassword: "", ApiUrl: "https://api.netgsm.com.tr/sms/send/xml"}
	api := &netgsm.API{config}
	request := netgsm.Request{}
	request.MainBody.Body.Msg = "test"
	request.MainBody.Body.No = "905555555555"
	send := api.Sms(request)
	if send {
		fmt.Println("mesaj iletildi")
	} else {
		fmt.Println("hata oluştu")
	}
}
```

```go
/*
	TO Send SMS-OTP
*/
package main

import (
	"fmt"

	netgsm "github.com/ozgur-soft/netgsm/src"
)

func main() {
	config := netgsm.Config{SmsCompany: "NETGSM", SmsMsgHeader: "", SmsUserCode: "", SmsPassword: "", ApiUrl: "https://api.netgsm.com.tr/sms/send/otp"}
	api := &netgsm.API{config}
	request := netgsm.Request{}
	request.MainBody.Body.Msg = "test"
	request.MainBody.Body.No = "905555555555"
	send, otp := api.Sms_Otp(request)
	if send {
		fmt.Println("mesaj gönderildi.")
		fmt.Println("returned otp : " + otp)
	} else {
		fmt.Println("hata oluştu")
	}
}
```

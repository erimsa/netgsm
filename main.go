package main

import (
	"fmt"

	netgsm "github.com/erimsa/netgsm/src"
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

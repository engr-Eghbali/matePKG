package services

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	structs "./basement"

	"gopkg.in/mgo.v2"
)

////////////send mail
func SendMail(body string, recipient string, origin MailOrigin) (er bool) {

	msg := "From: " + origin.From + "\n" +
		"To: " + recipient + "\n" +
		"Subject: Your verification code is : \n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", origin.From, origin.Password, "smtp.gmail.com"),
		origin.From, []string{recipient}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	log.Print("verification code sent to : " + recipient)
	return true
}

////////////////////////////////////////////////////////

////////////sned sms
func SendSms(txt string, recipient string, origin SmsOrigin) bool {

	resp, err := http.Get("https://login.parsgreen.com/UrlService/sendSMS.ashx?from=" + origin.From + "&to=" + recipient + "&&text=" + txt + "&signature=" + origin.ApiKey)
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	content := string(body)
	log.Println(content)
	log.Println(err2)
	if err != nil {
		// handle error
		log.Println("Sending SMS to: " + recipient + " Failed 006 <=End")
		return false
	} else {
		log.Println("SMS Sent to: " + recipient)
		return true
	}

}

///////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

func CreateVcRecord(UID string, session *mgo.Session) string,bool {

	//generate
	rand.Seed(time.Now().UnixNano())
	vc := strconv.Itoa(100000 + rand.Intn(999999-100000))

	VcRecord := structs.VcTable{ID:bson.NewObjectId(),UserID:UID,VC:vc}
	collection := session.DB("bkbfbtpiza46rc3").C("loginRequests")
	InsertErr:=collection.Insert(&VcRecord)

	if InsertErr!=nil{
		log.Println("Creating vc record failed:")
		log.Println(InsertErr)
		log.Println("<=End")
		return "",false
	}else{
		return vc,true
	}


}

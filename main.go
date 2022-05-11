package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Number struct {
	Text            string `xml:",chardata"`
	SendDigits      string `xml:"sendDigits,attr,omitempty"`
	SendOnPreanswer string `xml:"sendOnPreanswer,attr,omitempty"`
}

type User struct {
	Text            string `xml:",chardata"`
	SendDigits      string `xml:"sendDigits,attr,omitempty"`
	SendOnPreanswer string `xml:"sendOnPreanswer,attr,omitempty"`
	SipHeaders      string `xml:"sipHeaders,attr,omitempty"`
}

type Sip struct {
	Text            string `xml:",chardata"`
	SendDigits      string `xml:"sendDigits,attr,omitempty"`
	SendOnPreanswer string `xml:"sendOnPreanswer,attr,omitempty"`
	SipHeaders      string `xml:"sipHeaders,attr,omitempty"`
}

type Play struct {
	XMLName        xml.Name `xml:"Play"`
	Text           string   `xml:",chardata"`
	Loop           string   `xml:"loop,attr,omitempty"`
	CallbackURL    string   `xml:"callback_url,attr,omitempty"`
	CallbackMethod string   `xml:"callback_method,attr,omitempty"`
}

type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Text     string   `xml:",chardata"`
	Loop     string   `xml:"loop,attr,omitempty"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
}

type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Text    string   `xml:",chardata"`
	Method  string   `xml:"method,attr,omitempty"`
}

type Dial struct {
	Text           string  `xml:",chardata"`
	Number         *Number `xml:"Number,omitempty"`
	User           *User   `xml:"User,omitempty"`
	Sip            *Sip    `xml:"Sip,omitempty"`
	Record         string  `xml:"record,attr,omitempty"`
	AnswerOnBridge bool    `xml:"answerOnBridge,attr,omitempty"`
	CallerId       string  `xml:"callerId,attr,omitempty"`
}

type Pause struct {
	Text   string `xml:",chardata"`
	Length int    `xml:"length,attr,omitempty"`
}

type Reject struct {
	Text   string `xml:",chardata"`
	Reason string `xml:"reason,attr,omitempty"`
}

type Gather struct {
	XMLName xml.Name `xml:"Gather"`
	Text    string   `xml:",chardata"`
	Say     *Say     `xml:"Say"`

	Action              string `xml:"action,attr,omitempty"`
	Method              string `xml:"method,attr,omitempty"`
	FinishOnKey         string `xml:"finishOnKey,attr,omitempty"`
	NumDigits           string `xml:"numDigits,attr,omitempty"`
	Timeout             string `xml:"timeout,attr,omitempty"`
	ActionOnEmptyResult string `xml:"actionOnEmptyResult,attr,omitempty"`
	Input               string `xml:"input,attr,omitempty"`
}

type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
	Text    string   `xml:",chardata"`
}
type Response struct {
	XMLName  xml.Name  `xml:"Response"`
	Text     string    `xml:",chardata"`
	Redirect *Redirect `xml:"Redirect,omitempty"`
	Reject   *Reject   `xml:"Reject"`
	Gather   *Gather   `xml:"Gather,omitempty"`
	Dial     *Dial     `xml:"Dial,omitempty"`
	Play     *Play     `xml:"Play,omitempty"`
	Pause    *Pause    `xml:"Pause,omitempty"`
	Say      *Say      `xml:"Say,omitempty"`
	Hangup   *Hangup   `xml:Hangup,omitempty`
}

type StatusCallback struct {
	CallSid       string `json:"CallSid" form:"CallSid" query:"CallSid"`
	AccountSid    string `json:"AccountSid" form:"AccountSid" query:"AccountSid"`
	From          string `json:"From" form:"From" query:"From"`
	To            string `json:"To" form:"To" query:"To"`
	CallStatus    string `json:"CallStatus" form:"CallStatus" query:"CallStatus"`
	ApiVersion    string `json:"ApiVersion" form:"ApiVersion" query:"ApiVersion"`
	Direction     string `json:"Direction" form:"Direction" query:"Direction"`
	ForwardedFrom string `json:"ForwardedFrom" form:"ForwardedFrom" query:"ForwardedFrom"`
	CallerName    string `json:"CallerName" form:"CallerName" query:"CallerName"`
	ParentCallSid string `json:"ParentCallSid" form:"ParentCallSid" query:"ParentCallSid"`

	CallDuration      string `json:"CallDuration,omitempty" form:"CallDuration" query:"CallDuration"`
	SipResponseCode   string `json:"SipResponseCode,omitempty" form:"SipResponseCode" query:"SipResponseCode"`
	RecordingUrl      string `json:"RecordingUrl,omitempty" form:"RecordingUrl" query:"RecordingUrl"`
	RecordingSid      string `json:"RecordingSid,omitempty" form:"RecordingSid" query:"RecordingSid"`
	RecordingDuration string `json:"RecordingDuration,omitempty" form:"RecordingDuration" query:"RecordingDuration"`
	Timestamp         string `json:"Timestamp,omitempty" form:"Timestamp" query:"Timestamp"`
	CallbackSource    string `json:"CallbackSource,omitempty" form:"CallbackSource" query:"CallbackSource"`
	SequenceNumber    string `json:"SequenceNumber,omitempty" form:"SequenceNumber" query:"SequenceNumber"`
	Digits            string `json:"Digits,omitempty" form:"Digits" query:"Digits"`
	UserIntent        string `json:"UserIntent,omitempty" form:"UserIntent" query:"UserIntent"` // base64 encoded, json data.
}

// https://66da3a82c5f1.ngrok.io/TiniyoApplications/MainRestaurantMenu

// var input = []byte(`{"text":"i want to book a table for 4","intent":{"name":"booking_with_count","confidence":0.9603162407875061},"entities":[{"start":27,"end":28,"text":"4","value":4,"confidence":1,"entity":"number"}],"intent_ranking":[{"name":"booking_with_count","confidence":0.9603162407875061},{"name":"booking","confidence":0.0258566252887249},{"name":"booking_with_count_time_day_hours_minute","confidence":0.0034164865501224995},{"name":"booking_with_count_time","confidence":0.0032817136961966753},{"name":"cancel_booking","confidence":0.00111346784979105},{"name":"goodbye","confidence":0.0010961686493828893},{"name":"order_pizza","confidence":0.0008867786964401603},{"name":"bot_challenge","confidence":0.0008118133991956711},{"name":"complain","confidence":0.000798406545072794},{"name":"booking_time_day","confidence":0.0006561993504874408}]}`)

//var base64Input = `eyJ0ZXh0IjoiaSB3YW50IHRvIGJvb2sgYSB0YWJsZSBmb3IgNCIsImludGVudCI6eyJuYW1lIjoiYm9va2luZ193aXRoX2NvdW50IiwiY29uZmlkZW5jZSI6MC45NjAzMTYyNDA3ODc1MDYxfSwiZW50aXRpZXMiOlt7InN0YXJ0IjoyNywiZW5kIjoyOCwidGV4dCI6IjQiLCJ2YWx1ZSI6NCwiY29uZmlkZW5jZSI6MSwiZW50aXR5IjoibnVtYmVyIn1dLCJpbnRlbnRfcmFua2luZyI6W3sibmFtZSI6ImJvb2tpbmdfd2l0aF9jb3VudCIsImNvbmZpZGVuY2UiOjAuOTYwMzE2MjQwNzg3NTA2MX0seyJuYW1lIjoiYm9va2luZyIsImNvbmZpZGVuY2UiOjAuMDI1ODU2NjI1Mjg4NzI0OX0seyJuYW1lIjoiYm9va2luZ193aXRoX2NvdW50X3RpbWVfZGF5X2hvdXJzX21pbnV0ZSIsImNvbmZpZGVuY2UiOjAuMDAzNDE2NDg2NTUwMTIyNDk5NX0seyJuYW1lIjoiYm9va2luZ193aXRoX2NvdW50X3RpbWUiLCJjb25maWRlbmNlIjowLjAwMzI4MTcxMzY5NjE5NjY3NTN9LHsibmFtZSI6ImNhbmNlbF9ib29raW5nIiwiY29uZmlkZW5jZSI6MC4wMDExMTM0Njc4NDk3OTEwNX0seyJuYW1lIjoiZ29vZGJ5ZSIsImNvbmZpZGVuY2UiOjAuMDAxMDk2MTY4NjQ5MzgyODg5M30seyJuYW1lIjoib3JkZXJfcGl6emEiLCJjb25maWRlbmNlIjowLjAwMDg4Njc3ODY5NjQ0MDE2MDN9LHsibmFtZSI6ImJvdF9jaGFsbGVuZ2UiLCJjb25maWRlbmNlIjowLjAwMDgxMTgxMzM5OTE5NTY3MTF9LHsibmFtZSI6ImNvbXBsYWluIiwiY29uZmlkZW5jZSI6MC4wMDA3OTg0MDY1NDUwNzI3OTR9LHsibmFtZSI6ImJvb2tpbmdfdGltZV9kYXkiLCJjb25maWRlbmNlIjowLjAwMDY1NjE5OTM1MDQ4NzQ0MDh9XX0=`

func main() {

	e := echo.New()
	numberRestMap := new(PhonenumberMap)

	e.GET("/v1/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!!!")
	})

	e.POST("/TiniyoApplications/DirectCall", func(c echo.Context) error {
		resp := &Response{}
		resp.Text = ""
		u := StatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := &Response{}
			rejectresp.Reject = &Reject{
				Reason: "rejected",
			}
			return c.XML(http.StatusOK, resp)
		}
		resp.Dial = &Dial{
			AnswerOnBridge: true,
			Number: &Number{
				Text: u.To,
			},
		}
		return c.XML(http.StatusOK, resp)
	})

	// from number : customer number
	// to number : inbound tiniyo number.
	e.GET("/TiniyoApplications/MainRestaurantMenu", func(c echo.Context) error {
		u := StatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := &Response{}
			rejectresp.Reject = &Reject{
				Reason: "rejected",
			}
			return c.XML(http.StatusOK, rejectresp)
		}
		ivrRest := new(RestaurentIVR)
		numberRestMap.StoreNumberInstance(u.From, ivrRest)
		resp := ivrRest.GetMainMenuResponse()
		return c.XML(http.StatusOK, resp)
	})

	e.POST("/TiniyoApplications/DtmfReceived", func(c echo.Context) error {
		u := StatusCallback{}
		err := c.Bind(&u)
		rejectresp := GetRejectedResponse()

		if err != nil {
			return c.XML(http.StatusOK, rejectresp)
		}

		fmt.Println("Dtmf Digit : ", u.Digits, u.From)

		if len(u.Digits) == 0 {
			return c.XML(http.StatusOK, rejectresp)
		}

		lastChar := u.Digits[len(u.Digits)-1:]

		if lastChar == "#" {
			u.Digits = u.Digits[0 : len(u.Digits)-1]
		}

		ivrRest := numberRestMap.GetNumberInstance(u.From)

		if ivrRest == nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}

		resp := ivrRest.ProcessDTMFDigits(u.Digits)

		return c.XML(http.StatusOK, resp)
	})

	e.POST("/TiniyoApplications/Voicebot", func(c echo.Context) error {
		u := StatusCallback{}
		err := c.Bind(&u)

		rejectresp := GetRejectedResponse()

		if err != nil {
			return c.XML(http.StatusOK, rejectresp)
		}

		ivrRest := numberRestMap.GetNumberInstance(u.From)

		return c.XML(http.StatusOK, ivrRest.CreateWelcomeVoiceBot("Welcome to Big Pitcher, How can i help you?"))
	})

	e.POST("/TiniyoApplications/UserIntent", func(c echo.Context) error {
		resp := &Response{}
		resp.Text = ""
		u := StatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := &Response{}
			rejectresp.Reject = &Reject{
				Reason: "rejected",
			}
			return c.XML(http.StatusOK, resp)
		}
		ProcessUserIntent(u.UserIntent)
		return c.XML(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

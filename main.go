package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	TextType string   `xml:"textType,attr,omitempty"`
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
	VoiceMaxDuration    string `xml:"voiceMaxDuration,attr,omitempty"`
	VoicePreSilence     string `xml:"voicePreSilence,attr,omitempty"`
	VoicePostSilence    string `xml:"voicePostSilence,attr,omitempty"`
	VoiceMode           string `xml:"voiceMode,attr,omitempty"`
	VoiceAckSay         string `xml:"voiceAckSay,attr,omitempty"`
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

var base64Input = `eyJ0ZXh0IjoiaSB3YW50IHRvIGJvb2sgYSB0YWJsZSBmb3IgNCIsImludGVudCI6eyJuYW1lIjoiYm9va2luZ193aXRoX2NvdW50IiwiY29uZmlkZW5jZSI6MC45NjAzMTYyNDA3ODc1MDYxfSwiZW50aXRpZXMiOlt7InN0YXJ0IjoyNywiZW5kIjoyOCwidGV4dCI6IjQiLCJ2YWx1ZSI6NCwiY29uZmlkZW5jZSI6MSwiZW50aXR5IjoibnVtYmVyIn1dLCJpbnRlbnRfcmFua2luZyI6W3sibmFtZSI6ImJvb2tpbmdfd2l0aF9jb3VudCIsImNvbmZpZGVuY2UiOjAuOTYwMzE2MjQwNzg3NTA2MX0seyJuYW1lIjoiYm9va2luZyIsImNvbmZpZGVuY2UiOjAuMDI1ODU2NjI1Mjg4NzI0OX0seyJuYW1lIjoiYm9va2luZ193aXRoX2NvdW50X3RpbWVfZGF5X2hvdXJzX21pbnV0ZSIsImNvbmZpZGVuY2UiOjAuMDAzNDE2NDg2NTUwMTIyNDk5NX0seyJuYW1lIjoiYm9va2luZ193aXRoX2NvdW50X3RpbWUiLCJjb25maWRlbmNlIjowLjAwMzI4MTcxMzY5NjE5NjY3NTN9LHsibmFtZSI6ImNhbmNlbF9ib29raW5nIiwiY29uZmlkZW5jZSI6MC4wMDExMTM0Njc4NDk3OTEwNX0seyJuYW1lIjoiZ29vZGJ5ZSIsImNvbmZpZGVuY2UiOjAuMDAxMDk2MTY4NjQ5MzgyODg5M30seyJuYW1lIjoib3JkZXJfcGl6emEiLCJjb25maWRlbmNlIjowLjAwMDg4Njc3ODY5NjQ0MDE2MDN9LHsibmFtZSI6ImJvdF9jaGFsbGVuZ2UiLCJjb25maWRlbmNlIjowLjAwMDgxMTgxMzM5OTE5NTY3MTF9LHsibmFtZSI6ImNvbXBsYWluIiwiY29uZmlkZW5jZSI6MC4wMDA3OTg0MDY1NDUwNzI3OTR9LHsibmFtZSI6ImJvb2tpbmdfdGltZV9kYXkiLCJjb25maWRlbmNlIjowLjAwMDY1NjE5OTM1MDQ4NzQ0MDh9XX0=`

const (
	Greet   = "greet"
	GoodBye = "goodbye"
	Affirm  = "affirm"
	Deny    = "deny"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	numberRestMap := new(PhonenumberMap)

	// ProcessUserIntent(base64Input)

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

	e.POST("/TiniyoApplications/VoicebotLoan", func(c echo.Context) error {
		u := StatusCallback{}
		err := c.Bind(&u)

		rejectresp := GetRejectedResponse()

		if err != nil {
			return c.XML(http.StatusOK, rejectresp)
		}

		ivrRest := numberRestMap.GetNumberInstance(u.From)

		if ivrRest == nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}

		return c.XML(http.StatusOK, ivrRest.CreateWelcomeVoiceBot("Welcome to ICICI Bank, Are you interested in home loan? Say YES or NO"))
	})

	e.POST("/TiniyoApplications/Voicebot", func(c echo.Context) error {
		u := StatusCallback{}
		err := c.Bind(&u)

		rejectresp := GetRejectedResponse()

		if err != nil {
			return c.XML(http.StatusOK, rejectresp)
		}

		fmt.Println("Voicebot FROM : ", u.From)

		ivrRest := new(RestaurentIVR)
		numberRestMap.StoreNumberInstance(u.From, ivrRest)

		// ssmlText := `<speak xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="http://www.w3.org/2001/mstts" xmlns:emo="http://www.w3.org/2009/10/emotionml" version="1.0" xml:lang="en-US"><voice name="hi-IN-MadhurNeural"><prosody rate="0%" pitch="0%">BigPitcher में आपका स्वागत है</prosody></voice></speak>`
		ssmlText := `<speak xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="http://www.w3.org/2001/mstts" xmlns:emo="http://www.w3.org/2009/10/emotionml" version="1.0" xml:lang="en-US"><voice name="en-IN-NeerjaNeural"><prosody rate="0%" pitch="0%">Welcome to Big Pitcher, How can i help you?</prosody></voice></speak>`
		// ssmlText := `<speak xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="http://www.w3.org/2001/mstts" xmlns:emo="http://www.w3.org/2009/10/emotionml" version="1.0" xml:lang="en-US"><voice name="hi-IN-MadhurNeural"><prosody rate="0%" pitch="0%">आप इस टेक्स्ट को अपनी इच्छानुसार किसी भी टेक्स्ट से बदल सकते हैं।  आप या तो इस टेक्स्ट बॉक्स में लिख सकते हैं या अपना टेक्स्ट यहां पेस्ट कर सकते हैं।
		// विभिन्न भाषाओं और आवाजों को आजमाएं। आवाज की गति और पिच को बदलें।  आप SSML (स्पीच सिंथेसिस मार्कअप लैंग्वेज) को भी बदल सकते हैं ताकि यह नियंत्रित किया जा सके कि टेक्स्ट के विभिन्न खंड कैसे ध्वनि करते हैं। इसे आजमाने के लिए ऊपर दिए गए SSML पर क्लिक करें!
		//  टेक्स्ट टू स्पीच का उपयोग करने का आनंद लें!</prosody></voice></speak>`
		return c.XML(http.StatusOK, ivrRest.CreateWelcomeVoiceBot(ssmlText))
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

		userIntent := ProcessUserIntent(u.UserIntent)

		fmt.Println("Voicebot FROM : ", u.From)

		ivrRest := numberRestMap.GetNumberInstance(u.From)

		if ivrRest == nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}

		fmt.Println(userIntent)

		prefix := `<speak xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="http://www.w3.org/2001/mstts" xmlns:emo="http://www.w3.org/2009/10/emotionml" version="1.0" xml:lang="en-US"><voice name="en-IN-NeerjaNeural"><prosody rate="0%" pitch="0%">`
		postfix := `</prosody></voice></speak>`
		switch userIntent.Intent.Name {
		case "greet":
			fmt.Println("greet")
			ssmlText := prefix + `Hello` + postfix
			resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
		case "goodbye":
			fmt.Println("goodbye")
			ssmlText := prefix + `Good Bye, it was nice talking to you.` + postfix
			resp = ivrRest.CreateSayHangupSSML(ssmlText)
		case "affirm":
			fmt.Println("affirm")
			ssmlText := prefix + `Thank you! for confirming that you are interested in booking table. how can i help you with booking?` + postfix
			resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
		case "deny":
			fmt.Println("deny")
			ssmlText := prefix + `Ok No problem, we will not call you back again. Thanks for your feedback on product. Bye` + postfix
			resp = ivrRest.CreateSayHangupSSML(ssmlText)
		case "complain":
			fmt.Println("complain")
			ssmlText := prefix + `Let me help you in raising complain. Please speak out for next 5 minute, your audio is recorded and on priority your complaing would be analysed and resolved.` + postfix
			resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
		case "booking":
			fmt.Println("booking")
			ssmlText := prefix + `I can help you with booking of table, For how many persons do you need reservation?` + postfix
			resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
		case "booking_with_count":
			fmt.Println("booking_with_count")
			if userIntent.Entities[0].Entity == "number" {
				fmt.Println("number of persongs : ", userIntent.Entities[0].Value)
				// ask for what time today or tomorrow?
				ivrRest.SetCount(userIntent.Entities[0].Value)
				ssmlText := prefix + `I would like to confirm that you need booking for ` + userIntent.Entities[0].Text + `, and do you need booking today, tomorrow or day after tomorrow?` + postfix
				resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
			}
		case "booking_with_count_time":
			fmt.Println("booking_with_count_time")
			// ask for the
		case "booking_time_day":
			fmt.Println("booking_time_day")
			if userIntent.Entities[0].Entity == "time" {
				fmt.Println("time of booking : ", userIntent.Entities[0].Value)
				ivrRest.SetDayTime(userIntent.Entities[0].Text)
				ssmlText := prefix + `Ok i will do booking ` + userIntent.Entities[0].Text + ` for you, What time do you need booking? you can say like 9:30PM or 10:30AM.` + postfix
				resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
			}
			// ivrRest.CreateWelcomeVoiceBot(ssmlText)
		case "booking_with_time_day_hours_minute":
			fmt.Println("booking_with_time_day_hours_minute")
			if userIntent.Entities[0].Entity == "time" {
				fmt.Println("time of booking : ", userIntent.Entities[0].Text, userIntent.Entities[0].Value)
				// ask for the time of today or tomorrow?
				ssmlText := prefix + `Ok i will do booking ` + userIntent.Entities[0].Text + ` for you.` + postfix
				resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
			}
		case "booking_with_count_time_day_hours_minute":
			fmt.Println("booking_with_count_time_day_hours_minute")
			for i := 0; i < len(userIntent.Entities); i++ {
				if userIntent.Entities[0].Entity == "time" {
					fmt.Println("time of booking : ", userIntent.Entities[0].Text, userIntent.Entities[0].Value)
					ivrRest.SetDayTime(userIntent.Entities[0].Text)
				} else if userIntent.Entities[0].Entity == "number" {
					fmt.Println("number of persongs : ", userIntent.Entities[0].Value)
				}
			}
			ssmlText := prefix + `Ok i will do booking ` + userIntent.Entities[0].Text + ` for you.` + postfix
			resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)

		case "booking_count":
			fmt.Println("booking_count")
			if userIntent.Entities[0].Entity == "number" {
				fmt.Println("number of persongs : ", userIntent.Entities[0].Value)
				ivrRest.SetCount(userIntent.Entities[0].Value)
				// ask for what time today or tomorrow?
				ssmlText := prefix + `I would like to confirm that you need booking for ` + userIntent.Entities[0].Text + `, and do you need booking today, tomorrow or day after tomorrow?` + postfix
				resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
			}
		case "talkto_agent":
			fmt.Println("talkto_agent")
			ssmlText := prefix + `Ok , Let me transfer your call to an agent. Transferring call now.` + postfix
			resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
		case "nlu_fallback":
			ssmlText := prefix + `I don't understand can you please say that again?` + postfix
			resp = ivrRest.CreateWelcomeVoiceBot(ssmlText)
		default:
			fmt.Println("Invalid")
		}

		return c.XML(http.StatusOK, resp)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

type RestaurentIVR struct {
	state                 string
	phoneNumber           string
	inboundDidNumber      string
	numberofPeopleJoining string
	timing                string
	level                 string
	subtime               string
}

func GetRejectedResponse() *Response {
	rejectresp := &Response{}
	rejectresp.Reject = &Reject{
		Reason: "rejected",
	}
	return rejectresp
}

func (rivr RestaurentIVR) createGatherSayResponse(gatherSayString string, digits string) *Response {
	resp := &Response{}
	resp.Text = ""
	resp.Gather = &Gather{
		Action:      "https://tiniyo.dev/TiniyoApplications/DtmfReceived",
		NumDigits:   digits,
		FinishOnKey: "#",
		Method:      "POST",
	}
	resp.Gather.Say = &Say{
		Text: gatherSayString,
	}
	resp.Say = &Say{
		Text: "We didn't receive any input. Goodbye!",
	}
	return resp
}

func (rivr RestaurentIVR) createSayHangup(sayString string) *Response {
	resp := &Response{}
	resp.Text = ""
	resp.Say = &Say{
		Text: sayString,
	}
	resp.Hangup = &Hangup{}
	return resp
}

func (rivr *RestaurentIVR) GetMainMenuResponse() *Response {
	Text := "Thank you for calling Barbeque nation, Where customer service is our priority. Press 1 to book table. Press 2 to retrive existing booking. " +
		" Press 3 to cancel table booking. Press 4 to get loyalty points. Press 0 to repeat the menu"
	rivr.state = "IDLE"
	return rivr.createGatherSayResponse(Text, "1")
}

func (rivr *RestaurentIVR) GetRepeatMenuResponse() *Response {
	Text := "Press 1 to book table. Press 2 to retrive existing booking. " +
		" Press 3 to cancel table booking. Press 4 to get loyalty points. Press 5 To speak with our sales representative for table booking. " +
		" Press 6 to give feedback. Press 7 for availability of table. Press 0 to repeat the menu"
	rivr.state = "IDLE"
	return rivr.createGatherSayResponse(Text, "1")
}

func (rivr *RestaurentIVR) GetPeopleNunberCollectionResponse() *Response {
	Text := "how many people are joining with you for dineout? press number from your touch pad. like 1, 4 and # for 14 people."
	rivr.state = "USER_COLLECT"
	return rivr.createGatherSayResponse(Text, "3")
}

func (rivr *RestaurentIVR) GetExistingBooking() *Response {
	Text := "We dont have booking for your table. Press 0 to go back to main menu."
	rivr.state = "IDLE"
	return rivr.createGatherSayResponse(Text, "1")
}

func (rivr *RestaurentIVR) CancelBooking() *Response {
	Text := "Your current booking cancelled. Press 0 to go back to main menu."
	rivr.state = "IDLE"
	return rivr.createGatherSayResponse(Text, "1")
}

func (rivr *RestaurentIVR) GetLoyaltyPoints() *Response {
	Text := "Your loyalty points are 1000. Press 0 to go back to main menu."
	rivr.state = "IDLE"
	return rivr.createGatherSayResponse(Text, "1")
}

func (rivr *RestaurentIVR) ProcessDTMFDigits(digits string) *Response {
	if rivr.state == "IDLE" {
		if digits == "0" {
			return rivr.GetRepeatMenuResponse()
		} else if digits == "1" {
			return rivr.GetPeopleNunberCollectionResponse()
		} else if digits == "2" {
			return rivr.GetExistingBooking()
		} else if digits == "3" {
			return rivr.CancelBooking()
		} else if digits == "4" {
			return rivr.GetLoyaltyPoints()
		}
	} else if rivr.state == "USER_COLLECT" {
		// digits is number of users
		rivr.numberofPeopleJoining = digits
		rivr.state = "USER_COLLECTED"
		return rivr.createGatherSayResponse("Press 1 to book table today, Press 2 to book table tomorrow and Press 3 for custom date.", "1")
	} else if rivr.state == "USER_COLLECTED" {
		if digits == "1" {
			rivr.timing = "today"
		} else if digits == "2" {
			rivr.timing = "tomorrow"
		} else if digits == "3" {
			rivr.state = "COLLECT_CUSTOM_TIME"
			return rivr.createGatherSayResponse("You can enter date in D D  M M  Y Y Y Y format. for example 10th May 2021 would be 10 05 2021.", "8")
		}
		rivr.state = "COLLECT_SUBTIME"
		return rivr.createGatherSayResponse("Press 1 for breakfast, press 2 for lunch, press 3 for dinner.", "1")
	} else if rivr.state == "COLLECT_CUSTOM_TIME" {
		// dtmf digits would be in ddmmyyyy format.

		date := digits[0:2] + "-" + digits[2:4] + "-" + digits[4:8]

		layout := "02-01-2006"
		t, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println(err)
		}
		y, m, d := t.Date()

		rivr.timing = strconv.Itoa(d) + " " + m.String() + " " + strconv.Itoa(y)

		rivr.state = "COLLECT_SUBTIME"
		return rivr.createGatherSayResponse("Press 1 for breakfast, press 2 for lunch, press 3 for dinner.", "1")
	} else if rivr.state == "COLLECT_SUBTIME" {
		if digits == "1" {
			rivr.subtime = "breakfast"
		} else if digits == "2" {
			rivr.subtime = "lunch"
		} else if digits == "3" {
			rivr.subtime = "dinner"
		}
		return rivr.createSayHangup("Your table booked for " + rivr.numberofPeopleJoining + " people, " + rivr.timing + " " + rivr.subtime + ", Thankyou for booking table with brabaque nation.")
	}
	return nil
}

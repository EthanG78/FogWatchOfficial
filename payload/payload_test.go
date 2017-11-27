package payload_test

import (
	"github.com/EthanG78/fog_watch/payload"
	"github.com/stretchr/testify/assert"
	"testing"
)

//This is to test if allocating new data through Payload methods works.
func TestPayloadMethods(t *testing.T) {
	data := payload.Payload{}

	//This is the test payload that will be compared with actual^^
	testData := payload.Payload{
		Date:     "11/12/2017",
		Location: "Uptown",
		Temp:     "15",
		Humidity: "3%",
		WindS:    "8km/h",
		Status:   "Active",
	}

	//Methods for inputing new data
	data.SetDate("11/12/2017")
	data.SetLocale("Uptown")
	data.SetTemp("15")
	data.SetHumidity("3%")
	data.SetWindSpeed("8km/h")
	data.SetStatus("Active")

	assert.Equal(t, testData, data, "Test if payload methods are functioning")
}

func TestPayloadRetrieval(t *testing.T) {
	firebase := "https://fogwatch-45fe5.firebaseio.com/"
	field := "TestData"
	data := payload.GetPayload(firebase, field)

	//Tests to see if map returned is nil or not
	if assert.NotNil(t, data, "Checks if GetPayload succesfully fetched firebase data") {
		t.Log(data)
		//Debugging
	} else if assert.Nil(t, data, "If the payload is nil, return err") {
		t.Error("Failed to fetch data from firebase")
	}
}

func TestPayload(t *testing.T) {
	firebase := "https://fogwatch-45fe5.firebaseio.com/"
	key := "TestData"

	date := payload.GetPayloadField(firebase, key, "Date")
	assert.Equal(t, "11-14-2017", date)

	humidity := payload.GetPayloadField(firebase, key, "Humidity")
	assert.Equal(t, "100%", humidity)

	location := payload.GetPayloadField(firebase, key, "Location")
	assert.Equal(t, "Uptown", location)

	status := payload.GetPayloadField(firebase, key, "Status")
	assert.Equal(t, "Active", status)

	temp := payload.GetPayloadField(firebase, key, "Temp")
	assert.Equal(t, "20", temp)

	windspeed := payload.GetPayloadField(firebase, key, "WindS")
	assert.Equal(t, "8km/h", windspeed)



}

//Continue Test

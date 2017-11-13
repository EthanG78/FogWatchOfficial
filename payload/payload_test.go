package payload_test

import (
	"github.com/EthanG78/fog_watch/payload"
	"github.com/stretchr/testify/assert"
	"testing"
)

//This is to test if allocating new data through Payload methods works.
func TestPayload(t *testing.T) {
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

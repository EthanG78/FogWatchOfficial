package payload

import (
	"encoding/json"
	"fmt"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
	//"time"
)

type Payload struct {
	Date     string
	Location string
	Temp     string
	Humidity string
	WindS    string
	Status   string
}

//For printing all data
func (p *Payload) PrintPayload() {
	fmt.Printf("Date: %v\nLocation: %v\nTemperature: %v\nHumidity: %v\nWind Speed: %v\nStatus: %v/n", p.Date, p.Location, p.Temp, p.Humidity, p.WindS, p.Status)
}

func GetPayload(firebase, key string) Payload {
	payload := Payload{
		Date:     GetPayloadField(firebase, key, "Date"),
		Location: GetPayloadField(firebase, key, "Location"),
		Temp:     GetPayloadField(firebase, key, "Temp"),
		Humidity: GetPayloadField(firebase, key, "Humidity"),
		WindS:    GetPayloadField(firebase, key, "WindS"),
		Status:   GetPayloadField(firebase, key, "Status"),
	}
	return payload
}

//For fetching individual information from firebase
func GetPayloadField(firebase, key, field string) string {
	//localDate := time.Now().Local()

	var unmData map[string]interface{}
	var mData map[string]interface{}
	var valField string

	f := firego.New(firebase, nil)

	if err := f.Value(&unmData); err != nil {
		log.Fatalf("Error retrieving firebase data: %v", err)
	}

	marshaled, err := json.Marshal(unmData)
	if err != nil {
		log.Fatalf("Failed to marshal: %v, %v", unmData, err)
	}

	if err := json.Unmarshal(marshaled, &mData); err != nil {
		log.Fatalf("Failed to Unmarshal: %v, %v", marshaled, err)
	}

	for _, valKey := range mData {
		keyMap, ok := valKey.(map[string]interface{})
		if !ok {
			continue
		}
		//field := localDate.Format("01-02-2006")
		//This field variable is for fetching data with this specific prefix (temp, date, etc..)
		//Will make this a parameter for people to call
		if keyData, ok := keyMap[key]; ok {
			fieldMap, ok := keyData.(map[string]interface{})
			if !ok {
				continue
			}
			for indexField, valField := range fieldMap {

				if indexField == field {
					//Make sure valField is a string
					valField, ok := valField.(string)
					if !ok {
						continue
					}
					return valField
				}
			}
		}

	}

	return valField

}

//These are for setting payload data
func (p *Payload) SetDate(date string) {
	p.Date = date
}

func (p *Payload) SetLocale(location string) {
	p.Location = location
}

func (p *Payload) SetTemp(temp string) {
	p.Temp = temp
}

func (p *Payload) SetHumidity(humidity string) {
	p.Humidity = humidity
}

func (p *Payload) SetWindSpeed(wind string) {
	p.WindS = wind
}

func (p *Payload) SetStatus(status string) {
	p.Status = status
}

//Reading payload data (for debugging mostly)
func (p *Payload) GetDate() string {
	return p.Date
}

func (p *Payload) GetLocale() string {
	return p.Location
}

func (p *Payload) GetTemp() string {
	return p.Temp
}

func (p *Payload) GetHumidity() string {
	return p.Humidity
}

func (p *Payload) GetWindSpeed() string {
	return p.WindS
}

func (p *Payload) GetStatus() string {
	return p.Status
}

//TODO: These metods are REALLY SLOW!!! fix for the future

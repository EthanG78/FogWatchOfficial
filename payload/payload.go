package payload

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/zabawaba99/firego.v1"
	"io/ioutil"
	"log"
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
	date, err := GetPayloadField(firebase, key, "Date")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	location, err := GetPayloadField(firebase, key, "Location")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	temp, err := GetPayloadField(firebase, key, "Temp")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	humidity, err := GetPayloadField(firebase, key, "Humidity")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	wind, err := GetPayloadField(firebase, key, "WindS")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	status, err := GetPayloadField(firebase, key, "Status")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	payload := Payload{
		Date:     date,
		Location: location,
		Temp:     temp,
		Humidity: humidity,
		WindS:    wind,
		Status:   status,
	}
	return payload
}

//For fetching individual information from firebase
func GetPayloadField(firebase, key, field string) (string, error) {
	//localDate := time.Now().Local()

	var unmData map[string]interface{}
	var mData map[string]interface{}
	var valField string

	//For webapp authentification
	d, err := ioutil.ReadFile("payload/service_account.json")
	if err != nil {
		return "", err
	}

	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		return "", err
	}

	f := firego.New(firebase, conf.Client(oauth2.NoContext))

	if err := f.Value(&unmData); err != nil {
		log.Fatalf("Error retrieving firebase data: %v", err)
		return "", err
	}

	marshaled, err := json.Marshal(unmData)
	if err != nil {
		log.Fatalf("Failed to marshal: %v, %v", unmData, err)
		return "", err
	}

	if err := json.Unmarshal(marshaled, &mData); err != nil {
		log.Fatalf("Failed to Unmarshal: %v, %v", marshaled, err)
		return "", err
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
					return valField, nil
				}
			}
		}

	}

	return valField, nil

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

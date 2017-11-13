package payload

import "fmt"

type Payload struct {
	Date     string
	Location string
	Temp     string
	Humidity string
	WindS    string
	Status   string
}

//This map will keep the paylod data along with it's unique Key
//data := make(map[string]Payload)

//For printing all data
func (p *Payload) PrintPayload() {
	fmt.Printf("Date: %v\nLocation: %v\nTemperature: %v\nHumidity: %v\nWind Speed: %v\nStatus: %v/n", p.Date, p.Location, p.Temp, p.Humidity, p.WindS, p.Status)
}

func (p *Payload) GetPayload() *Payload {
	return p
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

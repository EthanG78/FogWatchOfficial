#!/usr/bin/python

#This will be the script used for weather data collction within the raspberry pi weather stations
import Adafruit_DHT #Library for reading DHT sensors
import json
import pytz
import datetime
from firebase import firebase

#Setting local timezone
tz = ytz.timezone('America/Moncton')
now = datetime.datetime.now(tz)
firebase = firebase.FirebaseApplication("https://fogwatch-45fe5.firebaseio.com", None)

#Configuring DHT22 humidity sensor
sensor = Adafruit_DHT.DHT22
pin = 22
hum,temp = Adafruit_DHT.read_retry(sensor,pin)


#This will be the payload of data being sent to firebase
payload = {
	"Date/": now.strftime("%m-%d-%Y"),
	"Temperature/": "{0:.1f}*C".format(temp),
	"Humidity/": "{0:.1f}%".format(hum),
	"Status/": "Active",
}

results = firebase.path("/prototype/UptownSJ/" + now.strftime("%H:%M:%S"),data)
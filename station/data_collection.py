#!/usr/bin/python

#This will be the script used for weather data collction within the raspberry pi weather stations
import Adafruit_DHT #Library for reading DHT sensors
import os
import glob 
import json
import pytz
import datetime
from firebase import firebase

#Setting local timezone
tz = pytz.timezone('America/Moncton')
now = datetime.datetime.now(tz)
firebase = firebase.FirebaseApplication("https://fogwatch-45fe5.firebaseio.com", None)

#For reading oneWire sensor
os.system('modprobe w1-gpio')
os.system('modprobe w1-therm')

base_dir = '/sys/bus/w1/devices/'
device_folder = glob.glob(base_dir + '28*')[0]
device_file = device_folder + '/w1_slave'

#Configuring DHT22 humidity sensor
sensor = Adafruit_DHT.DHT22
pin = 22 #Wiring pi pin 3
#The DHT22 will only be used for humidity and maybe internal temo
hum,internalTemp = Adafruit_DHT.read_retry(sensor,pin)


#For reading oneWire data from RPI
def read_temp_raw():
	f = open(device_file, 'r')
	lines = f.readlines()
	f.close()
	return lines

def read_temp():
	lines = read_temp_raw()
	while lines[0].strip()[-3:] != 'YES':
		time.sleep(0.2)
		lines = read_temp_raw()
	equals_pos = lines[1].find('t=')
	if equals_pos != -1:
		temp_string = lines[1][equals_pos+2:]
		temp_c = float(temp_string) / 1000.0
		return str(temp_c)


#This will be the payload of data being sent to firebase
payload = {
	"Date/": now.strftime("%m-%d-%Y"),
	"Temperature/": read_temp(),
	"Humidity/": "{0:.1f}%".format(hum), 
	"Status/": "Active",
}

results = firebase.patch("/prototype/UptownSJ/" + now.strftime("%H:%M:%S"),data)
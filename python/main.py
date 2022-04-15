import requests
import sys

if len(sys.argv) != 2:
    print("First argument should be an ICAO station.")
    exit()

weather_report = requests.get(f'https://api.paradox.ovh/metar?station={sys.argv[1]}').json()['details']['descriptions']['full_description']
print(weather_report)
# Watertemp app

Go program that fetches water temperature information from an API.

## Datasource

The data retrieved is published by Södertälje kommun. It represents the water temperature at locations around Södertälje. The endpoint used in this project exposes the most recently available data, i.e. it answers the question "What temperature is it right now?". Example:  

````json
[
  {
    "type": "Watertemp",
    "temp_water": 7.8,
    "formatted_time": "Nov 02 2021 20:10:14",
    "alias": "Eklundsnäsbadet",
    "ts": 1635880214580,
    "latitude": 59.16883,
    "longitude": 17.59184,
    "gmap": "https://www.google.com/maps/search/?api=1&query=59.16883,17.59184"
  },
  ...
]
``` 

For more information, see 

* https://www.sodertalje.se/kommun-och-politik/for-medborgare/oppna-data/oppen-data/#esc_entry=870&esc_context=1
* https://catalog.sodertalje.se/store/1/resource/862

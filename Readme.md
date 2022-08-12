# Musement CLI App
#### Author: Federico Ayme

This is a Go CLI App for getting the forecast printed in console for all the cities where Musement has activities to sell.

The forecast is printed in format: 

`Processed city [city name] | [weather today] - [wheather tomorrow]`

For this purpose the information about the cities is get from musementAPI `/api/v3/cities` (more information [here](https://api.musement.com/swagger_3.5.0.json))

The information about the forecast for each city is get from weatherAPI `http://api.weatherapi.com/v1/forecast.json?key=[your-key]&q=[lat],[long]&days=2` (more information [here](https://www.weatherapi.com/docs/))

## Install

In root folder `docker build --tag docker-forecast-cities .`

`docker run -it --rm docker-forecast-cities`

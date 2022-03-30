# Making concurrent/parallel requests

This sample application shows how we can make multiple, parallel requests to 
a weather api to bring the data of several locations in parallel

For this example, we will be using the [Open Meteo](https://open-meteo.com/en/docs) api
to fetch the following data points for each locations -

1. Elevation
2. Temperature
3. Wind Speed
4. Wind Direction


Since Open Meteo only works with *Latitude and Longitudes* to retrieve weather
we will be using a pre defined json of cities and their lat longs to get the
actual coordinates of the cities. 

The list is pulled from the **[lutangar/cities.json](https://github.com/lutangar/cities.json)** project
on github which is also available as an npm module.
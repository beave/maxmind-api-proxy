# maxmind-api-proxy
The "maxmind-api-proxy" is a small Golang program that acts as a "proxy" between a "client" and the Maxmind GeoIP
API (https://www.maxmind.com/en/solutions/ip-geolocation-databases-api-services).

As requests are made to the "proxy", results are stored (cached) to a Redis database.  The database will automatically
"expire" entire "cached" entires after a specified amount of time (default: 24 hours).

The idea is to query Maxmind less,  thus saving you money on queries. 


## Groupie Tracker 

This project is a web application that tracks musical artists, their locations, and dates of events. It serves data from an external API and displays artist information using Google Maps API integration for locations.

## Live website:
https://groupie-tracker.fly.dev/

#### How to use

```bash
git clone https://learn.zone01oujda.ma/git/aammar/groupie-tracker-geolocalization
```

```bash
go run .
```

The project contains two main components:

- **The server**: responsible for fetching data from the api and serving the web application.
- **The API**: responsible for fetching and serving artist-related data from the external source.

> Setting ports (to use a specific port):

```bash
export PORT=<port number>
export APIPORT=<api port number>"
```

#### Google Maps Integration
This project integrates the Google Maps API to display artist locations on an interactive map. The locations are fetched from an external API and are dynamically displayed on the map as markers, representing various places where the artists have performed or will perform in the future.
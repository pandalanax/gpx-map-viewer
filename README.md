# gpx-map-viewer
[![Docker image](https://github.com/pandalanax/gpx-map-viewer/actions/workflows/publish_docker.yml/badge.svg)](https://github.com/pandalanax/gpx-map-viewer/actions/workflows/publish_docker.yml)

Really simple website that just puts gpx files on a map and displays them in the browser.
It automatically tries to fit the map so that every gpx route is displayed.

I use it to have a "heatmap" with all my gpx tracks on a map.

## Run
Use docker-compose.yml in the repo:
```yml
version: "3.8"

services:
  gpx-viewer:
    image: pandalanax/gpx-map-viewer:latest
    ports:
      - "8080:8080"
    volumes:
      - ./gpx_tracks:/app/gpx_tracks
    restart: unless-stopped
```


![image](https://github.com/user-attachments/assets/4933987c-a0b5-447b-857c-e4868e8cbfd2)

# Acknowledgements

- [leaflet](https://github.com/Leaflet/Leaflet)
- [leaflet-gpx](https://github.com/mpetazzoni/leaflet-gpx)

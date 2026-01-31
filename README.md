# GPX Manipulator

## Purpose

This app is designed to read GPX data files in and out, while changing some information such as the timezone information.

This is essentially to validate and format GPX data files for use with tools such as `exiftool` for geotagging photos.

## Usage

**Convert a GPX file to use Australia/Sydney time**
```
go run main.go -timezone='Australia/Sydney' -input=/path/to/input.gpx -output=/path/to/output.gpx
```

**Convert a GPX file to use UTC/Zulu time**
```
go run main.go -timezone='UTC' -input=/path/to/input.gpx -output=/path/to/output.gpx
```

## License

MIT

## Other

AI was used for some aspects of this project.
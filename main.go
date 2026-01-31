package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	gpxtypes "github.com/fubarhouse/gpx-manipulator/internal/lib/gpx/types"
)

var (
	timezone         string
	input            string
	output           string
	disableRewriting bool
)

func main() {
	// Validate input
	flag.StringVar(&timezone, "timezone", "Australia/Sydney", "")
	flag.StringVar(&input, "input", "", "")
	flag.StringVar(&output, "output", "", "")
	flag.BoolVar(&disableRewriting, "disable-rewriting", false, "")
	flag.Parse()

	// Open input file
	f, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Decode the input.
	dec := xml.NewDecoder(f)
	var gpx gpxtypes.GPX
	if err := dec.Decode(&gpx); err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "XML decode error: %v\n", err)
		os.Exit(1)
	}

	// Process the data in the input GPX data file.
	ProcessGPX(&gpx)

	// Marshall/read new data into memory.
	marshaled, err := xml.MarshalIndent(gpx, "", "  ")
	if err != nil {
		log.Fatalf("marshal error: %v", err)
	}
	marshaled = append([]byte(xml.Header), marshaled...)

	// Write the file.
	if err := os.WriteFile(output, marshaled, 0o644); err != nil {
		log.Fatalf("write error: %v", err)
	}
	fmt.Printf("Successfully rewritten GPX to %s\n", output)
}

// toZone will convert the input time to the destination timezone.
func toZone(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}

// ProcessGPX will process the input GPX file.
func ProcessGPX(gpx *gpxtypes.GPX) {
	for i, trk := range gpx.Tracks {
		for n, seg := range trk.Segs {
			trkPts := make([]gpxtypes.TrkPt, 0, len(gpx.Tracks))
			for _, pt := range seg.Points {
				if pt.Time != nil {
					var wantedTime = *pt.Time

					// Rework time if desired.
					if !disableRewriting {
						originalTime, _ := time.Parse(time.RFC3339, *pt.Time)
						syd, _ := toZone(originalTime)
						wantedTime = syd.Format(time.RFC3339)
					}

					// Rewrite output
					trkPts = append(trkPts, gpxtypes.TrkPt{
						XMLName:       pt.XMLName,
						Lat:           pt.Lat,
						Lon:           pt.Lon,
						Elevation:     pt.Elevation,
						Time:          &wantedTime,
						MagVar:        pt.MagVar,
						GeoIDHeight:   pt.GeoIDHeight,
						Name:          pt.Name,
						Cmt:           pt.Cmt,
						Desc:          pt.Desc,
						Src:           pt.Src,
						Link:          pt.Link,
						Sym:           pt.Sym,
						Type:          pt.Type,
						Fix:           pt.Fix,
						Sat:           pt.Sat,
						HDOP:          pt.HDOP,
						VDOP:          pt.VDOP,
						PDOP:          pt.PDOP,
						AgeOfDGPSData: pt.AgeOfDGPSData,
						DGPSID:        pt.DGPSID,
						Extensions:    pt.Extensions,
					})
				}
			}

			// Replace existing trkpts.
			if !disableRewriting {
				gpx.Tracks[i].Segs[n].Points = trkPts
			}
		}
	}
}

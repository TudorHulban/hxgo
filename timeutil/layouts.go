package timeutil

import "time"

const (
	LayoutDate         = "2006-01-02"                    //𒁛-MM-DD
	LayoutDateTime     = "2006-01-02T15:04"              //𒁛-MM-DDTHH:MM
	LayoutDateTimeSec  = "2006-01-02T15:04:05"           //𒁛-MM-DDTHH:MM:SS
	LayoutDateTimeMs   = "2006-01-02T15:04:05.000"       //𒁛-MM-DDTHH:MM:SS.SSS (milliseconds)
	LayoutDateTimeNano = "2006-01-02T15:04:05.000000000" //𒁛-MM-DDTHH:MM:SS.SSSSSSSSS (nanoseconds)

	LayoutRFC3339     = time.RFC3339     // "2006-01-02T15:04:05Z07:00"
	LayoutRFC3339Nano = time.RFC3339Nano // "2006-01-02T15:04:05.999999999Z07:00"
	LayoutRFC822      = time.RFC822      // "02 Jan 06 15:04 MST"
	LayoutRFC822Z     = time.RFC822Z     // "02 Jan 06 15:04 -0700" // With numeric zone
	LayoutRFC1123     = time.RFC1123     // "Mon, 02 Jan 2006 15:04:05 MST"
	LayoutRFC1123Z    = time.RFC1123Z    // "Mon, 02 Jan 2006 15:04:05 -0700"
	LayoutANSIC       = time.ANSIC       // "Mon Jan _2 15:04:05 2006"
	LayoutUnixDate    = time.UnixDate    // "Mon Jan _2 15:04:05 MST 2006"
	LayoutRubyDate    = time.RubyDate    // "Mon Jan 02 15:04:05 -0700 2006"
	LayoutKitchen     = time.Kitchen     // "3:04PM"
	LayoutStamp       = time.Stamp       // "Jan _2 15:04:05"
	LayoutStampMilli  = time.StampMilli  // "Jan _2 15:04:05.000"
	LayoutStampMicro  = time.StampMicro  // "Jan _2 15:04:05.000000"
	LayoutStampNano   = time.StampNano   // "Jan _2 15:04:05.000000000"
)

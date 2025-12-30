package request

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/timeutil"
)

func TestExtractOptionalTimestamp(t *testing.T) {
	data := RequestData{
		Content: map[string]string{
			"time": "2024-09-16T18:16",
		},
	}

	fmt.Println(
		data.ExtractOptionalTimestampInLocation(
			"time",
			timeutil.LayoutDateTime,
			timeutil.LocationRomania,
		),

		data.ExtractOptionalTimestampInLocation(
			"time",
			timeutil.LayoutDateTime,
			timeutil.LocationRomania,
		).Time.UnixNano(),
	)
}

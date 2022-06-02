package wgnTime

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://ikso.us/posts/unmarshal-timestamp-as-time/

// UnixTime is our magic type
type UnixTime struct {
	time.Time
}

// UnmarshalJSON is the method that satisfies the Unmarshaller interface
func (uTime *UnixTime) UnmarshalJSON(data []byte) error {
	var timestamp int64
	err := json.Unmarshal(data, &timestamp)
	if err != nil {
		return err
	}
	uTime.Time = time.Unix(timestamp, 0)
	return nil
}

// MarshalJSON turns our time.Time back into an int
func (uTime UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", uTime.Time.Unix())), nil
}

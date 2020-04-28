package snowflake

import (
	"strconv"
	"time"
)

var epoch = time.Unix(1420070400, 0)

// T represents an ID for an object, unique within all of Discord.
type T uint64

// Parse parses a Discord snowflake.
func Parse(s string) (T, error) {
	t, err := strconv.ParseUint(s, 10, 64)
	return T(t), err
}

// Time retrieves and returns the snowflake's timestamp.
func (s T) Time() time.Time {
	return epoch.Add(time.Millisecond * time.Duration(s>>22))
}

// Stamp returns a copy of the snowflake stamped with the given timestamp.
func (s T) Stamp(t time.Time) T {
	ms := t.Sub(epoch) / time.Millisecond
	return s | T(ms<<22)
}

// Worker retrieves and returns the snowflake's internal five-bit origin worker
// ID.
func (s T) Worker() uint8 {
	return uint8((s & 0x3e0000) >> 17)
}

// Process retrieves and returns the snowflake's internal five-bit origin PID.
func (s T) Process() uint8 {
	return uint8((s & 0x1f000) >> 12)
}

// Seq returns the process seqnum: for every object created by the origin
// process this number is incremented.
func (s T) Seq() uint16 {
	return uint16(s & 0xfff)
}

func (s T) String() string {
	return strconv.FormatUint(s, 10)
}

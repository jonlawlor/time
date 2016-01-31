package ts

import (
	"fmt"
	"time"
)

// TimeSpan is a period of time with a beginning and an end.
type TimeSpan struct {
	Start time.Time
	End   time.Time

	// I use the convoluted "NotStartInclusive" so that the zero value for TimeSpan's
	// bounds is a sensible default, and callers can just use ts.TimeSpan{Start: foo, End: bar}
	NotStartInclusive bool
	EndInclusive      bool
}

// String returns the TimeSpan formatted using the format string
//	"2006-01-02 15:04:05.999999999 -0700 MST", same as time.Time
func (ts1 TimeSpan) String() string {
	return ts1.Format("2006-01-02 15:04:05.999999999 -0700 MST")
}

// Format returns the TimeSpan formatted according to the provided format, with
// bounds from [] and () when they are inclusive or exclusive, respectively.
func (ts1 TimeSpan) Format(layout string) string {
	var f string
	switch {
	case !ts1.NotStartInclusive && !ts1.EndInclusive:
		f = "[%s, %s)"
	case ts1.NotStartInclusive && ts1.EndInclusive:
		f = "(%s, %s]"
	case ts1.NotStartInclusive && !ts1.EndInclusive:
		f = "(%s, %s)"
	case !ts1.NotStartInclusive && ts1.EndInclusive:
		f = "[%s, %s]"
	}
	return fmt.Sprintf(f, ts1.Start.Format(layout), ts1.End.Format(layout))
}

// Copy produces a timespan that is identical to the input.
func (ts1 TimeSpan) Copy() TimeSpan {
	return TimeSpan{
		Start:             ts1.Start,
		End:               ts1.End,
		NotStartInclusive: ts1.NotStartInclusive,
		EndInclusive:      ts1.EndInclusive,
	}
}

// Union returns the smallest TimeSpan that contains both input TimeSpans.
func (ts1 TimeSpan) Union(ts2 TimeSpan) (ts3 TimeSpan) {
	switch {
	case ts1.Start.Before(ts2.Start):
		ts3.Start = ts1.Start
		ts3.NotStartInclusive = ts1.NotStartInclusive
	case ts2.Start.Before(ts1.Start):
		ts3.Start = ts2.Start
		ts3.NotStartInclusive = ts2.NotStartInclusive
	default:
		ts3.Start = ts1.Start
		ts3.NotStartInclusive = ts1.NotStartInclusive && ts2.NotStartInclusive
	}

	switch {
	case ts1.End.After(ts2.End):
		ts3.End = ts1.End
		ts3.EndInclusive = ts1.EndInclusive
	case ts2.End.After(ts1.End):
		ts3.End = ts2.End
		ts3.EndInclusive = ts2.EndInclusive
	default:
		ts3.End = ts1.End
		ts3.EndInclusive = ts1.EndInclusive || ts2.EndInclusive
	}
	return
}

// Intersect returns the TimeSpan that is common to both inputs.
func (ts1 TimeSpan) Intersect(ts2 TimeSpan) (ts3 TimeSpan) {
	switch {
	case ts1.Start.Before(ts2.Start):
		ts3.Start = ts2.Start
		ts3.NotStartInclusive = ts2.NotStartInclusive
	case ts2.Start.Before(ts1.Start):
		ts3.Start = ts1.Start
		ts3.NotStartInclusive = ts1.NotStartInclusive
	default:
		ts3.Start = ts1.Start
		ts3.NotStartInclusive = ts1.NotStartInclusive || ts2.NotStartInclusive
	}

	switch {
	case ts1.End.After(ts2.End):
		ts3.End = ts2.End
		ts3.EndInclusive = ts2.EndInclusive
	case ts2.End.After(ts1.End):
		ts3.End = ts1.End
		ts3.EndInclusive = ts1.EndInclusive
	default:
		ts3.End = ts1.End
		ts3.EndInclusive = ts1.EndInclusive && ts2.EndInclusive
	}
	return
}

// Diff returns the TimeSpans that are in the first but not in the second, and
// that are in the second but not the first.
func (ts1 TimeSpan) Diff(ts2 TimeSpan) (TimeSpan, TimeSpan) {
	if ts1.End.Before(ts2.Start) || ts1.Start.After(ts2.End) {
		return ts1.Copy(), ts2.Copy()
	}
	var ts3, ts4 TimeSpan
	if ts1.Start.Before(ts2.Start) {
		ts3.Start = ts1.Start
		ts3.NotStartInclusive = ts1.NotStartInclusive
		ts3.End = ts2.Start
		ts3.EndInclusive = ts2.NotStartInclusive

		ts4.Start = ts1.End
		ts4.NotStartInclusive = ts1.EndInclusive
		ts4.End = ts2.End
		ts4.EndInclusive = ts2.EndInclusive
		return ts3, ts4
	}
	ts4.Start = ts2.Start
	ts4.NotStartInclusive = ts2.NotStartInclusive
	ts4.End = ts1.Start
	ts4.EndInclusive = ts1.NotStartInclusive

	ts3.Start = ts2.End
	ts3.NotStartInclusive = ts2.EndInclusive
	ts3.End = ts1.End
	ts3.EndInclusive = ts1.EndInclusive
	return ts3, ts4
}

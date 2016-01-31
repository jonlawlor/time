package ts

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		// you maniacs you blew it up
		panic(err)
	}
	Nov0920148AMNYC := time.Date(2014, 11, 9, 8, 0, 0, 0, loc)
	Nov0920158AMNYC := time.Date(2015, 11, 9, 8, 0, 0, 0, loc)
	for i, test := range []struct {
		ts     TimeSpan
		expect string
	}{
		{
			ts:     TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC},
			expect: "[2014-11-09 08:00:00 -0500 EST, 2015-11-09 08:00:00 -0500 EST)",
		}, {
			ts:     TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, NotStartInclusive: true},
			expect: "(2014-11-09 08:00:00 -0500 EST, 2015-11-09 08:00:00 -0500 EST)",
		}, {
			ts:     TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, EndInclusive: true},
			expect: "[2014-11-09 08:00:00 -0500 EST, 2015-11-09 08:00:00 -0500 EST]",
		}, {
			ts:     TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, NotStartInclusive: true, EndInclusive: true},
			expect: "(2014-11-09 08:00:00 -0500 EST, 2015-11-09 08:00:00 -0500 EST]",
		}} {

		if got := test.ts.String(); got != test.expect {
			t.Errorf("%d: expected: %s, got %s", i, test.expect, got)
		}
	}
}

func TestUnion(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		// you maniacs you blew it up
		panic(err)
	}
	Nov0920148AMNYC := time.Date(2014, 11, 9, 8, 0, 0, 0, loc)
	Nov0920158AMNYC := time.Date(2015, 11, 9, 8, 0, 0, 0, loc)
	Apr1720147AMNYC := time.Date(2014, 4, 17, 7, 0, 0, 0, loc)
	Apr1720157AMNYC := time.Date(2015, 4, 17, 7, 0, 0, 0, loc)
	for i, test := range []struct {
		ts1, ts2, expect TimeSpan
	}{
		{
			ts1:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC},
			ts2:    TimeSpan{Start: Apr1720147AMNYC, End: Apr1720157AMNYC},
			expect: TimeSpan{Start: Apr1720147AMNYC, End: Nov0920158AMNYC},
		}, {
			ts1:    TimeSpan{Start: Apr1720147AMNYC, End: Apr1720157AMNYC},
			ts2:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC},
			expect: TimeSpan{Start: Apr1720147AMNYC, End: Nov0920158AMNYC},
		}, {
			ts1:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, NotStartInclusive: true},
			ts2:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, EndInclusive: true},
			expect: TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, EndInclusive: true},
		}} {
		res := test.ts1.Union(test.ts2)
		if got := res.String(); got != test.expect.String() {
			t.Errorf("%d: expected: %s, got %s", i, test.expect.String(), got)
		}
	}
}

func TestIntersect(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		// you maniacs you blew it up
		panic(err)
	}
	Nov0920148AMNYC := time.Date(2014, 11, 9, 8, 0, 0, 0, loc)
	Nov0920158AMNYC := time.Date(2015, 11, 9, 8, 0, 0, 0, loc)
	Apr1720147AMNYC := time.Date(2014, 4, 17, 7, 0, 0, 0, loc)
	Apr1720157AMNYC := time.Date(2015, 4, 17, 7, 0, 0, 0, loc)
	for i, test := range []struct {
		ts1, ts2, expect TimeSpan
	}{
		{
			ts1:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC},
			ts2:    TimeSpan{Start: Apr1720147AMNYC, End: Apr1720157AMNYC},
			expect: TimeSpan{Start: Nov0920148AMNYC, End: Apr1720157AMNYC},
		}, {
			ts1:    TimeSpan{Start: Apr1720147AMNYC, End: Apr1720157AMNYC},
			ts2:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC},
			expect: TimeSpan{Start: Nov0920148AMNYC, End: Apr1720157AMNYC},
		}, {
			ts1:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, NotStartInclusive: true},
			ts2:    TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, EndInclusive: true},
			expect: TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC, NotStartInclusive: true},
		}} {
		res := test.ts1.Intersect(test.ts2)
		if got := res.String(); got != test.expect.String() {
			t.Errorf("%d: expected: %s, got %s", i, test.expect.String(), got)
		}
	}
}

func TestDiff(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		// you maniacs you blew it up
		panic(err)
	}
	Nov0920148AMNYC := time.Date(2014, 11, 9, 8, 0, 0, 0, loc)
	Nov0920158AMNYC := time.Date(2015, 11, 9, 8, 0, 0, 0, loc)
	Apr1720147AMNYC := time.Date(2014, 4, 17, 7, 0, 0, 0, loc)
	Apr1720157AMNYC := time.Date(2015, 4, 17, 7, 0, 0, 0, loc)
	for i, test := range []struct {
		ts1, ts2, expect1, expect2 TimeSpan
	}{
		{
			ts1:     TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC},
			ts2:     TimeSpan{Start: Apr1720147AMNYC, End: Apr1720157AMNYC},
			expect1: TimeSpan{Start: Apr1720157AMNYC, End: Nov0920158AMNYC},
			expect2: TimeSpan{Start: Apr1720147AMNYC, End: Nov0920148AMNYC},
		}, {
			ts2:     TimeSpan{Start: Nov0920148AMNYC, End: Nov0920158AMNYC},
			ts1:     TimeSpan{Start: Apr1720147AMNYC, End: Apr1720157AMNYC},
			expect2: TimeSpan{Start: Apr1720157AMNYC, End: Nov0920158AMNYC},
			expect1: TimeSpan{Start: Apr1720147AMNYC, End: Nov0920148AMNYC},
		}, {
			ts1:     TimeSpan{Start: Apr1720147AMNYC, End: Nov0920148AMNYC},
			ts2:     TimeSpan{Start: Apr1720157AMNYC, End: Nov0920158AMNYC},
			expect1: TimeSpan{Start: Apr1720147AMNYC, End: Nov0920148AMNYC},
			expect2: TimeSpan{Start: Apr1720157AMNYC, End: Nov0920158AMNYC},
		}} {
		res1, res2 := test.ts1.Diff(test.ts2)
		if got := res1.String(); got != test.expect1.String() {
			t.Errorf("%d: expected first output: %s, got %s", i, test.expect1.String(), got)
		}
		if got := res2.String(); got != test.expect2.String() {
			t.Errorf("%d: expected second output: %s, got %s", i, test.expect2.String(), got)
		}
	}
}

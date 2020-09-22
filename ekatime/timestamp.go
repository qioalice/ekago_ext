// Copyright © 2020. All rights reserved.
// Author: Ilya Stroy.
// Contacts: qioalice@gmail.com, https://github.com/qioalice
// License: https://opensource.org/licenses/MIT

package ekatime

//goland:noinspection GoSnakeCaseUsage
import (
	"time"

	ekatime_orig "github.com/qioalice/ekago/v2/ekatime"
)

type (
	// Timestamp is the same as ekatime.Timestamp but with supporting go-pg (v10).
	//
	// Read more:
	// https://github.com/qioalice/ekago/ekatime/timestamp_encode.go ,
	// https://github.com/go-pg/pg ,
	// https://github.com/go-pg/pg/blob/v10/example_custom_test.go .
	Timestamp ekatime_orig.Timestamp
)

// WrapTimestamp returns a Timestamp object as modified ekatime.Timestamp object
// for being able to use it with go-pg.
//
// See also: WrapTimestampPtr().
func WrapTimestamp(ts ekatime_orig.Timestamp) Timestamp {
	return Timestamp(ts)
}

// WrapTimestampPtr returns a Timestamp object by ptr as modified ekatime.Time object
// for being able to use it with go-pg.
//
// See also: WrapTimestamp().
func WrapTimestampPtr(ts *ekatime_orig.Timestamp) *Timestamp {
	return (*Timestamp)(ts)
}

func (ts Timestamp) ToOrig() ekatime_orig.Timestamp {
	return ekatime_orig.Timestamp(ts)
}

func (ts *Timestamp) ToOrigPtr() *ekatime_orig.Timestamp {
	return (*ekatime_orig.Timestamp)(ts)
}

func (ts Timestamp) I64() int64 {
	return ts.ToOrig().I64()
}

func (ts Timestamp) Std() time.Time {
	return ts.ToOrig().Std()
}

func (ts Timestamp) Date() Date {
	return WrapDate(ts.ToOrig().Date())
}

func (ts Timestamp) Time() Time {
	return WrapTime(ts.ToOrig().Time())
}

func (ts Timestamp) Year() Year {
	return ts.ToOrig().Year()
}

func (ts Timestamp) Month() Month {
	return ts.ToOrig().Month()
}

func (ts Timestamp) Day() Day {
	return ts.ToOrig().Day()
}

func (ts Timestamp) Hour() Hour {
	return ts.ToOrig().Hour()
}

func (ts Timestamp) Minute() Minute {
	return ts.ToOrig().Minute()
}

func (ts Timestamp) Second() Second {
	return ts.ToOrig().Second()
}

func (ts Timestamp) Split() (d Date, t Time) {
	return WrapDate(ts.ToOrig().Date()), WrapTime(ts.ToOrig().Time())
}

func Now() Timestamp {
	return WrapTimestamp(ekatime_orig.Now())
}

func UnixFrom(y Year, m Month, d Day, hh Hour, mm Minute, ss Second) Timestamp {
	return WrapTimestamp(ekatime_orig.UnixFrom(y, m, d, hh, mm, ss))
}

func UnixFromStd(t time.Time) Timestamp {
	return WrapTimestamp(ekatime_orig.UnixFromStd(t))
}

func (ts Timestamp) BeginningOfDay() Timestamp {
	return WrapTimestamp(ts.ToOrig().BeginningOfDay())
}

func (ts Timestamp) EndOfDay() Timestamp {
	return WrapTimestamp(ts.ToOrig().EndOfDay())
}

func (ts Timestamp) BeginningAndEndOfDay() TimestampPair {
	return WrapTimestampPair(ts.ToOrig().BeginningAndEndOfDay())
}

func (ts Timestamp) BeginningOfMonth() Timestamp {
	return WrapTimestamp(ts.ToOrig().BeginningOfMonth())
}

func (ts Timestamp) EndOfMonth() Timestamp {
	return WrapTimestamp(ts.ToOrig().EndOfMonth())
}

func (ts Timestamp) BeginningAndEndOfMonth() TimestampPair {
	return WrapTimestampPair(ts.ToOrig().BeginningAndEndOfMonth())
}

func (ts Timestamp) BeginningOfYear() Timestamp {
	return WrapTimestamp(ts.ToOrig().BeginningOfYear())
}

func (ts Timestamp) EndOfYear() Timestamp {
	return WrapTimestamp(ts.ToOrig().EndOfYear())
}

func (ts Timestamp) BeginningAndEndOfYear() TimestampPair {
	return WrapTimestampPair(ts.ToOrig().BeginningAndEndOfYear())
}

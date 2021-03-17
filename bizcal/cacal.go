package bizcal

import (
	"time"
)

/*
This part of the code is a golang adaptation of the QuantLib project
Calendar implementation, originally written in C++

It uses the QuantLib license as described and linked below
*/

/*
 Copyright (C) 2006 Piter Dias
 Copyright (C) 2011 StatPro Italia srl

 This file is part of QuantLib, a free-software/open-source library
 for financial quantitative analysts and developers - http://quantlib.org/

 QuantLib is free software: you can redistribute it and/or modify it
 under the terms of the QuantLib license.  You should have received a
 copy of the license along with this program; if not, please email
 <quantlib-dev@lists.sf.net>. The license is also available online at
 <http://quantlib.org/license.shtml>.

 This program is distributed in the hope that it will be useful, but WITHOUT
 ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
 FOR A PARTICULAR PURPOSE.  See the license for more details.
*/

/*
End QuantLib license text
*/

//CACal for all Canadian holidays
//has all BasicCal methods
type CACal struct {
	BasicCal
}

//IsNewYearsDay checks if a particular day is on new year's day
func (cal CACal) IsNewYearsDay(y int, m time.Month, d int, w time.Weekday) bool {
	// New Year's Day
	// Could be following Monday
	// Or preceeding Friday
	return ((d == 1 || ((d == 2 || d == 3) && w == time.Monday)) && m == time.January)
}

//IsFamilyDay checks if a particular day is family day
func (cal CACal) IsFamilyDay(y int, m time.Month, d int, w time.Weekday) bool {
	// Family Day (third Monday in February, since 2008)
	return ((d >= 15 && d <= 21) && w == time.Monday && m == time.February && y >= 2008)
}

//IsGoodFriday checks for Good Friday
func (cal CACal) IsGoodFriday(y int, dd int) bool {
	// get Easter Monday for that year first
	em := cal.EasterMonday(y)
	// dd is day in year, the ddth day of year y
	return dd == em-3
}

//IsVictoriaDay checks for Victoria Day
func (cal CACal) IsVictoriaDay(y int, m time.Month, d int, w time.Weekday) bool {
	// The Monday on or preceding 24 May (Victoria Day)
	return (d > 17 && d <= 24 && w == time.Monday && m == time.May)
}

//IsCanadaDay checks for Canada Day
func (cal CACal) IsCanadaDay(y int, m time.Month, d int, w time.Weekday) bool {
	// July 1st, possibly moved to Monday (Canada Day)
	return ((d == 1 || ((d == 2 || d == 3) && w == time.Monday)) && m == time.July)
}

//IsProvincialHoliday checks for provincial holiday
func (cal CACal) IsProvincialHoliday(y int, m time.Month, d int, w time.Weekday) bool {
	// first Monday of August (Provincial Holiday)
	return (d <= 7 && w == time.Monday && m == time.August)
}

//IsLaborDay checks for Labor Day
func (cal CACal) IsLaborDay(y int, m time.Month, d int, w time.Weekday) bool {
	// first Monday of September (Labor Day)
	return (d <= 7 && w == time.Monday && m == time.September)
}

//IsThanksgiving checks for Thanksgiving
func (cal CACal) IsThanksgiving(y int, m time.Month, d int, w time.Weekday) bool {
	// second Monday of October (Thanksgiving Day)
	return (d > 7 && d <= 14 && w == time.Monday && m == time.October)
}

//IsRemeberanceDay checks for Rememberance Day
func (cal CACal) IsRememberanceDay(y int, m time.Month, d int, w time.Weekday) bool {
	// November 11th (possibly moved to Monday)
	return ((d == 11 || ((d == 12 || d == 13) && w == time.Monday)) && m == time.November)
}

//IsChristmas checks for Christmas
func (cal CACal) IsChristmas(y int, m time.Month, d int, w time.Weekday) bool {
	// Christmas (possibly moved to Monday or Tuesday)
	return (m == time.December &&
		(d == 25 || (d == 27 && (w == time.Monday || w == time.Tuesday))))
}

//IsBoxingDay checks for Boxing Day
func (cal CACal) IsBoxingDay(y int, m time.Month, d int, w time.Weekday) bool {
	// Boxing Day (possibly moved to Monday or Tuesday)
	return (m == time.December &&
		(d == 26 || (d == 28 && (w == time.Monday || w == time.Tuesday))))
}

//CASettleCal, calendar for CA Settlement
//has all CACal methods
//It also satisfies BizCal interface
type CASettleCal struct {
	CACal
}

//IsBusinessDay checks for business day according to US Settlement Calendar
func (cal CASettleCal) IsBusinessDay(t time.Time) bool {
	if cal.IsWeekend(t) {
		return false
	}

	y, m, d := t.Date()
	w := t.Weekday()
	dd := t.YearDay()

	if cal.IsNewYearsDay(y, m, d, w) ||
		cal.IsFamilyDay(y, m, d, w) ||
		cal.IsGoodFriday(y, dd) ||
		cal.IsVictoriaDay(y, m, d, w) ||
		cal.IsCanadaDay(y, m, d, w) ||
		cal.IsProvincialHoliday(y, m, d, w) ||
		cal.IsLaborDay(y, m, d, w) ||
		cal.IsThanksgiving(y, m, d, w) ||
		cal.IsRememberanceDay(y, m, d, w) ||
		cal.IsChristmas(y, m, d, w) ||
		cal.IsBoxingDay(y, m, d, w) {
		// holidays
		return false
	}

	return true
}

//TSXCal, calendar for Toronto Stock Exchange
//has all CACal methods
//It also satisfies BizCal interface
type TSXCal struct {
	CACal
}

//IsBusinessDay checks for business day according to TSX Calendar
func (cal TSXCal) IsBusinessDay(t time.Time) bool {
	if cal.IsWeekend(t) {
		return false
	}

	y, m, d := t.Date()
	w := t.Weekday()
	dd := t.YearDay()

	if cal.IsNewYearsDay(y, m, d, w) ||
		cal.IsFamilyDay(y, m, d, w) ||
		cal.IsGoodFriday(y, dd) ||
		cal.IsVictoriaDay(y, m, d, w) ||
		cal.IsCanadaDay(y, m, d, w) ||
		cal.IsProvincialHoliday(y, m, d, w) ||
		cal.IsLaborDay(y, m, d, w) ||
		cal.IsThanksgiving(y, m, d, w) ||
		cal.IsRememberanceDay(y, m, d, w) ||
		cal.IsChristmas(y, m, d, w) ||
		cal.IsBoxingDay(y, m, d, w) {
		// holidays
		return false
	}

	return true
}

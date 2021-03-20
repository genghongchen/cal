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

//USCal for all US holidays
//has all BasicCal methods
type USCal struct {
	BasicCal
}

//IsNewYearsDay checks if a particular day is on new year's day
func (cal USCal) IsNewYearsDay(y int, m time.Month, d int, w time.Weekday) bool {
	// New Year's Day
	// Could be following Monday
	// Or preceeding Friday
	if (d == 1 && m == time.January) ||
		(d == 2 && m == time.January && w == time.Monday) {
		return true
	}

	return false
}

//IsMLKDay checks for MLK Day
func (cal USCal) IsMLKDay(y int, m time.Month, d int, w time.Weekday) bool {
	// MLK Day (third Monday in Jan)
	if (d >= 15 && d <= 21) && m == time.January && w == time.Monday {
		return true
	}

	return false
}

//IsPresidentsDay checks for President's Day
func (cal USCal) IsPresidentsDay(y int, m time.Month, d int, w time.Weekday) bool {
	if y >= 1971 {
		// third Monday in February
		return (d >= 15 && d <= 21) && m == time.February && w == time.Monday
	} else {
		// February 22nd, or as adjusted
		return m == time.February && (d == 22 ||
			(d == 23 && w == time.Monday) ||
			(d == 21 && w == time.Friday))
	}
}

//IsGoodFriday checks for Good Friday
func (cal USCal) IsGoodFriday(y int, dd int) bool {
	// get Easter Monday for that year first
	em := cal.EasterMonday(y)
	// dd is day in year, the ddth day of year y
	return dd == em-3
}

//IsMemorialDay checks for Memorial Day
func (cal USCal) IsMemorialDay(y int, m time.Month, d int, w time.Weekday) bool {
	if y >= 1971 {
		// last Monday in May
		return d >= 25 && w == time.Monday && m == time.May
	} else {
		// May 30th, or as adjusted
		return m == time.May && (d == 30 ||
			(d == 31 && w == time.Monday) ||
			(d == 29 && w == time.Friday))
	}
}

//IsIndependenceDay checks for 4th of July
func (cal USCal) IsIndependenceDay(y int, m time.Month, d int, w time.Weekday) bool {
	return m == time.July && (d == 4 ||
		(d == 5 && w == time.Monday) ||
		(d == 3 && w == time.Friday))
}

//IsLaborDay checks for Labor Day
func (cal USCal) IsLaborDay(y int, m time.Month, d int, w time.Weekday) bool {
	// first Monday in September
	return m == time.September && d <= 7 && w == time.Monday
}

//IsColumbusDay checks for Columbus Day
func (cal USCal) IsColumbusDay(y int, m time.Month, d int, w time.Weekday) bool {
	// second Monday in October
	return m == time.October && y >= 1971 &&
		(d >= 8 && d <= 14) && w == time.Monday
}

//IsVeteransDay checks for Veteran's Day
func (cal USCal) IsVeteransDay(y int, m time.Month, d int, w time.Weekday) bool {
	if y <= 1970 || y >= 1978 {
		// November 11th, as adjusted
		return m == time.November && (d == 11 ||
			(d == 12 && w == time.Monday) ||
			(d == 10 && w == time.Friday))
	} else {
		// fourth Monday in October
		return m == time.October && (d >= 22 && d <= 28) && w == time.Monday
	}
}

//IsVeteransDayNoSaturday checks for Veteran's Day
func (cal USCal) IsVeteransDayNoSaturday(y int, m time.Month, d int, w time.Weekday) bool {
	if y <= 1970 || y >= 1978 {
		// November 11th, as adjusted, but no Saturday to Friday adjustment
		return m == time.November && (d == 11 ||
			(d == 12 && w == time.Monday))
	} else {
		// fourth Monday in October
		return m == time.October && (d >= 22 && d <= 28) && w == time.Monday
	}
}

//IsThanksgiving checks for Thanksgiving Day
func (cal USCal) IsThanksgiving(y int, m time.Month, d int, w time.Weekday) bool {
	// Thanksgiving Day (fourth Thursday in November)
	return m == time.November && w == time.Thursday && (d >= 22 && d <= 28)
}

//IsChristmas checks for Christmas Day
func (cal USCal) IsChristmas(y int, m time.Month, d int, w time.Weekday) bool {
	return m == time.December && (d == 25 ||
		(d == 26 && w == time.Monday) ||
		(d == 24 && w == time.Friday))
}

//BizCal interface, Business calendar
type BizCal interface {
	BaseCal
	IsBusinessDay(t time.Time) bool
}

//USSettleCal, calendar for US Settlement
//has all USCal methods
//It also satisfies BizCal interface
type USSettleCal struct {
	USCal
}

//IsBusinessDay checks for business day according to US Settlement Calendar
func (cal USSettleCal) IsBusinessDay(t time.Time) bool {
	if cal.IsWeekend(t) {
		return false
	}

	y, m, d := t.Date()
	w := t.Weekday()

	if cal.IsNewYearsDay(y, m, d, w) ||
		// only for US Settlement Calendar
		// Preceeding 12/31 if 1/1 is on Saturday
		(d == 31 && m == time.December && w == time.Friday) ||
		cal.IsMLKDay(y, m, d, w) ||
		cal.IsPresidentsDay(y, m, d, w) ||
		cal.IsMemorialDay(y, m, d, w) ||
		cal.IsIndependenceDay(y, m, d, w) ||
		cal.IsLaborDay(y, m, d, w) ||
		cal.IsColumbusDay(y, m, d, w) ||
		cal.IsVeteransDay(y, m, d, w) ||
		cal.IsThanksgiving(y, m, d, w) ||
		cal.IsChristmas(y, m, d, w) {
		// holidays
		return false
	}

	return true
}

//USLiborCal, calendar for US Libor
//has all USCal methods
//It also satisfies BizCal interface
type USLiborCal struct {
	USCal
}

//IsBusinessDay checks for business day according to US Settlement Calendar
func (cal USLiborCal) IsBusinessDay(t time.Time) bool {
	if cal.IsWeekend(t) {
		return false
	}

	y, m, d := t.Date()
	w := t.Weekday()

	// Since 2015 Independence Day only impacts Libor if it falls
	// on a weekday
	if m == time.July && y >= 2015 && d != 4 {
		return true
	}

	if cal.IsNewYearsDay(y, m, d, w) ||
		// only for US Settlement Calendar
		// Preceeding 12/31 if 1/1 is on Saturday
		(d == 31 && m == time.December && w == time.Friday) ||
		cal.IsMLKDay(y, m, d, w) ||
		cal.IsPresidentsDay(y, m, d, w) ||
		cal.IsMemorialDay(y, m, d, w) ||
		cal.IsIndependenceDay(y, m, d, w) ||
		cal.IsLaborDay(y, m, d, w) ||
		cal.IsColumbusDay(y, m, d, w) ||
		cal.IsVeteransDay(y, m, d, w) ||
		cal.IsThanksgiving(y, m, d, w) ||
		cal.IsChristmas(y, m, d, w) {
		// holidays
		return false
	}

	return true
}

//USGovBondCal, calendar for US Government bonds
//has all USCal methods
//It also satisfies BizCal interface
type USGovBondCal struct {
	USCal
}

//IsBusinessDay checks for business day according to US Settlement Calendar
func (cal USGovBondCal) IsBusinessDay(t time.Time) bool {
	if cal.IsWeekend(t) {
		return false
	}

	y, m, d := t.Date()
	w := t.Weekday()
	dd := t.YearDay()

	if cal.IsNewYearsDay(y, m, d, w) ||
		cal.IsMLKDay(y, m, d, w) ||
		cal.IsPresidentsDay(y, m, d, w) ||
		(y != 2015 && cal.IsGoodFriday(y, dd)) ||
		cal.IsMemorialDay(y, m, d, w) ||
		cal.IsIndependenceDay(y, m, d, w) ||
		cal.IsLaborDay(y, m, d, w) ||
		cal.IsColumbusDay(y, m, d, w) ||
		cal.IsVeteransDayNoSaturday(y, m, d, w) ||
		cal.IsThanksgiving(y, m, d, w) ||
		cal.IsChristmas(y, m, d, w) {
		// holidays
		return false
	}

	// Special closings
	if (y == 2018 && m == time.December && d == 5) || // President Bush's Funeral
		// Hurricane Sandy
		(y == 2012 && m == time.October && (d == 30)) ||
		// President Reagan's funeral
		(y == 2004 && m == time.June && d == 11) {
		return false
	}

	return true
}

//USFedCal, calendar for US Settlement
//has all USCal methods
//It also satisfies BizCal interface
type USFedCal struct {
	USCal
}

//IsBusinessDay checks for business day according to Federal Reserve Calendar
func (cal USFedCal) IsBusinessDay(t time.Time) bool {
	if cal.IsWeekend(t) {
		return false
	}

	y, m, d := t.Date()
	w := t.Weekday()

	if cal.IsNewYearsDay(y, m, d, w) ||
		cal.IsMLKDay(y, m, d, w) ||
		cal.IsPresidentsDay(y, m, d, w) ||
		cal.IsMemorialDay(y, m, d, w) ||
		// a little bit different for independence day
		// no 7/3 holiday if 7/4 is a Saturday
		(m == time.July && (d == 4 || (d == 5 && w == time.Monday))) ||
		cal.IsLaborDay(y, m, d, w) ||
		cal.IsColumbusDay(y, m, d, w) ||
		cal.IsVeteransDayNoSaturday(y, m, d, w) ||
		cal.IsThanksgiving(y, m, d, w) ||
		// subtle difference for Christmas too
		// no 12/24 holiday if 12/25 is a Saturday
		(m == time.December && (d == 25 || (d == 26 && w == time.Monday))) {
		// holidays
		return false
	}

	return true
}

//NYSECal, calendar for New York Stock Exchange
//has all USCal methods
//It also satisfies BizCal interface
type NYSECal struct {
	USCal
}

//IsBusinessDay checks for business day according to NYSE calendar
func (cal NYSECal) IsBusinessDay(t time.Time) bool {
	if cal.IsWeekend(t) {
		return false
	}

	y, m, d := t.Date()
	w := t.Weekday()
	dd := t.YearDay()

	if cal.IsNewYearsDay(y, m, d, w) ||
		(y >= 1998 && cal.IsMLKDay(y, m, d, w)) ||
		cal.IsPresidentsDay(y, m, d, w) ||
		cal.IsGoodFriday(y, dd) ||
		cal.IsMemorialDay(y, m, d, w) ||
		cal.IsIndependenceDay(y, m, d, w) ||
		cal.IsLaborDay(y, m, d, w) ||
		cal.IsThanksgiving(y, m, d, w) ||
		cal.IsChristmas(y, m, d, w) {
		// holidays
		return false
	}

	if (y <= 1968 || (y <= 1980 && y%4 == 0)) &&
		m == time.November && d <= 7 && w == time.Tuesday {
		// Presidential election days
		return false
	}

	// Special closings
	if // President Bush's Funeral
	(y == 2018 && m == time.December && d == 5) ||
		// Hurricane Sandy
		(y == 2012 && m == time.October && (d == 29 || d == 30)) ||
		// President Ford's funeral
		(y == 2007 && m == time.January && d == 2) ||
		// President Reagan's funeral
		(y == 2004 && m == time.June && d == 11) ||
		// September 11-14, 2001
		(y == 2001 && m == time.September && (11 <= d && d <= 14)) ||
		// President Nixon's funeral
		(y == 1994 && m == time.April && d == 27) ||
		// Hurricane Gloria
		(y == 1985 && m == time.September && d == 27) ||
		// 1977 Blackout
		(y == 1977 && m == time.July && d == 14) ||
		// Funeral of former President Lyndon B. Johnson.
		(y == 1973 && m == time.January && d == 25) ||
		// Funeral of former President Harry S. Truman
		(y == 1972 && m == time.December && d == 28) ||
		// National Day of Participation for the lunar exploration.
		(y == 1969 && m == time.July && d == 21) ||
		// Funeral of former President Eisenhower.
		(y == 1969 && m == time.March && d == 31) ||
		// Closed all day - heavy snow.
		(y == 1969 && m == time.February && d == 10) ||
		// Day after Independence Day.
		(y == 1968 && m == time.July && d == 5) ||
		// June 12-Dec. 31, 1968
		// Four day week (closed on Wednesdays) - Paperwork Crisis
		(y == 1968 && dd >= 163 && w == time.Wednesday) ||
		// Day of mourning for Martin Luther King Jr.
		(y == 1968 && m == time.April && d == 9) ||
		// Funeral of President Kennedy
		(y == 1963 && m == time.November && d == 25) ||
		// Day before Decoration Day
		(y == 1961 && m == time.May && d == 29) ||
		// Day after Christmas
		(y == 1958 && m == time.December && d == 26) ||
		// Christmas Eve
		((y == 1954 || y == 1956 || y == 1965) && m == time.December && d == 24) {
		return false
	}

	return true
}

// End QuantLib code adaptation

//AdjForBusinessDay take one date and either returns itself
//if it is already a business day
//or returns the next business day
func AdjForBusinessDay(cal BizCal, t time.Time) time.Time {
	rt := t

	for {
		if cal.IsBusinessDay(rt) {
			return rt
		}
		rt = rt.AddDate(0, 0, 1)
	}
}

//NextBusinessDay takes one day and returns
//the next business day after that day
func NextBusinessDay(cal BizCal, t time.Time) time.Time {
	rt := t

	for {
		rt = rt.AddDate(0, 0, 1)
		if cal.IsBusinessDay(rt) {
			return rt
		}
	}
}

//AdjLastBusinessDay take one date and either returns itself
//if it is already a business day
//or returns the last business day
func AdjLastBusinessDay(cal BizCal, t time.Time) time.Time {
	rt := t

	for {
		if cal.IsBusinessDay(rt) {
			return rt
		}
		rt = rt.AddDate(0, 0, -1)
	}
}

//PrevBusinessDay takes one day and returns
//the previous business day before that day
func PrevBusinessDay(cal BizCal, t time.Time) time.Time {
	rt := t

	for {
		rt = rt.AddDate(0, 0, -1)
		if cal.IsBusinessDay(rt) {
			return rt
		}
	}
}

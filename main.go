package main

import (
	"fmt"
)

func main() {
	fmt.Print("deteksi kalender\n")
	var tmstmp = *newTimestamp(*newTime(0, 0, 59), *newDate(2024, 3, -1, 1))
	fmt.Print("h-m-s: ", tmstmp.time, " d-d-m-y: ", tmstmp.date)
}

type Timestamp struct {
	time Time
	date Date
}

func newTimestamp(time Time, date Date) *Timestamp {
	tp := new(Timestamp)
	tp.time = *newTime(time.hour, time.minute, time.second)
	tp.date = *newDate(date.year, date.month, date.date, date.day)
	return tp
}

type Time struct {
	hour   int
	minute int
	second int
}

func newTime(hour int, minute int, second int) *Time {
	p := new(Time)
	if hour <= 23 && hour >= 0 {
		p.hour = hour
	}
	if minute <= 59 && minute >= 0 {
		p.minute = minute
	}
	if second <= 59 && second >= 0 {
		p.second = second
	}
	return p
}

type Date struct {
	day   int
	date  int
	month int
	year  int
}

func newDate(y int, month int, date int, day int) *Date {
	d := new(Date)
	d.year = y
	//bulan harus 12, jika tidak:
	if month <= 12 && month > 0 {
		d.month = month
	} else if month > 12 {
		d.year += 1
		d.month = month - 12
	} else {
		d.year -= 1
		d.month = 12 - (month * -1)
	}
	//tanggal berdasar tahun kabisat
	//jika tahun kabisat dan di bulan februari
	if d.year%4 == 0 && d.month == 2 {
		if date > 0 && date <= 29 {
			d.date = date
		} else if date > 29 {
			d.month += 1
			d.date = date - 29
		} else {
			d.month -= 1
			d.date = 31 - (date * -1)
		}
		//jika berada pada tahun kabisat namun tidak pada bulan februari
	} else if d.year%4 == 0 && d.month != 2 {
		//jika pada bulan ganjil {januari, maret, mei, juli}
		if d.month%2 != 0 {
			if date > 0 && date <= 31 {
				d.date = date
			} else if date > 31 {
				d.month += 1
				d.date = date - 31
			} else {
				d.month -= 1
				d.date = 30 - (date * -1)
			}
			//jika pada bulan genap {April, juni, agustus}
		} else {
			if date > 0 && date <= 30 {
				d.date = date
			} else if date > 30 {
				d.month += 1
				d.date = date - 30
			} else {
				d.month -= 1
				d.date = 31 - (date * -1)
			}
		}
		//jika tidak berada pada tahun kabisat tapi di bulan februari
	} else if d.year%4 != 0 && d.month == 2 {
		// februari: 1 - 28
		if date > 0 && date <= 28 {
			d.date = date
		} else if date > 28 {
			d.month += 1
			d.date = date - 28
		} else {
			d.month -= 1
			d.date = 31 - date
		}
		//jika tidak di tahun kabisat dan tidak juga di februari
		// bulan ganjil{januari, maret, mei}
	} else if d.year%4 != 0 && d.month%2 != 0 {
		if date > 0 && date <= 31 {
			d.date = date
		} else if date > 31 {
			d.month += 1
			d.date = date - 31
		} else {
			d.month -= 1
			d.date = 30 - (date * -1)
		}
	} else {
		if date > 0 && date <= 30 {
			d.date = date
		} else if date > 30 {
			d.month += 1
			d.date = date - 30
		} else {
			d.month -= 1
			d.date = 31 - (date * -1)
		}
	}

	if day <= 7 && day > 0 {
		d.day = day
	} else if day > 7 {
		d.day = day - 7
	} else {
		d.day = 7 - day
	}
	return d
}

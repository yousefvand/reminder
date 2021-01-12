package main

// JalaliToGregorian :
func JalaliToGregorian(jy int, jm int, jd int) (int, int, int) {
	var gy, gm, gd, days int
	var salA = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	jy += 1595
	days = -355668 + (365 * jy) + ((jy / 33) * 8) + (((jy % 33) + 3) / 4) + jd
	if jm < 7 {
		days += (jm - 1) * 31
	} else {
		days += ((jm - 7) * 30) + 186
	}
	gy = 400 * (days / 146097)
	days %= 146097
	if days > 36524 {
		days--
		gy += 100 * (days / 36524)
		days %= 36524
		if days >= 365 {
			days++
		}
	}
	gy += 4 * (days / 1461)
	days %= 1461
	if days > 365 {
		gy += (days - 1) / 365
		days = (days - 1) % 365
	}
	gd = days + 1
	if (gy%4 == 0 && gy%100 != 0) || (gy%400 == 0) {
		salA[2] = 29
	}
	gm = 0
	for gm < 13 && gd > salA[gm] {
		gd -= salA[gm]
		gm++
	}
	return gy, gm, gd
}

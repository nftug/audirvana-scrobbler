package internal

import "time"

func float64ToTime(unixTime float64) time.Time {
	seconds := int64(unixTime)                                // 秒部分
	nanoseconds := int64((unixTime - float64(seconds)) * 1e9) // ナノ秒部分
	return time.Unix(seconds, nanoseconds).UTC()
}

func mjdToTime(mjd float64) time.Time {
	// MJDをユリウス日形式に変換
	julianDate := mjd + 2400000.5
	// ユリウス日からユニックスエポックに変換
	unixTime := (julianDate - 2440587.5) * 86400 // 秒換算
	return float64ToTime(unixTime)
}

func ParseSQLiteFloatToTime(value float64) time.Time {
	// 修正ユリウス日かユニックスエポックか判定
	if value > 1e10 { // ユニックスエポックの範囲
		return float64ToTime(value)
	}
	return mjdToTime(value)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	fmt.Printf("nowTime -> %s\n", nowTime)

	const format1 = "2006/01/02 15:04:05" // 24h表現、0埋めあり
	fmt.Printf("now -> %s\n", nowTime.Format(format1))

	const format2 = "2006/1/2 3:4:5" // 12h表現、0埋め無し
	fmt.Printf("now -> %s\n", nowTime.Format(format2))

	const DateFormat = "2006/01/02"
	const TimeFormat = "15:04:05"
	const TimeFormat2 = "15"
	const MilliFormat = "2006/01/02 15:04:05.000"
	const MicroFormat = "2006/01/02 15:04:05.000000"
	const NanoFormat = "2006/01/02 15:04:05.000000000"

	fmt.Printf("yyyy/MM/dd -> %s\n", nowTime.Format(DateFormat))
	fmt.Printf("HH:mm:ss   -> %s\n", nowTime.Format(TimeFormat))
	fmt.Printf("HH -> %s\n", nowTime.Format(TimeFormat2))

	// ミリ秒まで出力
	fmt.Printf("Milli -> %s\n", nowTime.Format(MilliFormat))

	// マイクロ秒まで出力
	fmt.Printf("Micro -> %s\n", nowTime.Format(MicroFormat))

	// ナノ秒まで出力
	fmt.Printf("Nano  -> %s\n", nowTime.Format(NanoFormat))

	// Unixtimeに変換
	unixTime := nowTime.Unix()
	fmt.Printf("unixTime -> %d\n", unixTime)
}

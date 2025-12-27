package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("🚀 年末カウントダウンを開始します（Ctrl+C で終了）")

	// 更新用のTickerを作成（1秒より少し短くして表示のラグを防ぐ）
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		// ターゲットは「翌年の1月1日 00:00:00」
		// 例: 今が2023年なら、2024年1月1日を目指す
		targetYear := now.Year() + 1
		target := time.Date(targetYear, 1, 1, 0, 0, 0, 0, now.Location())

		// 残り時間を計算
		diff := target.Sub(now)

		// もし年を越していたら（差分がマイナスなら）、おめでとうメッセージを表示
		if diff <= 0 {
			fmt.Printf("\r🎉 Happy New Year %d! 🎍                        \n", targetYear)
			break
		}

		// --- プログレスバーの計算（今年の1月1日から翌年1月1日まで） ---
		yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		totalYearDuration := target.Sub(yearStart)
		elapsed := now.Sub(yearStart)
		progress := float64(elapsed) / float64(totalYearDuration)

		barLength := 30
		completed := int(progress * float64(barLength))
		// 範囲外エラー防止
		if completed > barLength {
			completed = barLength
		}
		bar := strings.Repeat("■", completed) + strings.Repeat("□", barLength-completed)

		// --- 表示用の時間計算 ---
		// diff.Seconds() を int にキャストすることで秒未満を切り捨て
		totalSeconds := int(diff.Seconds())
		days := totalSeconds / (24 * 3600)
		hours := (totalSeconds % (24 * 3600)) / 3600
		minutes := (totalSeconds % 3600) / 60
		seconds := totalSeconds % 60

		// 表示（\rで行頭に戻る）
		// 以前の文字が残らないよう末尾に空白を入れています
		fmt.Printf("\r[%s] %3.3f%% | 残り: %d日 %02d時間 %02d分 %02d秒   ",
			bar, progress*100, days, hours, minutes, seconds)
	}
}

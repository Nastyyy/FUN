package main

import (
	"fmt"
	"time"
)

const epochs = 120

const div = "======================================================"

func main() {
	tickTimes := []tickData{}

	start := time.Now()
	m4 := Weapon{"M4A1", 100, 5}
	m16 := Weapon{"M16A3", 100, 3}
	player1 := &Player{"Player1", 1, 100, m4}
	player2 := &Player{"Player2", 2, 100, m16}
	game := NewGame(player1, player2)
	fmt.Println(div)
	for i := 0; i < epochs; i++ {
		tickStart := time.Now()

		game.tick(i)

		tickStop := time.Now()
		meta := tickData{tickStart, tickStop}
		fmt.Print(meta.GetDuration())

		fmt.Print("\n")
		tickTimes = append(tickTimes, meta)
		fmt.Println(div)
	}
	stop := time.Now()

	avgTick := getAvgTick(tickTimes)
	fmt.Printf("Total duration: %s | Avg tick: %dms\n", stop.Sub(start), avgTick)
}

func getAvgTick(ticks []tickData) int64 {
	var sum int64
	for _, tick := range ticks {
		time := tick.GetDuration().Milliseconds()
		sum += time
	}

	return sum / int64(len(ticks))
}

type tickData struct {
	start time.Time
	stop  time.Time
}

func (td tickData) GetDuration() time.Duration {
	return td.stop.Sub(td.start)
}

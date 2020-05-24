package musicplayer

import (
	"fmt"
	"time"
)

/**

 */
type MP4Player struct {
	stat     int
	progress int
}

func (p *MP4Player) Play(source string) {
	fmt.Println("Playing mp4 music", source)
	p.progress = 0
	for p.progress < 100 {
		// 假装正在播放
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinised playing", source)
}
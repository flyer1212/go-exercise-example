package musicplayer

import (
	"fmt"
)

/**

 */

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "MP4":
		p = &MP4Player{}
	default:
		fmt.Println("Unsupported music type:", mtype)
	}
	p.Play(source)
}

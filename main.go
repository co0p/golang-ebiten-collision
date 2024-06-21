package main

import (
	_ "embed"
	"flag"
	"fmt"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed crate.jpg
var crateImage []byte

//go:embed dummy.jpg
var dummyImage []byte

var ctFlag = flag.String("t", "", "define the type of collision detection, 'line' or 'aabb'")

func main() {

	flag.Parse()
	if *ctFlag != "line" && *ctFlag != "aabb" {
		fmt.Println("usage: -t <line|aabb>")
		return
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("collision")

	game := NewGame(*ctFlag)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

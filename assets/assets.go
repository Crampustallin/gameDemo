package assets

import (
	"embed"
	"io"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed all:fonts
	gameAssets embed.FS
	//go:embed audio/kick-hard.mp3
	Kick_hard []byte
)

func LoadFont() font.Face {
	r, err := gameAssets.Open("fonts/MPLUS1p-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := r.Close(); err != nil {
			log.Print("closing MPLUS1p-Regular: font reader: %v", err)
		}
	}()
	fontData, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions {
		Size: 24,
		DPI: 72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return fontFace
}

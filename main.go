package main

import (
	"os"
	"log"

	_ "image/png"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/tinne26/etxt"
	"github.com/yuin/gopher-lua"
	"golang.org/x/image/font/sfnt"

	"bufio"
)

var (
	teto *ebiten.Image
	intr string = "TETO WORD OF THE DAY"
	word string = ""
)

type Game struct{
	funnyFont	*etxt.Renderer
}

func init() {
	if teto == nil {
		var err error
		teto, _, err = ebitenutil.NewImageFromFile("data/teto.png")
		if err != nil {
			log.Fatal(err)
		}
	}

	tiktok, err := os.Open("data/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer tiktok.Close()
	
	var baguettes []string
	mesmerize := bufio.NewScanner(tiktok)
	for mesmerize.Scan() {
		baguettes = append(baguettes, mesmerize.Text())
	}

	if err := mesmerize.Err(); err != nil {
		log.Fatal(err)
	}

	if len(baguettes) == 0 {
		word = "CANCELLED"
	}

	L := lua.NewState()
	defer L.Close()

	argTable := L.CreateTable(1, 0)
	if len(os.Args) == 2 {
		L.RawSetInt(argTable, 1, lua.LString(string(os.Args[1])))
	} else if len(os.Args) > 2 {
		L.RawSetInt(argTable, 1, lua.LString("OVERFLOW!!!"))
	} else {
		L.RawSetInt(argTable, 1, SliceToTable(L, baguettes))
	}
	L.SetGlobal("arg", argTable)
	
	if err := L.DoFile("lib/randomText.lua"); err != nil {
		panic(err)
	}
	ret := L.Get(-1)

	if str, ok := ret.(lua.LString); ok {
        word = string(str)
	}
	L.Pop(1)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	sw, sh := screen.Size()
	tetoSize := teto.Bounds().Size()
	op := &ebiten.DrawImageOptions{}

	screen.Fill(color.White)

	x := float64(sw-tetoSize.X) / 2
    y := float64(sh-tetoSize.Y) / 2

    op.GeoM.Translate(x,y)
    screen.DrawImage(teto, op)

	g.funnyFont.SetSize(64)
	g.funnyFont.SetAlign(etxt.Center)
	g.funnyFont.SetColor(color.Black)

	offsets := []struct{ dx, dy int }{
		{-3, -3}, {-3, 3}, {3, -3}, {3, 3},
		{-3, 0}, {3, 0}, {0, -3}, {0, 3},
	}

	topX, topY := sw/2, 40
	bottomX, bottomY := sw/2, sh-40
	for _, offset := range offsets {
		g.funnyFont.Draw(screen, intr, topX+offset.dx, topY+offset.dy)
		g.funnyFont.Draw(screen, word, bottomX+offset.dx, bottomY+offset.dy)
	}

	g.funnyFont.SetColor(color.White)
	g.funnyFont.Draw(screen, intr, topX, topY)
	g.funnyFont.Draw(screen, word, bottomX, bottomY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	ebiten.SetWindowSize(640,640)
	ebiten.SetWindowTitle("Teto Word of the Day!")
    
    fontBytes, err := os.ReadFile("data/impact.ttf")
    if err != nil {
        log.Fatal(err)
    }
    impactful, err := sfnt.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	renderer := etxt.NewRenderer()
	renderer.SetFont(impactful)
	renderer.Utils().SetCache8MiB()

	dailyWord := &Game{
		funnyFont: renderer,
	}

	if err := ebiten.RunGame(dailyWord); err != nil {
		log.Fatal(err)
	}
}

func SliceToTable(L *lua.LState, slice []string) *lua.LTable {
	tbl := L.NewTable()

	for _, item := range slice {
		tbl.Append(lua.LString(item))
	}

	return tbl
}
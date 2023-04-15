package menu

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const ()

func GetScreenCenter(screen *ebiten.Image) (int, int) {
	sw, sh := screen.Size()
	return sw / 2, sh / 2
}

func DrawAnt(screen *ebiten.Image) {
	//800x 600y
	//Center -400 Center-300
	//[0][0] = [240][160]
	ebitenutil.DrawRect(screen, 240+10, 160+10, 1, 1, color.RGBA{255, 255, 255, 255})
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	centerX, centerY := GetScreenCenter(screen)
	screen.Fill(color.RGBA{0x3d, 0x3d, 0x42, 0xff})

	//Draw Playground
	ebitenutil.DrawRect(screen, float64(centerX)-400, float64(centerY)-300, 800, 600, color.Black)
	//Name
	ebitenutil.DebugPrintAt(screen, "Ant Playground", centerX-40, centerY-330)
	//Legend
	ebitenutil.DebugPrintAt(screen, "Legend:\n   Ant\n   Food\n   Anthill", centerX-400, centerY+320)
	ebitenutil.DrawRect(screen, float64(centerX)-395, float64(centerY)+340, 8, 8, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawRect(screen, float64(centerX)-395, float64(centerY)+356, 8, 8, color.RGBA{0, 255, 0, 255})
	ebitenutil.DrawRect(screen, float64(centerX)-395, float64(centerY)+372, 8, 8, color.RGBA{255, 0, 255, 255})
	//FPS
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()), centerX+330, centerY-330)
	DrawAnt(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 920
}

func StartIntro() {
	fmt.Println("Created by Mateusz Redzimski")
}

func InitializeGame() {
	ebiten.SetWindowSize(1280, 920)
	ebiten.SetWindowTitle("Ants Simulator")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

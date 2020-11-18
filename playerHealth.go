package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)
const healthSize = 105
var pHealth = 5
type playerHealth struct {
	
	container *element
	tex       *sdl.Texture

	width, height float64

}


func newPlayerHealth(container *element, renderer *sdl.Renderer, filename string) *playerHealth {
	tex2 := hTextureFromBMP(renderer, filename)

	_, _, width, height, err := tex2.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &playerHealth{
		container: container,
		tex:       hTextureFromBMP(renderer, filename),
		width:     float64(width),
		height:    float64(height),
	}
}

func (hp *playerHealth) onDraw(renderer *sdl.Renderer) error {
	// Converting coordinates to top left of sprite
	

	if pHealth == 5 {
		renderer.CopyEx(
			hp.tex,
			&sdl.Rect{X: 105, Y: 105, W: 210, H: 210},
			&sdl.Rect{X: 0, Y: 650, W: 105, H: 105},
			hp.container.rotation,
			&sdl.Point{X: 105, Y: 105},
			sdl.FLIP_NONE)
	}
		
	if pHealth == 4 {
		renderer.CopyEx(
		hp.tex,
		&sdl.Rect{X: 70, Y: 105, W: 170, H: 205},
		&sdl.Rect{X: 0, Y: 650, W: 105, H: 105},
		hp.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
		
	if pHealth == 3 {
		renderer.CopyEx(
		hp.tex,
		&sdl.Rect{X: 40, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 0, Y: 650, W: 105, H: 105},
		hp.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
	
	if pHealth == 2 {
		renderer.CopyEx(
		hp.tex,
		&sdl.Rect{X: 30, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 0, Y: 650, W: 105, H: 105},
		hp.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
	if pHealth == 1 {
		renderer.CopyEx(
		hp.tex,
		&sdl.Rect{X: 15, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 0, Y: 650, W: 105, H: 105},
		hp.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
	if pHealth == 0 {
		renderer.CopyEx(
		hp.tex,
		&sdl.Rect{X: 0, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 0, Y: 650, W: 105, H: 105},
		hp.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
	
	
	


	return nil
}


func (hp *playerHealth) onUpdate() error {

	
	return nil
}

func hTextureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}

func (hp *playerHealth) onCollision(other *element) error {

	if other.tag == "enemyBullet" {
		pHealth = pHealth - 1
		print("Player Health lost")
		
	}	

	return nil
}




type vulnerablePlayer struct {
	container *element

}

func newVulnerablePlayer (container *element) *vulnerablePlayer {
	return &vulnerablePlayer{container: container}
}

func(vp *vulnerablePlayer) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func(vp *vulnerablePlayer) onUpdate() error {
	return nil
}

func(vp *vulnerablePlayer) onCollision(other *element) error {
	
	
	if pHealth == 0 {
		vp.container.active = false
	}
	
	
	return nil	
	
}



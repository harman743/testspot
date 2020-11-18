package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)
const enemyHealthSize = 105
var eHealth = 5
type enemyHealth struct {
	
	container *element
	tex       *sdl.Texture

	width, height float64

}


func newEnemyHealth(container *element, renderer *sdl.Renderer, filename string) *enemyHealth {
	tex := hETextureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &enemyHealth{
		container: container,
		tex:       hETextureFromBMP(renderer, filename),
		width:     float64(width),
		height:    float64(height),
	}
}

func (he *enemyHealth) onDraw(renderer *sdl.Renderer) error {
	// Converting coordinates to top left of sprite
	
	if eHealth == 5 {
		renderer.CopyEx(
			he.tex,
			&sdl.Rect{X: 105, Y: 105, W: 210, H: 210},
			&sdl.Rect{X: 1180, Y: -134, W: 105, H: 105},
			he.container.rotation,
			&sdl.Point{X: 105, Y: 105},
			sdl.FLIP_NONE)
	}
		
	if eHealth == 4 {
		renderer.CopyEx(
		he.tex,
		&sdl.Rect{X: 135, Y: 105, W: 170, H: 205},
		&sdl.Rect{X: 1180, Y: -134, W: 105, H: 105},
		he.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
		
	if eHealth == 3 {
		renderer.CopyEx(
		he.tex,
		&sdl.Rect{X: 175, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 1180, Y: -134, W: 105, H: 105},
		he.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
	
	if eHealth == 2 {
		renderer.CopyEx(
		he.tex,
		&sdl.Rect{X: 200, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 1180, Y: -134, W: 105, H: 105},
		he.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
	if eHealth == 1 {
		renderer.CopyEx(
		he.tex,
		&sdl.Rect{X: 210, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 1180, Y: -134, W: 105, H: 105},
		he.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	
	if eHealth == 0 {
		renderer.CopyEx(
		he.tex,
		&sdl.Rect{X: 0, Y: 105, W: 140, H: 205},
		&sdl.Rect{X: 1180, Y: -134, W: 105, H: 105},
		he.container.rotation,
		&sdl.Point{X: 105, Y: 105},
		sdl.FLIP_NONE)
	
	}	

	return nil
}


func (he *enemyHealth) onUpdate() error {

	
	return nil
}

func hETextureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
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

func (he *enemyHealth) onCollision(other *element) error {

	if other.tag == "playerBullet" {
		eHealth = eHealth - 1
		print("Enemy Health lost")
		
	}	

	return nil
}




type vulnerableEnemy struct {
	container *element

}

func newVulnerableEnemy (container *element) *vulnerableEnemy {
	return &vulnerableEnemy{container: container}
}

func(ve *vulnerableEnemy) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func(ve *vulnerableEnemy) onUpdate() error {
	return nil
}

func(ve *vulnerableEnemy) onCollision(other *element) error {
	
	
	if eHealth == 0 {
		ve.container.active = false
	}
	
	
	return nil	
	
}



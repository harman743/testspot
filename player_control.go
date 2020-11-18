package main

import (
	"math"
	"time"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover2 struct {
	container *element
	speed     float64

	sr *spriteRenderer
}

func newkeyboardMover2(container *element, speed float64) *keyboardMover2 {
	
	return &keyboardMover2{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover2 *keyboardMover2) onCollision(other *element) error {

	return nil
}

func (mover2 *keyboardMover2) onUpdate() error {
	
	keys := sdl.GetKeyboardState()
	
	cont := mover2.container
	
	col := circle{
		center: cont.position,
		radius: 80,
		
	}
	cont.collisions = append(cont.collisions, col)

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if cont.position.x-(mover2.sr.width/2.0) > 0 {
			cont.position.x -= mover2.speed
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if cont.position.x+(mover2.sr.height/2.0) < screenWidth {
			cont.position.x += mover2.speed
		}
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		if cont.position.y+(mover2.sr.height/2.0) < screenWidth {
			cont.position.y -= mover2.speed
		}
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		if cont.position.y+(mover2.sr.height/2.0) < screenWidth {
			cont.position.y += mover2.speed
		}
	}

	return nil
}

func (mover2 *keyboardMover2) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter2 struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
	tag	  string
}

func newkeyboardShooter2(container *element, cooldown time.Duration) *keyboardShooter2 {
	return &keyboardShooter2{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover2 *keyboardShooter2) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover2.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover2.lastShot) >= mover2.cooldown {
			mover2.shoot(pos.x-40, pos.y-110)
			

			mover2.lastShot = time.Now()
		}
	}

	return nil
}

func (mover2 *keyboardShooter2) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover2 *keyboardShooter2) shoot(x, y float64) {
	if bul, ok := playerBulletFromPool(); ok {
	
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.rotation = 270 * (math.Pi / 180)
	}
}

func (mover2 *keyboardShooter2) onCollision(other *element) error {
	
	return nil
}

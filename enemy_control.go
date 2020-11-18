package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64

	sr *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	
	cont := mover.container
	
	col := circle{
		center: cont.position,
		radius: 90,
		
	}
	
	cont.collisions = append(cont.collisions, col)
	
	if keys[sdl.SCANCODE_A] == 1 {
		if cont.position.x-(mover.sr.width/2.0) > 0 {
			cont.position.x -= mover.speed
		}
	}
	if keys[sdl.SCANCODE_D] == 1 {
		if cont.position.x+(mover.sr.height/2.0) < screenWidth {
			cont.position.x += mover.speed
		}
	}
	if keys[sdl.SCANCODE_W] == 1 {
		if cont.position.y+(mover.sr.height/2.0) < screenWidth {
			cont.position.y -= mover.speed
		}
	}
	if keys[sdl.SCANCODE_S] == 1 {
		if cont.position.y+(mover.sr.height/2.0) < screenWidth {
			cont.position.y += mover.speed
		}
	}


	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
	tag       string
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.position

	if keys[sdl.SCANCODE_R] == 1 {
		if time.Since(mover.lastShot) >= mover.cooldown {
			mover.shoot(pos.x, pos.y+100)

			mover.lastShot = time.Now()
		}
	}

	return nil
}

func (mover *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardShooter) shoot(x, y float64) {
	if bul2, ok := enemyBulletFromPool(); ok {
		
		bul2.active = true
		bul2.position.x = x
		bul2.position.y = y
		bul2.rotation = 90 * (math.Pi / 180)
	}
}

func (mover *keyboardShooter) onCollision(other *element) error {
	return nil
}

package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 9
	playerSize  = 105

	playerShotCooldown = time.Millisecond * 250
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - playerSize/2.0,
	}

	player.active = true

	sr := newSpriteRenderer(player, renderer, "Sprites/player.bmp")
	player.addComponent(sr)
	
	hp := newPlayerHealth(player, renderer, "Sprites/playerHealth.bmp")
	player.addComponent(hp)
	
	vp := newVulnerablePlayer(player)
	player.addComponent(vp)


	

	mover2 := newkeyboardMover2(player, playerSpeed)
	player.addComponent(mover2)

	shooter := newkeyboardShooter2(player, playerShotCooldown)
	player.addComponent(shooter)

	return player
}



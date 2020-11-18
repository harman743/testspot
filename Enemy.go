package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const basicEnemySize = 105
const enemySpeed = 9

const enemyShotCooldown = time.Millisecond * 250

func newBasicEnemy(renderer *sdl.Renderer) *element {
	basicEnemy := &element{}

	basicEnemy.position = vector{
		x: screenWidth /1.5,
		y: screenHeight - basicEnemySize*6.0,
	}
	
	
	basicEnemy.rotation = 180

	basicEnemy.active = true

	sr := newSpriteRenderer(basicEnemy, renderer, "Sprites/Enemy.bmp")
	basicEnemy.addComponent(sr)
	
	he := newEnemyHealth(basicEnemy, renderer, "Sprites/enemyHealth.bmp")
	basicEnemy.addComponent(he)
	
	ve := newVulnerableEnemy(basicEnemy)
	basicEnemy.addComponent(ve)

	
	
	

	mover := newKeyboardMover(basicEnemy, enemySpeed)
	basicEnemy.addComponent(mover)

	shooter := newKeyboardShooter(basicEnemy, playerShotCooldown)
	basicEnemy.addComponent(shooter)

	return basicEnemy
}



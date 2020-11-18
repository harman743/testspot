package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 10
)

func newEnemyBullet(renderer *sdl.Renderer) *element {
	bullet2 := &element{}

	sr := newSpriteRenderer(bullet2, renderer, "Sprites/pbullet.bmp")
	bullet2.addComponent(sr)

	mover := newBulletMover(bullet2, bulletSpeed)
	bullet2.addComponent(mover)

	col := circle{
		center: bullet2.position,
		radius: 4,
	}
	bullet2.collisions = append(bullet2.collisions, col)
	
	bullet2.tag = "enemyBullet"
	

	return bullet2
}

func newPlayerBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "Sprites/pbullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	col := circle{
		center: bullet.position,
		radius: 8,
	}
	bullet.collisions = append(bullet.collisions, col)
	
	bullet.tag = "playerBullet"
	
	
	return bullet
}


var playerBulletPool []*element

var enemyBulletPool []*element

func initEnemyBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul2 := newEnemyBullet(renderer)
		elements = append(elements, bul2)
		enemyBulletPool = append(enemyBulletPool, bul2)
	}
}

func initPlayerBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newPlayerBullet(renderer)
		elements = append(elements, bul)
		playerBulletPool = append(playerBulletPool, bul)
	}
}


func enemyBulletFromPool() (*element, bool) {
	for _, bul2 := range enemyBulletPool {
		if !bul2.active {
			return bul2, true
		}
	}

	return nil, false
}

func playerBulletFromPool() (*element, bool) {
	for _, bul := range playerBulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}



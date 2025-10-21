package main

import "fmt"

type Player struct {
	username, password string
	game_count         int8
}

func (player Player) struct_method() string {
	return fmt.Sprintf("halo %s, anda sudah bermain sebanyak %dx", player.username, player.game_count)
}

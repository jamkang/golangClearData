package heartPackage

func (self *Server) GetPlayer(uid int) *Player {
	player, ok := self.players[uid]
	if !ok {
		return nil
	}

	return player
}

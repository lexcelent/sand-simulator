package base

// Material should be able to update
type Material interface {
	Update(g *Game)
	NameID() int
}

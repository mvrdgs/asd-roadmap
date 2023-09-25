package colors

type blue struct{}

func Blue() Color {
	return &blue{}
}

func (b *blue) Fill() string {
	return "blue"
}

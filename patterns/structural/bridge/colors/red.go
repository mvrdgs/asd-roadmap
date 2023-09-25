package colors

type red struct{}

func Red() Color {
	return &red{}
}

func (r *red) Fill() string {
	return "red"
}

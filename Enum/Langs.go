package enum


type Lang int

const (
	JavaScript Lang = iota
	GoLang 
	Java
)

func (l Lang) String() string {
	return [...]string{"JavaScript", "GoLang", "Java"}[l]
}

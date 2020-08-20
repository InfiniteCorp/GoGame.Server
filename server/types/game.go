package types

// Interface similar to every game
type IGame interface {
	Init()

	Play([]byte, uint8) error

	GetState() interface{}
}

type GameDefinition struct {
	ID              string             `yaml:"id"`
	LibPath         string             `yaml:"lib_path"`
	CreatorFunction func() interface{} `yaml:",omitempty"`
}

type Game struct {
	Name string `yaml:"name"`
	Game IGame  `yaml:",omitempty"`
}

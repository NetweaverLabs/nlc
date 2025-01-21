package request

type Request struct {
	Cmd  string `gob:"Cmd"`
	Args any    `gob:"Args"`
}

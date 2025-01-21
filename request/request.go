package request

type Request struct {
	Cmd  string   `gob:"Cmd"`
	Args []string `gob:"Args"`
}

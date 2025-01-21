package response

type Response struct {
	Status  string   `gob:"Status"`
	Payload []string `gob:"Payload"`
}

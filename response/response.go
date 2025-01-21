package response

type Response struct {
	Status  string `gob:"Status"`
	Payload any    `gob:"Payload"`
}

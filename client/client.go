package client

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/NetweaverLab/nlc/request"
	"github.com/NetweaverLab/nlc/response"
)

type DaemonClient struct {
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func NewDaemonClient() (*DaemonClient, error) {
	conn, err := net.Dial("unix", socketpath)
	if err != nil {
		return nil, err
	}
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)
	return &DaemonClient{
		encoder: enc,
		decoder: dec,
	}, nil
}

func (dc *DaemonClient) Send(req *request.Request) error {
	if err := dc.encoder.Encode(req); err != nil {
		return fmt.Errorf("daemon client couldnot send the request: %s", err.Error())
	}
	return nil
}

func (dc *DaemonClient) Recieve(resp *response.Response) error {
	if err := dc.decoder.Decode(resp); err != nil {
		return fmt.Errorf("deamon client failed to recieve the response: %s", err.Error())
	}
	return nil
}

package harmony

import (
	"context"
	"encoding/json"

	"github.com/skwair/harmony/internal/payload"
)

// sendPayload sends a single Payload to the Gateway with
// the given op and data.
func (c *Client) sendPayload(ctx context.Context, op int, d interface{}) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	p := &payload.Payload{Op: op, D: b}
	c.logger.Debugf("sent payload: %s", p)
	return payload.Send(ctx, c.conn, p)
}

// recvPayload receives a single Payload from the Gateway.
func (c *Client) recvPayload() (*payload.Payload, error) {
	p, err := payload.Recv(c.ctx, &c.connRMu, c.conn)
	if err != nil {
		return nil, err
	}

	c.logger.Debugf("received payload: %s", p)

	return p, nil
}

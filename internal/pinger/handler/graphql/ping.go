package pinger_resolver

import "fmt"

func (p *PingResolver) Message() string {
	return fmt.Sprintf("Pong! from %s", p.message)
}

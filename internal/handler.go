package internal

import (
	"context"

	"github.com/jkstack/anet"
)

func (agent *Agent) OnConnect() {
}

func (agent *Agent) OnDisconnect() {
}

func (agent *Agent) OnReportMonitor() {
}

func (agent *Agent) OnMessage(msg *anet.Msg) error {
	switch msg.Type {
	case anet.TypeHMStaticReq:
		var rep anet.Msg
		rep.Type = anet.TypeHMStaticRep
		rep.TaskID = msg.TaskID
		rep.HMStatic = getStatic()
		agent.chWrite <- &rep
	}
	return nil
}

func (agent *Agent) LoopWrite(ctx context.Context, ch chan *anet.Msg) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-agent.chWrite:
			ch <- msg
		}
	}
}

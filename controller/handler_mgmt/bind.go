/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package handler_mgmt

import (
	"github.com/openziti/channel/v2"
	"github.com/openziti/ziti/common/trace"
	"github.com/openziti/ziti/controller/network"
	"github.com/openziti/ziti/controller/xmgmt"
)

type BindHandler struct {
	network          *network.Network
	xmgmts           []xmgmt.Xmgmt
	sshTunnelManager *sshTunnelManager
}

func NewBindHandler(network *network.Network, xmgmts []xmgmt.Xmgmt) channel.BindHandler {
	return &BindHandler{
		network:          network,
		xmgmts:           xmgmts,
		sshTunnelManager: &sshTunnelManager{},
	}
}

func (bindHandler *BindHandler) BindChannel(binding channel.Binding) error {
	binding.AddTypedReceiveHandler(newInspectHandler(bindHandler.network))

	tracesHandler := newStreamTracesHandler(bindHandler.network)
	binding.AddTypedReceiveHandler(tracesHandler)
	binding.AddCloseHandler(tracesHandler)

	eventsHandler := newStreamEventsHandler(bindHandler.network)
	binding.AddTypedReceiveHandler(eventsHandler)
	binding.AddCloseHandler(eventsHandler)

	binding.AddTypedReceiveHandler(newTogglePipeTracesHandler(bindHandler.network))

	binding.AddPeekHandler(trace.NewChannelPeekHandler(bindHandler.network.GetAppId(), binding.GetChannel(), bindHandler.network.GetTraceController()))

	sshTunnelRequestHandler := newSshTunnelHandler(bindHandler.network, bindHandler.sshTunnelManager)
	binding.AddTypedReceiveHandler(&channel.AsyncFunctionReceiveAdapter{
		Type:    sshTunnelRequestHandler.ContentType(),
		Handler: sshTunnelRequestHandler.HandleReceive,
	})
	binding.AddCloseHandler(sshTunnelRequestHandler)
	binding.AddTypedReceiveHandler(newSshTunnelDataHandler(bindHandler.sshTunnelManager))

	xmgmtDone := make(chan struct{})
	for _, x := range bindHandler.xmgmts {
		if err := binding.Bind(x); err != nil {
			return err
		}
		if err := x.Run(binding.GetChannel(), xmgmtDone); err != nil {
			return err
		}
	}
	if len(bindHandler.xmgmts) > 0 {
		binding.AddCloseHandler(newXmgmtCloseHandler(xmgmtDone))
	}

	return nil
}

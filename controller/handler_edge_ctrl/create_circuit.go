/*
	Copyright NetFoundry, Inc.

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

package handler_edge_ctrl

import (
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/edge/pb/edge_ctrl_pb"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/sdk-golang/ziti/edge"
	"github.com/pkg/errors"
)

type createCircuitHandler struct {
	baseRequestHandler
}

func NewCreateCircuitHandler(appEnv *env.AppEnv, ch channel2.Channel) channel2.ReceiveHandler {
	return &createCircuitHandler{
		baseRequestHandler: baseRequestHandler{
			ch:     ch,
			appEnv: appEnv,
		},
	}
}

func (self *createCircuitHandler) ContentType() int32 {
	return int32(edge_ctrl_pb.ContentType_CreateCircuitRequestType)
}

func (self *createCircuitHandler) Label() string {
	return "create.circuit"
}

func (self *createCircuitHandler) sendResponse(ctx *CreateCircuitRequestContext, response *edge_ctrl_pb.CreateCircuitResponse) {
	log := pfxlog.ContextLogger(self.ch.Label())

	body, err := proto.Marshal(response)
	if err != nil {
		log.WithError(err).WithField("token", ctx.req.SessionToken).Error("failed to marshal create circuit response")
		return
	}

	responseMsg := channel2.NewMessage(response.GetContentType(), body)
	responseMsg.ReplyTo(ctx.msg)
	if err = self.ch.Send(responseMsg); err != nil {
		log.WithError(err).WithField("token", ctx.req.SessionToken).Error("failed to send create circuit response")
	}
}

func (self *createCircuitHandler) HandleReceive(msg *channel2.Message, ch channel2.Channel) {
	req := &edge_ctrl_pb.CreateCircuitRequest{}
	if err := proto.Unmarshal(msg.Body, req); err != nil {
		pfxlog.ContextLogger(ch.Label()).WithError(err).Error("could not unmarshal CreateCircuitRequest")
		return
	}

	ctx := &CreateCircuitRequestContext{
		baseRequestContext: baseRequestContext{handler: self, msg: msg},
		req:                req,
	}

	go self.CreateCircuit(ctx)
}

func (self *createCircuitHandler) CreateCircuit(ctx *CreateCircuitRequestContext) {
	if !ctx.loadRouter() {
		return
	}
	ctx.loadSession(ctx.req.SessionToken)
	ctx.checkSessionType(persistence.SessionTypeDial)
	ctx.checkSessionFingerprints(ctx.req.Fingerprints)
	ctx.verifyEdgeRouterAccess()
	ctx.loadService()

	if ctx.err != nil {
		self.returnError(ctx, ctx.err)
		return
	}

	if ctx.service.EncryptionRequired && ctx.req.PeerData[edge.PublicKeyHeader] == nil {
		self.returnError(ctx, errors.New("encryption required on service, initiator did not send public header"))
		return
	}

	serviceId := ctx.session.ServiceId
	if ctx.req.TerminatorIdentity != "" {
		serviceId = ctx.req.TerminatorIdentity + "@" + serviceId
	}

	clientId := &identity.TokenId{Token: ctx.session.Id, Data: ctx.req.PeerData}

	n := self.appEnv.GetHostController().GetNetwork()
	sessionInfo, err := n.CreateSession(ctx.sourceRouter, clientId, serviceId)
	if err != nil {
		self.returnError(ctx, err)
		return
	}

	log := pfxlog.ContextLogger(self.ch.Label()).WithField("token", ctx.req.SessionToken)

	response := &edge_ctrl_pb.CreateCircuitResponse{
		SessionId: sessionInfo.Id.Token,
		Address:   sessionInfo.Circuit.IngressId,
		PeerData:  map[uint32][]byte{},
	}

	//static terminator peer data
	for k, v := range sessionInfo.Terminator.GetPeerData() {
		response.PeerData[k] = v
	}

	//runtime peer data
	for k, v := range sessionInfo.PeerData {
		response.PeerData[k] = v
	}

	if ctx.service.EncryptionRequired && response.PeerData[edge.PublicKeyHeader] == nil {
		self.returnError(ctx, errors.New("encryption required on service, terminator did not send public header"))
		if err := n.RemoveSession(sessionInfo.Id, true); err != nil {
			log.WithError(err).Error("failed to remove session")
		}
		return
	}

	log.Debugf("responding with successful circuit setup")
	self.sendResponse(ctx, response)
}

type CreateCircuitRequestContext struct {
	baseRequestContext
	req *edge_ctrl_pb.CreateCircuitRequest
}

func (self *CreateCircuitRequestContext) GetSessionToken() string {
	return self.req.SessionToken
}

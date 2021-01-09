package sfu

import (
	"github.com/Derek-X-Wang/relais/pkg/log"
	"github.com/Derek-X-Wang/relais/pkg/proto"
	"github.com/Derek-X-Wang/relais/pkg/rtc"
	"github.com/Derek-X-Wang/relais/pkg/util"
	nprotoo "github.com/cloudwebrtc/nats-protoo"
)

var (
	//nolint:unused
	dc = "default"
	//nolint:unused
	nid         = "sfu-unkown-node-id"
	protoo      *nprotoo.NatsProtoo
	broadcaster *nprotoo.Broadcaster
)

// Init func
func Init(dcID, nodeID, rpcID, eventID, natsURL string) {
	dc = dcID
	nid = nodeID
	protoo = nprotoo.NewNatsProtoo(natsURL)
	broadcaster = protoo.NewBroadcaster(eventID)
	handleRequest(rpcID)
	checkRTC()
}

// checkRTC send `stream-remove` msg to islb when some pub has been cleaned
func checkRTC() {
	log.Infof("SFU.checkRTC start")
	go func() {
		for mid := range rtc.CleanChannel {
			broadcaster.Say(proto.SFUStreamRemove, util.Map("mid", mid))
		}
	}()
}

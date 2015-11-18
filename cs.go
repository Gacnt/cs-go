package csgo

import (
	"log"

	"github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/internal"
	"github.com/Philipp15b/go-steam/internal/gamecoordinator"
	proto "github.com/golang/protobuf/proto"
)

const AppId = 730

type CS struct {
	client *steam.Client
}

func NewCSGO(client *steam.Client) *CS {
	c := &CS{client}
	client.GC.RegisterPacketHandler(c)
	return c
}

func (c *CS) SetPlaying(playing bool) {
	if playing {
		c.client.GC.SetGamesPlayed(730)
	} else {
		c.client.GC.SetGamesPlayed()
	}
}

func (c *CS) GetPlayerProfile(accountid uint64) {
	newAccId := accountid - 76561197960265728

	c.client.Write(internal.NewClientMsgProtobuf(k_EMsgGCCstrike15_v2_ClientRequestPlayersProfile, &CMsgGCCStrike15_v2_ClientRequestPlayersProfile{
		AccountId: proto.Uint32(uint32(newAccId)),
	}))
}

func (c *CS) ShakeHands() {
}

type GCReadyEvent struct{}

func (c *CS) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppId {
		return
	}

	switch EGCBaseClientMsg(packet.MsgType) {
	case EGCBaseClientMsg_k_EmsgGCClientWelcome:
		log.Println(packet)
		c.client.Emit(&GCReadyEvent{})
	}
	log.Println(packet)
}

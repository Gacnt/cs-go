package csgo

import (
	"log"
	"time"

	"github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/internal/gamecoordinator"
	proto "github.com/golang/protobuf/proto"
)

const AppId = 730

type CS struct {
	client      *steam.Client
	isConnected bool
}

func NewCSGO(client *steam.Client) *CS {
	c := &CS{client, false}
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

	c.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(k_EMsgGCCstrike15_v2_ClientRequestPlayersProfile), &CMsgGCCStrike15_v2_ClientRequestPlayersProfile{
		AccountId:    proto.Uint32(uint32(newAccId)),
		RequestLevel: proto.Uint32(32),
	}))
}

func (c *CS) ShakeHands() {
	var _ = time.Second
	c.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(EGCBaseClientMsg_k_EmsgGCClientHello), &CMsgClientHello{
		Version: proto.Uint32(1),
	}))

	time.Sleep(10 * time.Second)
	if !c.isConnected {
		log.Println("Gamecoordinator didn't shake our hand, gonna try talking again.")
		c.ShakeHands()
	} else {
		return
	}
}

type GCReadyEvent struct{}

func (c *CS) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppId {
		return
	}

	log.Println(packet)
	switch EGCBaseClientMsg(packet.MsgType) {
	case EGCBaseClientMsg_k_EmsgGCClientWelcome:
		c.isConnected = true
		c.client.Emit(&GCReadyEvent{})
		return
	case EGCBaseClientMsg(k_EMsgGCCstrike15_v2_PlayersProfile):
	}
}

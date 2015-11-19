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
	// Try to avoid not being ready on instant call of connection
	time.Sleep(5 * time.Second)
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

func (c *CS) RankString(rank uint32) (rankString string) {
	switch rank {
	case 1:
		rankString = "Silver 1"
	case 2:
		rankString = "Silver 2"
	case 3:
		rankString = "Silver 3"
	case 4:
		rankString = "Silver 4"
	case 5:
		rankString = "Silver Elite"
	case 6:
		rankString = "Silver Elite Master"
	case 7:
		rankString = "Gold Nova 1"
	case 8:
		rankString = "Gold Nova 2"
	case 9:
		rankString = "Gold Nova 3"
	case 10:
		rankString = "Gold Nova Master"
	case 11:
		rankString = "Master Guardian 1"
	case 12:
		rankString = "Master Guardian 2"
	case 13:
		rankString = "Master Guardian Elite"
	case 14:
		rankString = "Distinguished Master Guardian"
	case 15:
		rankString = "Legendary Eagle"
	case 16:
		rankString = "Legendary Eagle Master"
	case 17:
		rankString = "Supreme Master First Class"
	case 18:
		rankString = "Global Elite"
	}
	return rankString
}

type GCReadyEvent struct{}

type GCProfileFoundEvent struct {
	Profile CMsgGCCStrike15_v2_MatchmakingGC2ClientHello
}

type GCProfileNotFoundEvent struct{}

func (c *CS) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppId {
		return
	}

	switch EGCBaseClientMsg(packet.MsgType) {
	case EGCBaseClientMsg_k_EmsgGCClientWelcome:
		if c.isConnected {
			break
		}
		c.isConnected = true
		c.client.Emit(&GCReadyEvent{})
		return
	case EGCBaseClientMsg(k_EMsgGCCstrike15_v2_PlayersProfile):
		profile := new(CMsgGCCStrike15_v2_PlayersProfile)
		packet.ReadProtoMsg(profile)
		if profile.AccountProfiles == nil {
			c.client.Emit(&GCProfileNotFoundEvent{})
			break
		} else {
			c.client.Emit(&GCProfileFoundEvent{Profile: *profile.AccountProfiles[0]})

		}
	default:
		log.Println(packet)
	}
}

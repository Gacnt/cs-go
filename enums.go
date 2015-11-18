package csgo

import (
	"github.com/Philipp15b/go-steam/internal/steamlang"
)

type EGCBaseClientMsg int32

const (
	EGCBaseClientMsg_k_EmsgGCClientWelcome          EGCBaseClientMsg = 4004
	EGCBaseClientMsg_k_EmsgGCServerWelcome          EGCBaseClientMsg = 4005
	EGCBaseClientMsg_k_EmsgGCClientHello            EGCBaseClientMsg = 4006
	EGCBaseClientMsg_k_EmsgGCServerHello            EGCBaseClientMsg = 4007
	EGCBaseClientMsg_k_EmsgGCClientConnectionStatus EGCBaseClientMsg = 4008
	EGCBaseClientMsg_k_EmsgGCServerConnectionStatus EGCBaseClientMsg = 4009
)

const (
	k_EMsgGCCstrike15_v2_ClientRequestPlayersProfile steamlang.EMsg = 9127
	k_EMsgGCCstrike15_v2_PlayersProfile              steamlang.EMsg = 9128
)

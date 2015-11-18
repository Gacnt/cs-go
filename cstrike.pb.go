package csgo

import (
	proto "github.com/golang/protobuf/proto"
)

type CMsgGCCStrike15_v2_ClientRequestPlayersProfile struct {
	RequestIdDeprecated  *uint32 `protobuf:"varint,1,opt,name=request_id_deprecated" json:"request_id_deprecated,omitempty"`
	RequestIdsDeprecated *uint32 `protobuf:"varint,2,opt,name=request_ids_deprecated" json:"request_ids_deprecated,omitempty"`
	AccountId            *uint32 `protobuf:"varint,3,opt,name=account_id" json:"account_id,omitempty"`
	RequestLevel         *uint32 `protobuf:"varint,4,opt,name=request_level" json:"request_level,omitempty"`
}

type CMsgGCCStrike15_v2_PlayersProfile struct {
	RequestId       *uint32                                         `protobuf:"varint,1,opt,name=request_id" json:"request_id,omitempty"`
	AccountProfiles []*CMsgGCCStrike15_v2_MatchmakingGC2ClientHello `protobuf:"bytes,2,rep,name=account_profiles" json:"account_profiles,omitempty"`
}

type CMsgGCCStrike15_v2_MatchmakingGC2ClientHello struct {
	AccountId            *uint32                                         `protobuf:"varint,1,opt,name=account_id" json:"account_id,omitempty"`
	OnGoingMatch         *CMsgGCCStrike15_v2_MatchmakingGC2ClientReserve `protobuf:"bytes,2,opt,name=ongoingmatch" json:"ongoingmatch,omitempty"`
	GlobalStats          *GlobalStatistics                               `protobuf:"bytes,3,opt,name=global_stats" json:"global_stats,omitempty"`
	PenaltySeconds       *uint32                                         `protobuf:"varint,4,opt,name=penalty_seconds" json:"penalty_seconds,omitempty"`
	PenaltyReason        *uint32                                         `protobuf:"varint,5,opt,name=penalty_reason" json:"penalty_reason,omitempty"`
	VacBanned            *int32                                          `protobuf:"varint,6,opt,name=vac_banned" json:"vac_banned,omitempty"`
	Ranking              *PlayerRankingInfo                              `protobuf:"bytes,7,opt,name=ranking" json:"ranking,omitempty"`
	Commendation         *PlayerCommendationInfo                         `protobuf:"bytes,8,opt,name=commendation" json:"commendation,omitempty"`
	Medals               *PlayerMedalsInfo                               `protobuf:"bytes,9,opt,name=medals" json:"medals,omitempty"`
	MyCurentEvent        *TournamentEvent                                `protobuf:"bytes,10,opt,name=my_current_event" json:"my_current_event,omitempty"`
	MyCurrentEventTeams  []*TournamentTeam                               `protobuf:"bytes,11,rep,name=my_current_event_teams" json:"my_current_event_teams,omitempty"`
	MyCurrentTeam        *TournamentTeam                                 `protobuf:"bytes,12,opt,name=my_current_team" json:"my_current_team,omitempty"`
	MyCurrentEventStages []*TournamentEvent                              `protobuf:"bytes,13,rep,name=my_current_event_stages" json:"my_current_event_stages,omitempty"`
	SurveyVote           *uint32                                         `protobuf:"varint,14,opt,name=suvey_vote" json:"survey_vote,omitempty"`
	Activity             *AccountActivity                                `protobuf:"bytes,15,opt,name=activity" json:"activity,omitempty"`
	PlayerLevel          *int32                                          `protobuf:"varint,17,opt,name=player_level" json:"player_level,omitempty"`
	PlayerCurXp          *int32                                          `protobuf:"varint,18,opt,name=player_cur_xp" json:"player_cur_xp,omitempty"`
	PlayerXpBonusFlags   *int32                                          `protobuf:"varint,19,opt,name=player_xp_bonus_flags" json:"player_xp_bonus_flags,omitempty"`
}

type GlobalStatistics struct {
	PlayersOnline           *uint32                    `protobuf:"varint,1,opt,name=players_online" json:"players_online,omitempty"`
	ServersOnline           *uint32                    `protobuf:"varint,2,opt,name=servers_online" json:"servers_online,omitempty"`
	PlayersSearching        *uint32                    `protobuf:"varint,3,opt,name=players_searching" json:"players_searching,omitempty"`
	ServersAvailable        *uint32                    `protobuf:"varint,4,opt,name=servers_available" json:"servers_available,omitempty"`
	OngoingMatches          *uint32                    `protobuf:"varint,5,opt,name=ongoing_matches" json:"ongoing_matches,omitempty"`
	SearchTimeAvg           *uint32                    `protobuf:"varint,6,opt,name=search_time_avg" json:"search_time_avg,omitempty"`
	SearchStatistics        []*DetailedSearchStatistic `protobuf:"bytes,7,rep,name=search_statistics" json:"search_statistics,omityempty"`
	MainPostUrl             *string                    `protobuf:"bytes,8,opt,name=main_post_url" json:"main_post_url,omitempty"`
	RequiredAppIdVersion    *uint32                    `protobuf:"varint,9,opt,name=required_appid_version" json:"required_appid_version,omitempty"`
	PricesheetVersion       *uint32                    `protobuf:"varint,10,opt,name=pricesheet_version" json:"pricesheet_version,omitempty"`
	TwitchStreamsVersion    *uint32                    `protobuf:"varint,11,opt,name=twitch_streams_version" json:"twitch_streams_version,omitempty"`
	ActiveTournamentEventId *uint32                    `protobuf:"varint,12,opt,name=active_tournament_eventid" json:"active_tournament_eventid,omitempty"`
	ActiveSurveyId          *uint32                    `protobuf:"varint,13,opt,name=active_survey_id", json:"active_survey_id,omitempty"`
}

type CMsgGCCStrike15_v2_MatchmakingGC2ClientReserve struct {
	ServerId    *uint64                                         `protobuf:"varint,1,opt,name=serverid" json:"serverid,omitempty"`
	ServerIp    *uint32                                         `protobuf:"varint,2,opt,name=serverip" json:"serverip,omitempty"`
	ServerPort  *uint32                                         `protobuf:"varint,3,opt,name=serverport" json:"serverport,omitempty"`
	Reservation *CMsgGCCStrike15_v2_MatchmakingGC2ServerReserve `protobuf:"bytes,4,opt,name=reservation" json:"reservation,omitempty"`
	Map         *string                                         `protobuf:"bytes,5,opt,name=map" json:"map,omitempty"`
}

type PlayerRankingInfo struct {
	AccountId  *uint32  `protobuf:"varint,1,opt,name=account_id" json:"account_id,omitempty"`
	RankId     *uint32  `protobuf:"varint,2,opt,name=rank_id", json:"rank_id,omitempty"`
	Wins       *uint32  `protobuf:"varint,3,opt,name=wins" json:"wins,omitempty"`
	RankChange *float32 `protobuf:"fixed32,4,opt,name=rank_change" json:"rank_change,omitempty"`
}

type PlayerCommendationInfo struct {
	CmdFriendly *uint32 `protobuf:"varint,1,opt,name=cmd_friendly" json:"cmd_friendly,omitempty"`
	CmdTeaching *uint32 `protobuf:"varint,2,opt,name=cmd_teaching" json:"cmd_teaching,omitempty"`
	CmdLeader   *uint32 `protobuf:"varint,4,opt,name=cmd_leader" json:"cmd_leader,omitempty"`
}

type PlayerMedalsInfo struct {
	MedalTeam                 *uint32  `protobuf:"varint,1,opt,name=medal_team" json:"medal_team,omitempty"`
	MedalCombat               *uint32  `protobuf:"varint,2,opt,name=medal_combat" json:"medal_combat,omitempty"`
	MedalWeapon               *uint32  `protobuf:"varint,3,opt,name=medal_weapon" json:"medal_weapon,omitempty"`
	MedalGlobal               *uint32  `protobuf:"varint,4,opt,name=medal_global" json:"medal_global,omitempty"`
	MedalArms                 *uint32  `protobuf:"varint,5,opt,name=medal_arms" json:"medal_arms,omitempty"`
	DisplayItemsDefidx        []uint32 `protobuf:"varint,7,rep,name=display_items_defidx" json:"display_items_defidx,omitempty"`
	FeaturedDisplayItemDefidx *uint32  `protobuf:"varint,8,opt,name=featured_display_item_defidx" json:"featured_display_item_defidx"`
}

type TournamentEvent struct {
	EventId         *int32  `protobuf:"varint,1,opt,name=event_id" json:"event_id,omitempty"`
	EventTag        *string `protobuf:"bytes,2,opt,name=event_tag" json:"event_tag,omitempty"`
	EventName       *string `protobuf:"bytes,3,opt,name=event_name" json:"event_name,omitempty"`
	EventTimeStart  *uint32 `protobuf:"varint,4,opt,name=event_time_start" json:"event_time_start,omitempty"`
	EventTimeEnd    *uint32 `protobuf:"varint,5,opt,name=event_time_end" json:"event_time_end,omitempty"`
	EventPublic     *int32  `protobuf:"varint,6,opt,name=event_public" json:"event_public,omitempty"`
	EventStageId    *int32  `protobuf:"varint,7,opt,name=event_stage_id" json:"event_stage_id,omitempty"`
	EventStageName  *string `protobuf:"bytes,8,opt,name=event_stage_name" json:"event_stage_name,omitempty"`
	ActiveSectionId *uint32 `protobuf:"varint,9,opt,name=active_Secton_id" json:"active_section_id,omitempty"`
}

type TournamentTeam struct {
	TeamId   *int32              `protobuf:"varint,1,opt,name=team_id" json:"team_id,omitempty"`
	TeamTag  *string             `protobuf:"bytes,2,opt,name=team_tag" json:"team_tag,omitempty"`
	TeamFlag *string             `protobuf:"bytes,3,opt,name=team_flag" json:"team_flag,omitempty"`
	TeamName *string             `protobuf:"bytes,4,opt,name=team_name" json:"team_name,omitempty"`
	Players  []*TournamentPlayer `protobuf:"bytes,5,opt,name=players" json:"players,omitempty"`
}

type TournamentPlayer struct {
	AccountId      *uint32 `protobuf:"varint,1,opt,name=account_id" json:"account_id,omitempty"`
	PlayerNick     *string `protobuf:"bytes,2,opt,name=player_nick" json:"player_nick,omitempty"`
	PlayerName     *string `protobuf:"bytes,3,opt,name=player_name" json:"player_name,omitempty"`
	PlayerDOB      *uint32 `protobuf:"varint,4,opt,name=player_dob" json:"player_dob,omitempty"`
	PlayerFlag     *string `protobuf:"bytes,5,opt,name=player_flag" json:"player_flag,omitempty"`
	PlayerLocation *string `protobuf:"bytes,6,opt,name=player_location" json:"player_location,omitempty"`
	PlayerDesc     *string `protobuf:"bytes,7,opt,name="player_desc" json:"player_desc,omitempty`
}

type AccountActivity struct {
	Activity *uint32 `protobuf:"varint,1,opt,name=activity" json:"activity,omitempty"`
	Mode     *uint32 `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
	Map      *uint32 `protobuf:"varint,3,opt,name=map" json:"map,omitempty"`
}

type DetailedSearchStatistic struct {
	GameType         *uint32 `protobuf:"varint,1,opt,name=game_type" json:"game_type,omitempty"`
	SearchTimeAvg    *uint32 `protobuf:"varint,2,opt,name=search_time_avg" json:"search_time_avg,omitempty"`
	PlayersSearching *uint32 `protobuf:"varint,4,opt,name=players_searching" json:"players_searching,omitempty"`
}

type CMsgGCCStrike15_v2_MatchmakingGC2ServerReserve struct {
	AccountIds                  []uint32             `protobuf:"varint,1,rep,name=account_ids" json:"account_ids,omitempty"`
	GameType                    *uint32              `protobuf:"varint,2,opt,name=game_type" json:"game_type,omitempty"`
	MatchId                     *uint64              `protobuf:"varint,3,opt,name=match_id" json:"match_id,omitempty"`
	ServerVersion               *uint32              `protobuf:"varint,4,opt,name=server_version" json:"server_version,omitempty"`
	Rankings                    []*PlayerRankingInfo `protobuf:"bytes,5,rep,name=rankings" json:"rankings,omitempty"`
	EncryptionKey               *uint64              `protobuf:"varint,6,opt,name=encryption_key" json:"encryption_key,omitempty"`
	EncryptionKeyPub            *uint64              `protobuf:"varint,7,opt,name=encryption_key_pub" json:"encryption_key_pub,omitempty"`
	PartyIds                    []uint32             `protobuf:"varint,8,rep,name=party_ids" json:"party_ids,omitempty"`
	Whitelist                   []*IpAddressMask     `protobuf:"bytes,9,rep,name=whitelist" json:"whitelist,omitempty"`
	TvMasterSteamid             *uint64              `protobuf:"varint,10,opt,name=tv_master_steamid" json:"tv_master_steamid,omitempty"`
	TournamentEvent             *TournamentEvent     `protobuf:"bytes,11,opt,name=tournament_event" json:"tournamnet_event,omitempty"`
	TournamentTeams             []*TournamentTeam    `protobuf:"bytes,12,rep,name=tournament_teams" json:"tournament_teams,omitempty"`
	TournamentCastersAccountIds []uint32             `protobuf:"varint,13,rep,name=tournament_casters_account_ids" json:"tournament_casters_account_ids,omitempty"`
	TvRelaySteamid              *uint32              `protobuf:"varint,14,opt,name=tv_relay_steamid" json:"tv_relay_steamid,omitempty"`
	PreMatchData                *CPreMatchInfoData   `protobuf:"bytes,15,opt,name=pre_match_data" json:"pre_match_data,omitempty"`
}

type IpAddressMask struct {
	A     *uint32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B     *uint32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	C     *uint32 `protobuf:"varint,3,opt,name=c" json:"c,omitempty"`
	D     *uint32 `protobuf:"varint,4,opt,name="d json:"d,omitempty"`
	Bits  *uint32 `protobuf:"varint,5,opt,name=bits" json:"bits,omitempty"`
	Token *uint32 `protobuf:"varint,6,opt,name=token" json:"token,omitempty"`
}

type CPreMatchInfoData struct {
	PredictionsPCT *int32                                    `protobuf:"varint,1,opt,name=predictions_pct" json:"predictions_pct,omitempty"`
	Draft          *CDataGCCStrike15_v2_TournamentMatchDraft `protobuf:"bytes,4,opt,name=draft" json:"draft,omitempty"`
	Stats          []*CPreMatchInfoData_TeamStats            `protobuf:"bytes,5,rep,name=stats" json:"stats,omitempty"`
}

type CPreMatchInfoData_TeamStats struct {
	MatchInfoIdxtxt *int32  `protobuf:"varint,1,opt,name=match_info_idxtxt" json:"match_info_idxtxt,omitempty"`
	MatchInfoTxt    *string `protobuf:"bytes,2,opt,name=match_info_txt" json:"match_info_txt,omitempty"`
	MatchInfoTeams  *string `protobuf:"bytes,3,rep,name=match_info_teams" json:"match_info_teams"`
}

type CDataGCCStrike15_v2_TournamentMatchDraft struct {
	EventId      *int32                                            `protobuf:"varint,1,opt,name=event_id" json:"event_id,omitempty"`
	EventStageId *int32                                            `protobuf:"varint,2,opt,name=event_stage_id" json:"event_stage_id,omitempty"`
	TeamId0      *int32                                            `protobuf:"varint,3,opt,name=team_id_0" json:"team_id_0,omitempty"`
	TeamId1      *int32                                            `protobuf:"varint,4,opt,name=team_id_1" json:"team_id_1,omitempy"`
	MapsCount    *int32                                            `protobuf:"varint,5,opt,name=maps_count" json:"maps_count,omitempty"`
	MapsCurrent  *int32                                            `protobuf:"varint,6,opt,name=maps_current" json:"maps_current,omitempty"`
	TeamIdStart  *int32                                            `protobuf:"varint,7,opt,name=team_id_start" json:"team_id_start,omitempty"`
	TeamIdVeto1  *int32                                            `protobuf:"varint,8,opt,name=team_id_veto1" json:"team_id_veto1,omitempty"`
	TeamIdPickn  *int32                                            `protobuf:"varint,9,opt,name=team_id_pickn" json:"team_id_pickn,omitempty"`
	Drafts       []*CDataGCCStrike15_v2_TournamentMatchDraft_Entry `protobuf:"bytes,10,rep,name=drafts" json:"drafts,omitempty"`
}

type CDataGCCStrike15_v2_TournamentMatchDraft_Entry struct {
	MapId    *int32 `protobuf:"varint,1,opt,name=mapid" json:"mapid,omitempty"`
	TeamIdCT *int32 `protobuf:"varint,2,opt,name=team_id_ct" json:"team_id_ct,omitempty"`
}

func (m *CDataGCCStrike15_v2_TournamentMatchDraft_Entry) Reset() {
	*m = CDataGCCStrike15_v2_TournamentMatchDraft_Entry{}
}

func (m *CDataGCCStrike15_v2_TournamentMatchDraft_Entry) String() string {
	return proto.CompactTextString(m)
}

func (*CDataGCCStrike15_v2_TournamentMatchDraft_Entry) ProtoMessage() {}

func (m *CDataGCCStrike15_v2_TournamentMatchDraft) Reset() {
	*m = CDataGCCStrike15_v2_TournamentMatchDraft{}
}

func (m *CDataGCCStrike15_v2_TournamentMatchDraft) String() string {
	return proto.CompactTextString(m)
}

func (*CDataGCCStrike15_v2_TournamentMatchDraft) ProtoMessage() {}

func (m *CPreMatchInfoData) Reset() {
	*m = CPreMatchInfoData{}
}

func (m *CPreMatchInfoData) String() string {
	return proto.CompactTextString(m)
}

func (*CPreMatchInfoData) ProtoMessage() {}

func (m *CPreMatchInfoData_TeamStats) Reset() {
	*m = CPreMatchInfoData_TeamStats{}
}

func (m *CPreMatchInfoData_TeamStats) String() string {
	return proto.CompactTextString(m)
}

func (*CPreMatchInfoData_TeamStats) ProtoMessage() {}

func (m *IpAddressMask) Reset() {
	*m = IpAddressMask{}
}

func (m *IpAddressMask) String() string {
	return proto.CompactTextString(m)
}

func (*IpAddressMask) ProtoMessage() {}

func (m *CMsgGCCStrike15_v2_MatchmakingGC2ServerReserve) Reset() {
	*m = CMsgGCCStrike15_v2_MatchmakingGC2ServerReserve{}
}

func (m *CMsgGCCStrike15_v2_MatchmakingGC2ServerReserve) String() string {
	return proto.CompactTextString(m)
}

func (*CMsgGCCStrike15_v2_MatchmakingGC2ServerReserve) ProtoMessage() {}

func (m *AccountActivity) Reset() {
	*m = AccountActivity{}
}

func (m *AccountActivity) String() string {
	return proto.CompactTextString(m)
}

func (*AccountActivity) ProtoMessage() {}

func (m *TournamentPlayer) Reset() {
	*m = TournamentPlayer{}
}

func (m *TournamentPlayer) String() string {
	return proto.CompactTextString(m)
}

func (TournamentPlayer) ProtoMessage() {}

func (m *TournamentTeam) Reset() {
	*m = TournamentTeam{}
}

func (m *TournamentTeam) String() string {
	return proto.CompactTextString(m)
}

func (*TournamentTeam) ProtoMessage() {}

func (m *TournamentEvent) Reset() {
	*m = TournamentEvent{}
}

func (m *TournamentEvent) String() string {
	return proto.CompactTextString(m)
}

func (*TournamentEvent) ProtoMessage() {}

func (m *PlayerMedalsInfo) Reset() {
	*m = PlayerMedalsInfo{}
}

func (m *PlayerMedalsInfo) String() string {
	return proto.CompactTextString(m)
}

func (*PlayerMedalsInfo) ProtoMessage() {}

func (m *PlayerCommendationInfo) Reset() {
	*m = PlayerCommendationInfo{}
}

func (m *PlayerCommendationInfo) String() string {
	return proto.CompactTextString(m)
}

func (*PlayerCommendationInfo) ProtoMessage() {}

func (m *PlayerRankingInfo) Reset() {
	*m = PlayerRankingInfo{}
}

func (m *PlayerRankingInfo) String() string {
	return proto.CompactTextString(m)
}

func (*PlayerRankingInfo) ProtoMessage() {}

func (m *GlobalStatistics) Reset() {
	*m = GlobalStatistics{}
}

func (m *GlobalStatistics) String() string {
	return proto.CompactTextString(m)
}

func (*GlobalStatistics) ProtoMessage() {}

func (m *CMsgGCCStrike15_v2_MatchmakingGC2ClientHello) Reset() {
	*m = CMsgGCCStrike15_v2_MatchmakingGC2ClientHello{}
}

func (m *CMsgGCCStrike15_v2_MatchmakingGC2ClientHello) String() string {
	return proto.CompactTextString(m)
}

func (*CMsgGCCStrike15_v2_MatchmakingGC2ClientHello) ProtoMessage() {}

func (m *CMsgGCCStrike15_v2_MatchmakingGC2ClientReserve) Reset() {
	*m = CMsgGCCStrike15_v2_MatchmakingGC2ClientReserve{}
}

func (m *CMsgGCCStrike15_v2_MatchmakingGC2ClientReserve) String() string {
	return proto.CompactTextString(m)
}

func (*CMsgGCCStrike15_v2_MatchmakingGC2ClientReserve) ProtoMessage() {}

func (m *CMsgGCCStrike15_v2_ClientRequestPlayersProfile) Reset() {
	*m = CMsgGCCStrike15_v2_ClientRequestPlayersProfile{}
}

func (m *CMsgGCCStrike15_v2_ClientRequestPlayersProfile) String() string {
	return proto.CompactTextString(m)
}

func (*CMsgGCCStrike15_v2_ClientRequestPlayersProfile) ProtoMessage() {}

func (m *CMsgGCCStrike15_v2_PlayersProfile) Reset() {
	*m = CMsgGCCStrike15_v2_PlayersProfile{}
}

func (m *CMsgGCCStrike15_v2_PlayersProfile) String() string {
	return proto.CompactTextString(m)
}

func (*CMsgGCCStrike15_v2_PlayersProfile) ProtoMessage() {}

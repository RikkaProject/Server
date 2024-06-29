package login

import (
	"HeroServer/db"
	"HeroServer/proto"
	"HeroServer/service/mail"
	"HeroServer/service/role"
	"fmt"
	"strconv"
)

func HandleLogin(buf *proto.Buffer, pt *proto.Proto) bool {
	account := buf.ReadString()        //账号
	password := buf.ReadString()       //密码
	timestamp := buf.ReadInt()         //时间
	region := buf.ReadString()         //区域
	platform := buf.ReadString()       //平台
	platform2 := buf.ReadString()      //平台
	string1 := buf.ReadString()        //未知字符串
	serverId := buf.ReadInt()          //服务器id
	string2 := buf.ReadString()        //未知字符串 base64编码的 (base64)
	deviceId := buf.ReadString()       //可能是设备id (base64)
	deviceMark := buf.ReadString()     //设备制造商 (base64)
	string3 := buf.ReadString()        //未知 (base64)
	androidVersion := buf.ReadString() //安卓版本 (base64)
	deviceModel := buf.ReadString()    //设备型号 (base64)
	androidId := buf.ReadString()      //可能是安卓id (base64)
	byte1 := buf.ReadByte()            //未知
	byte2 := buf.ReadByte()            //未知S
	byte3 := buf.ReadByte()            //未知
	int1 := buf.ReadInt()              //未知
	appVersion := buf.ReadString()     // 应用版本
	resVersion := buf.ReadString()     // 可能是资源版本
	encodeMode := buf.ReadString()     // 加密方法
	fmt.Println("登录数据：", account, password, timestamp, region, platform, platform2, string1, serverId, string2, deviceId, deviceMark, string3, androidVersion, deviceModel, androidId, byte1, byte2, byte3, int1, appVersion, resVersion, encodeMode)
	rid, _ := strconv.Atoi(account)
	_role := role.GetRole(rid)
	if _role != nil {
		pt.Rid = rid
		pt.InitRoleData(_role)
		pt.Login(_role)
		return true
	} else {
		//登录失败
		fmt.Println("登录失败")
		return false
	}
	//pt.Login()
}

func HandleStartRecord(pt *proto.Proto) {
	pt.SendStepOne()
	pt.SendStepTwo()
	pt.SendTaskNewInit()
	pt.MPtTaskDailyInit()
	pt.MsgPlayerAssistRole() //12032
	pt.SendRoleInfo()
	sendCurrencyNumber(pt)
	pt.MsgCurrencyCurCrucibleInit() //18007
	sendItemAll(pt)
	sendNewBagInit(pt)
	pt.SendEquipInit()
	pt.MPtStorageInit()
	pt.MPtOtherRoleItemAll()
	pt.MsgRaidEquipAllItem()   //55101
	pt.MsgRaidEquipBagInit()   //55201
	pt.MsgRaidEquipEquipInit() //55301
	pt.MsgPassInfo()           //40001
	pt.MPtCostumeUnlock()
	pt.MPtAdventureLevelInfo()
	pt.MPtAdventureWeeklyRewardInfo()
	pt.MPtBuyGoldTimes()
	pt.MsgStoreNormalInfo()                //26003
	pt.MsgStoreRandRefreshInfo()           //26008
	pt.MsgStoreRandInfo()                  //26005
	pt.MsgStoreRandCardRefreshInfo()       //26018
	pt.MsgStoreRandCardInfo()              //26015
	pt.MsgStoreRandPetRefreshInfo()        //26024
	pt.MsgStoreRandPetInfo()               //26021
	pt.MsgStoreRandPetSkinRefreshInfo()    //26028
	pt.MsgStoreRandPetRefreshInfo()        //26025
	pt.MsgStoreRandArmorRefreshInfo()      //26032
	pt.MsgStoreRandArmorInfo()             //26029
	pt.MsgStoreRandSuperarmorRefreshInfo() //26036
	pt.MsgStoreRandSuperarmorInfo()        //26033
	pt.MPtShareScreenShotInfo()
	pt.MPtCardList()
	pt.MPtCardCollectionInfo()
	pt.MPtDungeonRiftAffix()
	pt.MPtDungeonCardAffix()
	pt.MPtLotteryFailCount()
	pt.MsgRechargeDoubleInfo()    //41014
	pt.MsgDungeonKeyRewardTimes() //13601
	pt.MPtRechargeTotalMoney()
	pt.MsgDungeonRankKeyTimes() //13527
	pt.MPtAbyssRewardLayer()
	pt.MPtRankSingleFirstPassAwards()
	pt.MsgPlayerBestScore() //13034
	pt.MPtDungeonCardLevelStar()
	pt.MPtDungeonCardBindInfo()
	pt.MPtDungeonCardRewardTimes()
	pt.MsgActivityPersonalActivityInfo()
	pt.MPtActivityInfo()
	pt.MPtActivityExchangeInfo()
	pt.MPtActivityRechargeInfo()
	pt.MPtActivityTaskInfo()
	pt.MPtActivityLotteryInit()
	pt.MPtActivityAccInit()
	pt.MsgActivityPoolInfo()
	pt.MPtActivityCurrencyInit()
	pt.MPtActivityComicInfo()
	pt.MPtActivityTimeLoginInit()
	pt.MPtActivityTimerLoginInfo()
	pt.MsgActivityInviteInfo()
	pt.MPtActivityBingoInit()
	pt.MPtLotteryFreeinfo()
	pt.MPtFestLetterInfo()
	pt.MPtActivityDiyGiftInfo()
	pt.MPtActDungeonInit()
	pt.MPtActivityNoticeInit()
	pt.MPtEquipCollectionInit()
	pt.MPtChatShieldInfo()
	pt.MPtChatShieldPlayerInfo()
	pt.MPtChatMaxLimit()
	pt.MPtActivityGiftInfo()
	pt.MPtGiftRandom()
	pt.MPtArenaWeekAmount()
	pt.MPtArenaInfo()
	pt.MPtArenaGroupPlayerInfo()
	pt.MPtSignIn7DaysInit()
	pt.MPtSignInDailyInit()
	pt.MPtAdInfo()
	pt.MPtAdLotteryInfo()
	pt.MPtBattlePassStoreInfo()
	pt.MPtBattlePassInfo()
	pt.MPtNewInnersBestLvTime()
	pt.MPtRegRewardInfo()
	pt.MPtCommunityShareInfo()
	pt.MPtVersionsTipsInfo()
	pt.MPtPlayerKVInit()
	pt.MPtNewAFKBackInfo()
	pt.MPtPetInit()
	pt.MPtPetExploreInfo()
	pt.MPtCareerTaskInit()
	pt.MPtActivityCacheInfo()
	pt.MPtArenaTopInfo()
	pt.MPtPetBattleTotalCoin()
	pt.MPtBDIInfo()
	pt.MPtBDIElementInfo()
	pt.MPtSLCostumeBuild()
	pt.MPtNewInnersRandom()
	pt.MPtLobbyTriggerDaily()
	pt.MPtRechargeMailInfo()
	pt.MPtPlanesGateInfo()
	pt.MPtBraveTeamPlayer()
	pt.MPtFashionPkInfo()
	pt.MPtMagicInfo()
	pt.MPtPlayerCostume()
	pt.MPtRankBestLevel()
	pt.MPtAttrInit()
	pt.MPtPassiveSkill()
	pt.MPtCardEquipInfoList()
	pt.MsgHeroCityInit()
	pt.MsgPlayerSkillInit()
	pt.MsgFunctionUnlock()
	//----
	pt.ProtoMptrankdungeonstarParse()
	pt.MsgActivityOwnExchangeInfo()
	pt.MsgActivityPotInfo()
	pt.MsgActivityGiftcardInfo()
	pt.MsgActivityBlindboxInfo()
	pt.MsgActivityBlindboxCodes()
	pt.MsgActivityCostLoginInfo()
	pt.MsgActivityRedEnvelopeInfo()
	pt.MsgActivityRechargeOnInfo()
	pt.MsgActivityMagicTower()
	pt.MsgActivityBpSginInfo()
	pt.MsgActivityPaintedInfo()
	pt.MsgActivityRouletteInfo()
	pt.MsgLuckytuInfo()
	pt.MsgFashionInfo()
	pt.MsgRebatermbInfo()
	pt.MsgActivityAnswerInfo()
	pt.MsgActivityTurntableInfo()
	pt.MsgDailywishInit()
	pt.MsgCardMatchInfo()
	pt.MsgHundredRaceInfo()
	pt.MsgPreOrderInfo()
	pt.MsgCookingInitInfo()
	pt.MsgCrazyLotteryInfo()
	pt.MsgLoveliveInitInfo()
	pt.MsgActivityFactoryInit()
	pt.MsgActivityStoreInfo()
	pt.MsgActivityDailyGiftInfo()
	pt.MsgLuckyBagInitInfo()
	pt.MsgActivitySpExploreInfo()
	pt.MsgGuildDonateInfo()
	pt.MsgGuildListRefreshTime()
	pt.MsgGuildStoreInfo()
	pt.MsgGuildBossTown()
	pt.MsgGuildAdvPersonalInfo()
	pt.MsgGuildAdvMonLayerPetUsed()
	pt.MsgDungeonWitchInfo()
	pt.MsgTotalRechargeAwardInit()
	pt.MsgTotalRechargeRmbAwardInit()
	pt.MsgRechargeMonthCardInfo()
	pt.MsgMatchPunishCd()
	pt.MsgMatchBeInviteBlock()
	pt.MsgDungeonWorldBossForeshowAffix()
	pt.MsgDungeonWorldBossAffix()
	pt.MsgDungeonWorldBossCd()
	pt.MsgNewInnersInfo()
	pt.MsgNewInnersStageRewardInfo()
	pt.MsgPlayerEventRewardInfo()
	pt.MsgPlayerExp()
	pt.MsgPlayerVersionRewardInfo()
	pt.MsgAstrologyAnswer()

	pt.MsgHeadRecordUpdate()
	pt.MsgHeadInfo()
	pt.MsgEmojiInfo()
	pt.MsgShareStoreInit()
	pt.MsgDungeonSolomonInfo()
	pt.MsgSoulballInfo()
	pt.MsgFundInfo()
	pt.MsgActivityVoteInfo()
	pt.MsgGiftPushGiftInfo()
	pt.MsgNewAbyssHuntInfo()
	pt.MsgNewAbyssHuntBossInfo()
	pt.MsgItemWeaponLinkInfo()
	pt.MsgRaidEquipStarChessInfo()
	pt.MsgRaidEquipStarRingAndTrackInfo()
	pt.MsgRaidEquipStarIllustrated()

	pt.MsgRechargeDungeonTeamCardInfo()
	pt.MsgTaskRaidHonorInfo()
	pt.MsgRaidEquipArgs()
	pt.MsgPlayerAbTest()
	pt.MsgGemInfo()
	pt.MsgGemFreeCombineInfo()
	pt.MsgGemResonanceInfo()
	pt.MsgPlayerNewTutorialInfo()
	pt.MsgTriggerBoxInfoDailyRemain()
	pt.MsgTriggerBoxInfoOnce()
	pt.MsgGemChooseRateUp()
	pt.MsgGemWellInfo()
	pt.MsgWorkshopInfo()
	pt.MsgTechInfo()
	pt.MsgGuildWishProUpdateTimes()
	pt.MsgGuildWishProPlay()
	pt.MsgGuildProStorePlayerInfo()
	pt.MsgWitchInviteInfo()
	pt.MsgHeroHouseLetter()
	pt.MsgHeroHouseName()
	pt.MsgHeroHouseInfo()
	pt.MsgHeroHouseDecorateInfo()
	pt.MsgStakeDamageGraphInfo()
	pt.MsgStakeRankInfo()
	pt.MsgStakePersonalInfo()
	pt.MsgStakeStrongInfo()

	pt.MsgRechargeNewestFirstInfo()
	pt.MsgRubbishBagUpdate()
	pt.MsgNewStoryInfo()
	pt.MsgGuildPetFetterList()
	pt.MsgGuildPetFetterInfo()
	pt.MsgGuildGvgPlayerInfo()
	pt.MsgGuildGvgRewardInfo()
	pt.MsgGuildGvgSignUpState()
	pt.MsgGuildGvgNoticeOp()
	pt.MsgCurrencyCashLog()
	pt.MsgPlayerBtnSnapshotInfo()
	pt.MsgNewStrongInit()
	pt.MsgCrossWorldBossPlayerInfo()

	pt.MsgLvPassiveSkills()
	pt.MsgEquipSnapshots()
	pt.MsgHistoryBookAll()
	pt.MsgPlayerDungeonLv()

	pt.MsgDungeonSkillTrialInfo()
	pt.MsgArmorUplevelInit()
	pt.MsgNoobActInfo()
	pt.MsgActivityPlayerActivityInfo()
	pt.MsgActivityPlayerActivityLotteryInfo()
	pt.MsgNewRuneInfo()
	pt.MsgNewRuneFirstCombine()
	pt.MsgInnateInfo()
	mail.SendMailCount(pt)
	//---
	pt.InitFinish()
}

func sendNewBagInit(pt *proto.Proto) {
	var items []*db.HeroItem
	db.Conn.Where(&db.HeroItem{Rid: pt.Rid}).Find(&items)
	pt.SendNewBagInit(items)
}

func sendCurrencyNumber(pt *proto.Proto) {
	var items []*db.Currency
	db.Conn.Where(&db.Currency{Rid: pt.Rid}).Find(&items)
	pt.SendCurrencyNumber(items)
}

func sendItemAll(pt *proto.Proto) {
	var items []*db.HeroItem
	db.Conn.Where(&db.HeroItem{Rid: pt.Rid}).Find(&items)
	pt.SendItemAll(items)
}

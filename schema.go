package pocket

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type liveListReq struct {
	ExtMsgType string `json:"extMsgType"` // "USER_LIVE"
	RoomId     string `json:"roomId"`     // ""
	OwnerId    string `json:"ownerId"`
	NextTime   int64  `json:"nextTime"` // 0
}

func newLiveListReq(ownerId string, nextTime int64) liveListReq {
	return liveListReq{
		ExtMsgType: "USER_LIVE",
		RoomId:     "",
		OwnerId:    ownerId,
		NextTime:   nextTime,
	}
}

type Resp[T any] struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content T      `json:"content"`
}

type UserInfoItem struct {
	UserId       int    `json:"userId"`
	Nickname     string `json:"nickname"` // this!
	Avatar       string `json:"avatar"`
	Exp          int    `json:"exp"`
	Level        int    `json:"level"`
	Gender       int    `json:"gender"`
	Birthday     string `json:"birthday"`
	City         string `json:"city"`
	Verification bool   `json:"verification"`
	Money        int    `json:"money"`
	Support      int    `json:"support"`
	Permission   struct {
		Post struct {
			View         bool        `json:"view"`
			Create       bool        `json:"create"`
			Update       bool        `json:"update"`
			Delete       bool        `json:"delete"`
			ManagerGroup interface{} `json:"managerGroup"`
			ManagerTeam  interface{} `json:"managerTeam"`
		} `json:"post"`
	} `json:"permission"`
	RoleName string `json:"roleName"`
	RoleId   int    `json:"roleId"`
	DeviceId string `json:"deviceId"`
	BindInfo []struct {
		BindType string `json:"bindType"`
		UniqueId string `json:"uniqueId"`
		Nickname string `json:"nickname"`
	} `json:"bindInfo"`
	BadgeCount   int         `json:"badgeCount"`
	Friends      int         `json:"friends"`
	Followers    int         `json:"followers"`
	Token        interface{} `json:"token"`
	BigSmallInfo struct {
		Relationship  bool        `json:"relationship"`
		BigUserInfo   interface{} `json:"bigUserInfo"`
		SmallUserInfo interface{} `json:"smallUserInfo"`
	} `json:"bigSmallInfo"`
	CommentStatus     int           `json:"commentStatus"`
	BgImg             string        `json:"bgImg"`
	Badge             []interface{} `json:"badge"`
	Vip               bool          `json:"vip"`
	TeamLogo          interface{}   `json:"teamLogo"`
	Card              int           `json:"card"`
	ExpArr            []int         `json:"expArr"`
	PfUrl             string        `json:"pfUrl"`
	EditImg           string        `json:"editImg"`
	EditName          string        `json:"editName"`
	EditBgImg         string        `json:"editBgImg"`
	TeenagersPassword interface{}   `json:"teenagersPassword"`
	Adult             bool          `json:"adult"`
	ContinueAuth      bool          `json:"continueAuth"`
	OutOfCn           bool          `json:"outOfCn"`
	ValidTime         int           `json:"validTime"`
	TeenagersTips     string        `json:"teenagersTips"`
	PMobile           string        `json:"pMobile"`
}

type LiveListItem struct {
	MsgidClient string `json:"msgidClient"`
	MsgTime     int64  `json:"msgTime"`
	MsgType     string `json:"msgType"`
	Bodys       string `json:"bodys"`
	ExtInfo     string `json:"extInfo"`
	Privacy     bool   `json:"privacy"`
}

type LiveExtInfo struct {
	ID             int64    `json:"id"`
	CoverUrl       string   `json:"coverUrl"`
	CoverUrlList   []string `json:"coverUrlList"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	Url            string   `json:"url"`
	JumpType       string   `json:"jumpType"`
	JumpPath       string   `json:"jumpPath"`
	ThirdAppName   string   `json:"thirdAppName"`
	ThirdAPPImgUrl string   `json:"thirdAPPImgUrl"`
	MessageType    string   `json:"messageType"`
	User           struct {
		UserId   int    `json:"userId"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	} `json:"user"`
}

func (l LiveListItem) FormatToLiveItem() LiveItem {
	var extInfo LiveExtInfo
	_ = json.Unmarshal([]byte(l.ExtInfo), &extInfo)
	return LiveItem{
		ID:    strconv.FormatInt(extInfo.ID, 10),
		Title: extInfo.Title,
		Time:  time.UnixMilli(l.MsgTime),
	}
}

type LiveItem struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Time  time.Time `json:"liveTime"`
}

type LiveListContent struct {
	Message  []LiveListItem `json:"message"`
	NextTime int64          `json:"nextTime"`
}

type LiveInfoContent struct {
	LiveId         string `json:"liveId"`
	RoomId         string `json:"roomId"`
	OnlineNum      int    `json:"onlineNum"`
	Type           int    `json:"type"`
	LiveType       int    `json:"liveType"`
	Review         bool   `json:"review"`
	NeedForward    bool   `json:"needForward"`
	SystemMsg      string `json:"systemMsg"`
	MsgFilePath    string `json:"msgFilePath"`    // important
	PlayStreamPath string `json:"playStreamPath"` // important
	User           struct {
		UserId     string `json:"userId"`
		UserName   string `json:"userName"`
		UserAvatar string `json:"userAvatar"`
		Level      int    `json:"level"`
	} `json:"user"`
	TopUser            []interface{} `json:"topUser"`
	Mute               bool          `json:"mute"`
	LiveMode           int           `json:"liveMode"`
	PictureOrientation int           `json:"pictureOrientation"`
	IsCollection       int           `json:"isCollection"`
	MergeStreamUrl     string        `json:"mergeStreamUrl"`
	Crm                string        `json:"crm"`
	CoverPath          string        `json:"coverPath"`
	Title              string        `json:"title"`
	Ctime              string        `json:"ctime"`
	Announcement       string        `json:"announcement"`
	SpecialBadge       []interface{} `json:"specialBadge"`
}

type VoiceUser struct {
	UserId      int    `json:"userId"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	PfUrl       string `json:"pfUrl"`
	VoiceStatus bool   `json:"voiceStatus"`
}

type VoiceStatusContent struct {
	VoiceUserList []VoiceUser `json:"voiceUserList"` // empty for not doing
	StreamUrl     string      `json:"streamUrl"`     // empty for not doing
}

type MessageItem struct {
	MsgIDServer string `json:"msgIdServer"`
	MsgIDClient string `json:"msgIdClient"`
	MsgTime     int64  `json:"msgTime"`
	MsgType     string `json:"msgType"`
	Bodys       string `json:"bodys"` // the content of Bodys depends on MsgType

	// ExtInfo will be unmarshalled by pocketgo before returning to caller
	// so there is no need to access ExtInfoStr
	ExtInfoStr string `json:"extInfo"`
	ExtInfo    MessageExtInfo
}

const (
	// text only
	MessageTypeText = "TEXT"

	// text reply
	MessageTypeReply     = "REPLY"
	MessageTypeGiftReply = "GIFTREPLY"
	MessageTypeFlipCard  = "FLIPCARD" // 文字翻牌（公开）

	// media
	MessageTypeLivePush     = "LIVEPUSH"
	MessageTypeImage        = "IMAGE"
	MessageTypeExpressImage = "EXPRESSIMAGE"
	MessageTypeAudio        = "AUDIO"
	MessageTypeVideo        = "VIDEO"

	// todo 语音翻牌、视频翻牌
)

type MessageFormatter[T any] struct {
	MessageItem
}

func (f MessageFormatter[T]) GetTypedBody() (T, error) {
	var body T
	err := json.Unmarshal([]byte(f.MessageItem.Bodys), &body)
	return body, err
}

type ReplyBody struct {
	ReplyInfo struct {
		ReplyText      string `json:"replyText"`
		ReplyName      string `json:"replyName"`
		ReplyMessageId string `json:"replyMessageId"`
		Text           string `json:"text"`
	} `json:"replyInfo"`
	MessageType string `json:"messageType"`
}

type GiftReplyBody struct {
	GiftReplyInfo struct {
		ReplyText      string `json:"replyText"`
		ReplyName      string `json:"replyName"`
		ReplyMessageId string `json:"replyMessageId"`
		Text           string `json:"text"`
	} `json:"giftReplyInfo"`
	MessageType string `json:"messageType"`
}

type FlipCardBody struct {
	FlipCardInfo struct {
		Question   string `json:"question"`
		Answer     string `json:"answer"`
		QuestionId string `json:"questionId"`
		AnswerId   string `json:"answerId"`
		AnswerType string `json:"answerType"`
	} `json:"filpCardInfo"` // 官方typo
	MessageType string `json:"messageType"`
}

type LivePushBody struct {
	LivePushInfo struct {
		LiveCover string `json:"liveCover"`
		LiveTitle string `json:"liveTitle"`
		LiveId    string `json:"liveId"`
		ShortPath string `json:"shortPath"`
	} `json:"livePushInfo"`
	MessageType string `json:"messageType"`
}

type ImageBody struct {
	Size int    `json:"size"`
	Ext  string `json:"ext"`
	W    int    `json:"w"`
	Url  string `json:"url"` // might have no extension!
	Md5  string `json:"md5"`
	H    int    `json:"h"`
}

type ExpressImageBody struct {
	ExpressImgInfo struct {
		EmotionRemote string `json:"emotionRemote"`
		Width         int    `json:"width"`
		Height        int    `json:"height"`
	} `json:"expressImgInfo"`
	MessageType string `json:"messageType"`
}

type AudioBody struct {
	Size int    `json:"size"`
	Ext  string `json:"ext"`
	Dur  int    `json:"dur"`
	Url  string `json:"url"`
	Md5  string `json:"md5"`
}

type VideoBody struct {
	Url  string `json:"url"`
	Md5  string `json:"md5"`
	Ext  string `json:"ext"`
	H    int    `json:"h"`
	Size int    `json:"size"`
	W    int    `json:"w"`
	Dur  int    `json:"dur"` // milliseconds
}

func (m *MessageItem) fillExtInfo() {
	var extInfo MessageExtInfo
	if err := json.Unmarshal([]byte(m.ExtInfoStr), &extInfo); err != nil {
		log.Println(err)
	}
	m.ExtInfo = extInfo
}

type MessageExtInfo struct {
	Module      string `json:"module"`
	ChannelRole string `json:"channelRole"`
	User        struct {
		UserId   int    `json:"userId"`
		NickName string `json:"nickName"`
		TeamLogo string `json:"teamLogo"`
		Avatar   string `json:"avatar"`
		Level    int    `json:"level"`
		RoleId   int    `json:"roleId"`
		Vip      bool   `json:"vip"`
		PfUrl    string `json:"pfUrl"`
	} `json:"user"`
	BubbleId string `json:"bubbleId"`
}

type MessageContent struct {
	Message  []MessageItem `json:"message"`
	NextTime int64         `json:"nextTime"` // unix millisecond
}

type TpBalanceContent struct {
	TpNum int `json:"tpNum"`
}

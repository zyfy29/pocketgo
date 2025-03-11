package pocket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatReply(t *testing.T) {
	body, err := MessageFormatter[ReplyBody]{MessageItem{
		MsgType: MessageTypeReply,
		Bodys:   "{\"replyInfo\":{\"replyText\":\"三倍速看完了\",\"replyName\":\"小天线💤\",\"replyMessageId\":\"8a96c48f-d468-4153-ba95-119878b583c3\",\"text\":\"万万不可，这是大忌\"},\"messageType\":\"REPLY\"}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

func TestFormatGiftReply(t *testing.T) {
	body, err := MessageFormatter[GiftReplyBody]{MessageItem{
		MsgType: MessageTypeGiftReply,
		Bodys:   "{\"giftReplyInfo\":{\"replyText\":\"送给 许雅兰 1个秋天的奶茶\",\"replyName\":\"河为之Aix（三月哥）\",\"replyMessageId\":\"e33138150bbe4cf6985cd07aea78469b\",\"text\":\"谢谢小河的奶茶\"},\"messageType\":\"GIFTREPLY\"}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

func TestFormatImage(t *testing.T) {
	body, err := MessageFormatter[ImageBody]{MessageItem{
		MsgType: MessageTypeImage,
		Bodys:   "{\"size\":1161084,\"ext\":\"jpg\",\"w\":3024,\"url\":\"https://nim-nosdn.netease.im/NDA5MzEwOA==/bmltYV8xMTMzMjAyMTU3MjRfMTcyMzA4NTgzNDUxNF9iNDcxMDM5MC03NTdkLTRiZDYtYjA2Yy03ZTJhZjI0NzA0YTU=\",\"md5\":\"56ce042a7962f0761f4ec73e773c005e\",\"h\":4032}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

func TestFormatAudio(t *testing.T) {
	body, err := MessageFormatter[AudioBody]{MessageItem{
		MsgType: MessageTypeAudio,
		Bodys:   "{\"size\":22272,\"ext\":\"aac\",\"dur\":6176,\"url\":\"https://nim-nosdn.netease.im/NDA5MzEwOA==/bmltYV83MjUxOTY2MTExXzE3NDE1MDE4MjYxMzFfMzljMDAzMzktM2I0My00MWM2LWIyOTctYTg3ZGZkMzI0ZjM5\",\"md5\":\"6ee6c45d56df4f1287ca17df31824e85\"}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

func TestFormatVideo(t *testing.T) {
	body, err := MessageFormatter[VideoBody]{MessageItem{
		MsgType: MessageTypeVideo,
		Bodys:   "{\"url\":\"https://nim-nosdn.netease.im/NDA5MzEwOA==/bmltYV8xMTMzMjAyMTU3MjRfMTcyMzA4NTgzNDUxNF9jMWEzZDJlZi05MzFkLTRkMDUtOWVlZC1kZGU4MDAyZDliMDI=\",\"md5\":\"56fd84ceae802afef8c823996c4dbb26\",\"ext\":\"mp4\",\"h\":480,\"size\":430908,\"w\":248,\"dur\":6012}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

func TestFormatFlipCard(t *testing.T) {
	body, err := MessageFormatter[FlipCardBody]{MessageItem{
		MsgType: MessageTypeFlipCard,
		Bodys:   "{\"filpCardInfo\":{\"question\":\"hello小羊我是安熠 前几天看到你玩瓦了哈哈哈哈哈 因为我也玩所以看到你玩还是很开心的呢 你是颜值和实力并存的主播 夸奖你一下\",\"answer\":\"哎呀，别说了，我可是太紧张了，跟大家一起玩还是很开心的，一开始大家邀请我，我还是非常不可思议，我就是想当一个凑数的而已啦，但是我觉得跟大家一起完了之后特别特别有意思，希望下次还能玩，哈哈哈哈，就是去增加一些效果的你看得开心就行了 哈哈哈哈好的好的并存并存 真不错！\",\"questionId\":\"1102930794879193088\",\"answerId\":\"5311114\",\"answerType\":\"1\"},\"messageType\":\"FLIPCARD\"}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

func TestFormatLivePush(t *testing.T) {
	body, err := MessageFormatter[LivePushBody]{MessageItem{
		MsgType: MessageTypeLivePush,
		Bodys:   "{\"livePushInfo\":{\"liveCover\":\"/2025/0309/538697xvgr47xr6qrsic9xya706t2s.jpg\",\"liveTitle\":\"来啦\",\"liveId\":\"1105632350841933824\",\"shortPath\":\"live/playdetail?id=1105632350841933824\"},\"messageType\":\"LIVEPUSH\"}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

func TestFormatExpressImage(t *testing.T) {
	body, err := MessageFormatter[ExpressImageBody]{MessageItem{
		MsgType: MessageTypeExpressImage,
		Bodys:   "{\"expressImgInfo\":{\"emotionRemote\":\"https://source.48.cn/2024/1028/000x7mg5j7at557eg60qvee8k6k.png\",\"width\":200,\"height\":200},\"messageType\":\"EXPRESSIMAGE\"}",
	}}.GetTypedBody()
	assert.NoError(t, err)
	t.Log(body)
}

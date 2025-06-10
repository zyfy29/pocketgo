package pocket

func (c *Client) ValidateToken() (string, error) {
	content, err := responseFormatter[UserInfoItem]{c}.doRestWithResult("", "https://pocketapi.48.cn/user/api/v1/user/info/reload", map[string]string{})
	if err != nil {
		return "", err
	}
	return content.Nickname, nil
}

func (c *Client) GetHistoryLives(memberID string, nextTime int64) ([]LiveItem, int64, error) {
	content, err := responseFormatter[LiveListContent]{c}.doRestWithResult("", "https://pocketapi.48.cn/im/api/v1/chatroom/msg/list/aim/type", newLiveListReq(memberID, nextTime))
	if err != nil {
		return nil, 0, err
	}
	var lives []LiveItem
	for _, item := range content.Message {
		lives = append(lives, item.FormatToLiveItem())
	}
	return lives, content.NextTime, nil
}

func (c *Client) GetLiveInfo(liveId string) (LiveInfoContent, error) {
	content, err := responseFormatter[LiveInfoContent]{c}.doRestWithResult("", "https://pocketapi.48.cn/live/api/v1/live/getLiveOne", map[string]string{
		"liveId": liveId,
	})
	if err != nil {
		return LiveInfoContent{}, err
	}
	return content, nil
}

func (c *Client) GetVoiceStatus(channelId, serverId int) (VoiceStatusContent, error) {
	content, err := responseFormatter[VoiceStatusContent]{c}.doRestWithResult("", "https://pocketapi.48.cn/im/api/v1/team/voice/operate", map[string]int{
		"channelId":   channelId,
		"serverId":    serverId,
		"operateCode": 2,
	})
	if err != nil {
		return VoiceStatusContent{}, err
	}
	return content, nil
}

func (c *Client) GetMessageList(channelId, serverId int, nextTime int64) ([]MessageItem, int64, error) {
	content, err := responseFormatter[MessageContent]{c}.doRestWithResult("", "https://pocketapi.48.cn/im/api/v1/team/message/list/homeowner", map[string]any{
		"channelId": channelId,
		"serverId":  serverId,
		"nextTime":  nextTime,
		"limit":     100,
	})
	if err != nil {
		return nil, 0, err
	}
	res := []MessageItem{}
	for _, msg := range content.Message {
		msg.fillExtInfo()
		res = append(res, msg)
	}
	return res, content.NextTime, nil
}

func (c *Client) GetTpBalance(ticketId int) (int, error) {
	content, err := responseFormatter[TpBalanceContent]{c}.doRestWithResult("", "https://pocketapi.48.cn/netface/api/v1/user/tp/balance", map[string]int{
		"ticketId": ticketId,
	})
	if err != nil {
		return 0, err
	}
	return content.TpNum, nil
}

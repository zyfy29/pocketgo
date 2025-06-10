package pocket

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func must2[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

var testClient = NewClient(os.Getenv("POCKET_TOKEN"), 1000, nil)

func TestValidateToken(t *testing.T) {
	tests := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nihao",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.ValidateToken()
			if !tt.wantErr(t, err, fmt.Sprintf("ValidateToken()")) {
				return
			}
			t.Log(got)
		})
	}
}

func TestGetHistoryLives(t *testing.T) {
	type args struct {
		memberID string
		nextTime int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				memberID: "63566",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := testClient.GetHistoryLives(tt.args.memberID, tt.args.nextTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GetHistoryLives(%v, %v)", tt.args.memberID, tt.args.nextTime)) {
				return
			}
			for i, v := range got {
				t.Log(i, JsonMarshal(v))
			}
			t.Log(got1)
		})
	}
}

func TestClient_GetLiveInfo(t *testing.T) {
	type args struct {
		liveId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "normal",
			args:    args{liveId: "1037118894660980736"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.GetLiveInfo(tt.args.liveId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetLiveInfo(%v)", tt.args.liveId)) {
				return
			}
			t.Log(JsonMarshal(got))
		})
	}
}

func TestClient_GetVoiceStatus(t *testing.T) {
	type args struct {
		channelId int
		serverId  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				channelId: 1279492,
				serverId:  1246459,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.GetVoiceStatus(tt.args.channelId, tt.args.serverId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetLiveInfo()")) {
				return
			}
			t.Log(JsonMarshal(got))
		})
	}
}

func TestClient_GetMessageList(t *testing.T) {
	type args struct {
		channelId int
		serverId  int
		nextTime  int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				channelId: 1279492,
				serverId:  1246459,
				nextTime:  0,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messages, nextTime, err := testClient.GetMessageList(tt.args.channelId, tt.args.serverId, tt.args.nextTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GetLiveInfo()")) {
				return
			}
			for i, msg := range messages {
				t.Log(i, JsonMarshal(msg))
			}
			t.Log(nextTime)
		})
	}
}

func TestClient_GetTpBalance(t *testing.T) {
	type args struct {
		ticketId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				ticketId: 3,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.GetTpBalance(tt.args.ticketId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetTpBalance(%v)", tt.args.ticketId)) {
				return
			}
			t.Log(got)
		})
	}
}

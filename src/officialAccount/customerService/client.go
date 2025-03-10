package customerService

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/customerService/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/customerService/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取客服基本信息
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (comp *Client) List(ctx context.Context) (*response.ResponseList, error) {
	result := &response.ResponseList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/customservice/getkflist", nil, nil, &result)

	return result, err
}

// 获取在线客服基本信息
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (comp *Client) Online(ctx context.Context) (*response.ResponseKFOnlineList, error) {
	result := &response.ResponseKFOnlineList{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/customservice/getonlinekflist", nil, nil, &result)

	return result, err
}

// 添加客服帐号
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (comp *Client) Create(ctx context.Context, account string, nickname string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"kf_account": account,
		"nickname":   nickname,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "customservice/kfaccount/add", params, nil, nil, &result)

	return result, err
}

// 邀请绑定客服帐号
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (comp *Client) Update(ctx context.Context, account string, nickname string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"kf_account": account,
		"nickname":   nickname,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "customservice/kfaccount/update", params, nil, nil, &result)

	return result, err
}

// 删除客服帐号
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (comp *Client) Delete(ctx context.Context, account string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	query := &object.StringMap{
		"kf_account": account,
	}

	_, err := comp.BaseClient.HttpGet(ctx, "customservice/kfaccount/delete", query, nil, &result)

	return result, err
}

// 邀请绑定客服帐号
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (comp *Client) Invite(ctx context.Context, account string, wechatID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"kf_account": account,
		"invite_wx":  wechatID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "customservice/kfaccount/inviteworker", params, nil, nil, &result)

	return result, err
}

// 上传客服头像
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
func (comp *Client) SetAvatar(ctx context.Context, account string, path string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	} else {
		return nil, errors.New("path is empty")
	}

	params := &object.StringMap{
		"kf_account": account,
	}
	_, err := comp.BaseClient.HttpUpload(ctx, "customservice/kfaccount/uploadheadimg", files, nil, params, nil, &result)

	return result, err
}

func (comp *Client) Message(ctx context.Context, message contract.MessageInterface) *Messenger {
	messageBuilder := NewMessenger(comp)
	return messageBuilder.SetMessage(&message)
}

// 客服接口 - 发消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
func (comp *Client) Send(ctx context.Context, message interface{}) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/custom/send", message, nil, nil, result)

	return result, err
}

// 显示收入状态给用户
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
func (comp *Client) ShowTypingStatusToUser(ctx context.Context, openID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"touser":  openID,
		"command": "Typing",
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/custom/typing", params, nil, nil, &result)

	return result, err
}

// 隐藏收入状态给用户
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
func (comp *Client) HideTypingStatusToUser(ctx context.Context, openID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"touser":  openID,
		"command": "CancelTyping",
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/custom/typing", params, nil, nil, &result)

	return result, err
}

// 获取聊天记录
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Obtain_chat_transcript.html
func (comp *Client) Messages(ctx context.Context, data *request.RequestMessages) (*response.ResponseMessages, error) {
	result := &response.ResponseMessages{}

	//params, err := object.StructToHashMapWithTag(data, "json")
	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}
	_, err = comp.BaseClient.HttpPostJson(ctx, "customservice/msgrecord/getmsglist", params, nil, nil, &result)

	return result, err
}

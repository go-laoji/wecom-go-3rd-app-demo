package logic

import (
	"context"
	"errors"
	"github.com/go-laoji/wxbizmsgcrypt"

	"wecom-go-3rd-app-demo/internal/svc"
	"wecom-go-3rd-app-demo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SuiteCmdCallbackPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuiteCmdCallbackPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) SuiteCmdCallbackPostLogic {
	return SuiteCmdCallbackPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SuiteCmdCallbackPostLogic) SuiteCmdCallbackPost(req types.MsgRequest, body []byte) (*types.MsgResponse, error) {
	// TODO: 注意post请求接收数据时配置为suiteid
	wxbiz := wxbizmsgcrypt.NewWXBizMsgCrypt(l.svcCtx.Config.SuiteToken,
		l.svcCtx.Config.SuiteEncodingAesKey,
		l.svcCtx.Config.SuiteId,
		wxbizmsgcrypt.XmlType)
	echoStr, cryptErr := wxbiz.DecryptMsg(req.MsgSignature, req.TimeStamp, req.Nonce, body)
	if cryptErr != nil {
		return &types.MsgResponse{}, errors.New(cryptErr.ErrMsg)
	}
	l.Info(string(echoStr))
	return &types.MsgResponse{ErrCode: 0, ErrMsg: "ok", Data: []byte("success")}, nil
}

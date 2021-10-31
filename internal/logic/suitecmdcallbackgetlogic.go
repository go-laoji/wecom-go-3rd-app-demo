package logic

import (
	"context"
	"errors"
	"fmt"

	"suite-zero-svr/internal/svc"
	"suite-zero-svr/internal/types"

	"github.com/go-laoji/wxbizmsgcrypt"
	"github.com/tal-tech/go-zero/core/logx"
)

type SuiteCmdCallbackGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuiteCmdCallbackGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) SuiteCmdCallbackGetLogic {
	return SuiteCmdCallbackGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SuiteCmdCallbackGetLogic) SuiteCmdCallbackGet(req types.MsgRequest) (*types.MsgResponse, error) {
	// TODO: 注意get请求接收数据时配置为corpid
	wxbiz := wxbizmsgcrypt.NewWXBizMsgCrypt(l.svcCtx.Config.SuiteToken,
		l.svcCtx.Config.SuiteEncodingAesKey,
		l.svcCtx.Config.CorpId,
		wxbizmsgcrypt.XmlType)
	echoStr, err := wxbiz.VerifyURL(req.MsgSignature, req.TimeStamp, req.Nonce, req.EchoStr)
	if err != nil {
		return &types.MsgResponse{}, errors.New(fmt.Sprintf("%s", err.ErrMsg))
	}
	return &types.MsgResponse{ErrCode: 0, ErrMsg: "ok", Data: echoStr}, nil
}

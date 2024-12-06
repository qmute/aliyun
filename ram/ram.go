package ram

import (
	"github.com/denverdino/aliyungo/sts"

	"github.com/qmute/aliyun"
)

type Client struct {
	stsClt *sts.STSClient
}

func New(conf aliyun.Config) *Client {
	return &Client{
		stsClt: sts.NewClient(conf.Key, conf.Secret),
	}
}

// AssumeRole 获取扮演角色的临时身份凭证
func (p *Client) AssumeRole(opts ...AssumeRoleOpt) (AssumeRoleResponse, error) {
	req := &assumeRoleReq{}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return AssumeRoleResponse{}, err
		}
	}

	rsp, err := p.stsClt.AssumeRole(req.ToStsAssumeRoleRequest())
	if err != nil {
		return AssumeRoleResponse{}, err
	}

	return AssumeRoleResponse{
		RequestId: rsp.RequestId,
		AssumedRoleUser: AssumedRoleUser{
			AssumedRoleId: rsp.AssumedRoleUser.AssumedRoleId,
			Arn:           rsp.AssumedRoleUser.Arn,
		},
		Credentials: AssumedRoleUserCredentials{
			AccessKeySecret: rsp.Credentials.AccessKeySecret,
			AccessKeyId:     rsp.Credentials.AccessKeyId,
			Expiration:      rsp.Credentials.Expiration,
			SecurityToken:   rsp.Credentials.SecurityToken,
		},
	}, nil
}

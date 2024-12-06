package ram

import (
	"errors"
	"strings"

	"github.com/denverdino/aliyungo/sts"
)

// assumeRoleReq 获取扮演角色的临时身份凭证请求参数
// doc:https://help.aliyun.com/document_detail/371864.html?spm=a2c4g.11186623.0.0.5feb73bdkPhzgc
type assumeRoleReq struct {
	// 要扮演的RAM角色ARN。
	// 该角色是可信实体为阿里云账号类型的RAM角色
	RoleArn string
	// 角色会话名称
	RoleSessionName string
	// 过期时间。单位：秒。
	// 过期时间最小值为900秒，最大值为要扮演角色的MaxSessionDuration时间。默认值为3600秒
	DurationSeconds int
	// 为STS Token额外添加的一个权限策略，进一步限制STS Token的权限。具体如下：
	// 如果指定该权限策略，则STS Token最终的权限策略取RAM角色权限策略与该权限策略的交集。
	// 如果不指定该权限策略，则STS Token最终的权限策略取RAM角色的权限策略。
	// 长度为1~1024个字符
	Policy string
}

func (p *assumeRoleReq) ToStsAssumeRoleRequest() sts.AssumeRoleRequest {
	return sts.AssumeRoleRequest{
		RoleArn:         p.RoleArn,
		RoleSessionName: p.RoleSessionName,
		DurationSeconds: p.DurationSeconds,
		Policy:          p.Policy,
	}
}

type AssumeRoleOpt func(opt *assumeRoleReq) error

// WithAssumeRoleArn 要扮演的RAM角色ARN。
// 该角色是可信实体为阿里云账号类型的RAM角色
func WithAssumeRoleArn(roleArn string) AssumeRoleOpt {
	return func(req *assumeRoleReq) error {
		arn := strings.TrimSpace(roleArn)
		if roleArn == "" {
			return errors.New("roleArn must be not empty")
		}

		req.RoleArn = arn

		return nil
	}
}

// WithAssumeRoleSessionName 角色会话名称
func WithAssumeRoleSessionName(sessionName string) AssumeRoleOpt {
	return func(req *assumeRoleReq) error {
		name := strings.TrimSpace(sessionName)
		if name == "" {
			return errors.New("roleSessionName must be not empty")
		}

		req.RoleSessionName = name

		return nil
	}
}

// WithAssumeRoleDurationSeconds
// 过期时间。单位：秒。
// 过期时间最小值为900秒，最大值为要扮演角色的MaxSessionDuration时间。默认值为3600秒
func WithAssumeRoleDurationSeconds(seconds int) AssumeRoleOpt {
	return func(req *assumeRoleReq) error {
		if seconds < 900 {
			return errors.New("过期时间最小值为900秒")
		}

		req.DurationSeconds = seconds

		return nil
	}
}

// WithAssumeRolePolicy
// 为STS Token额外添加的一个权限策略，进一步限制STS Token的权限。具体如下：
// 如果指定该权限策略，则STS Token最终的权限策略取RAM角色权限策略与该权限策略的交集。
// 如果不指定该权限策略，则STS Token最终的权限策略取RAM角色的权限策略。
// 长度为1~1024个字符
func WithAssumeRolePolicy(policy string) AssumeRoleOpt {
	return func(req *assumeRoleReq) error {
		policy = strings.TrimSpace(policy)

		req.Policy = policy

		return nil
	}
}

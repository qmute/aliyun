package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"log"
	"strings"
	"time"

	"github.com/denverdino/aliyungo/ram"
	"github.com/denverdino/aliyungo/sts"
	"github.com/pkg/errors"
)

type WebToken struct {
	BaseUrl     string
	PolicyToken *PolicyToken
}
type PolicyToken struct {
	AccessKeyId string
	Host        string
	Expire      int64
	Signature   string
	Policy      string
	Directory   string
	Callback    string
	Filename    string // 上传文件名
	Download    string // 下载文件地址
	Path        string // 带文件名的上传路径

}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

// WebToken H5上传token
// @Deprecated 已废弃，不安全，使用WebTokenV2
func (p *Client) WebToken() (*WebToken, error) {
	now := time.Now().Unix()
	expire_end := now + p.opt.Expire
	var tokenExpire = p.fmtGmtIso8601(expire_end)

	dir := p.opt.Dir
	// create post policy json
	var conf ConfigStruct
	conf.Expiration = tokenExpire
	condition := []string{}
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, dir)
	conf.Conditions = append(conf.Conditions, condition)

	// calucate signature
	result, err := json.Marshal(conf)
	if err != nil {
		fmt.Println("callback json err:", err)
	}

	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(p.secret))
	if _, err := io.WriteString(h, debyte); err != nil {
		return nil, errors.WithStack(err)
	}
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var policyToken PolicyToken
	policyToken.AccessKeyId = p.key
	policyToken.Host = strings.TrimSpace(p.opt.Base)
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = dir
	policyToken.Policy = string(debyte)

	ret := &WebToken{
		BaseUrl:     p.objectBaseURL(p.opt.Dir),
		PolicyToken: &policyToken,
	}

	return ret, nil
}

func (p *Client) getPath(dir string, filename string) string {
	return fmt.Sprintf("%s%s", dir, filename)
}

// WebTokenV2 H5上传token
func (p *Client) WebTokenV2(dir string, originalName string, contentType string) (*WebToken, error) {
	now := time.Now().Unix()
	expire_end := now + p.opt.Expire
	var tokenExpire = p.fmtGmtIso8601(expire_end)

	log.Println("tokenExpire:", tokenExpire)

	if dir == "" {
		dir = p.opt.Dir
	}

	if !strings.HasSuffix(dir, "/") {
		dir = fmt.Sprintf("%s/", dir)
	}

	filename := p.GetFilename(originalName)

	path := p.getPath(dir, filename)

	// create post policy json
	var conf ConfigStruct
	conf.Expiration = tokenExpire
	condition := []string{}
	condition = append(condition, "eq")
	condition = append(condition, "$key")
	condition = append(condition, path)
	conf.Conditions = append(conf.Conditions, condition)

	condition = []string{}
	condition = append(condition, "starts-with")
	condition = append(condition, "$Content-Type")
	condition = append(condition, "multipart/form-data")
	conf.Conditions = append(conf.Conditions, condition)

	condition = []string{}
	condition = append(condition, "starts-with")
	condition = append(condition, "$Content-Type")

	if contentType == "" {
		contentType = p.guessContextType(originalName)
	}
	condition = append(condition, strings.Split(contentType, "/")[0]+"/")
	conf.Conditions = append(conf.Conditions, condition)

	// calucate signature
	result, err := json.Marshal(conf)
	if err != nil {
		fmt.Println("callback json err:", err)
	}

	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(p.secret))
	if _, err := io.WriteString(h, debyte); err != nil {
		return nil, errors.WithStack(err)
	}
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var policyToken PolicyToken
	policyToken.AccessKeyId = p.key
	policyToken.Host = strings.TrimSpace(p.opt.Base)
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = dir
	policyToken.Policy = string(debyte)
	policyToken.Filename = filename
	policyToken.Download = p.ObjectURL(dir, filename)
	policyToken.Path = path

	ret := &WebToken{
		BaseUrl:     p.objectBaseURL(dir),
		PolicyToken: &policyToken,
	}

	return ret, nil
}

type Credentials struct {
	AccessKeySecret string
	AccessKeyId     string
	Expiration      string
	SecurityToken   string
}

type StsToken struct {
	BaseUrl     string
	Bucket      string
	Dir         string
	Endpoint    string
	Callback    string
	Credentials *Credentials
}

func (p *Client) StsToken() (*StsToken, error) {
	roleRsp, err := p.assumeRole()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ret := &StsToken{
		BaseUrl:  p.objectBaseURL(p.opt.Dir),
		Bucket:   p.opt.Bucket,
		Dir:      p.opt.Dir,
		Endpoint: p.opt.Endpoint,
		Callback: "",
		Credentials: &Credentials{
			AccessKeySecret: roleRsp.Credentials.AccessKeySecret,
			AccessKeyId:     roleRsp.Credentials.AccessKeyId,
			Expiration:      roleRsp.Credentials.Expiration,
			SecurityToken:   roleRsp.Credentials.SecurityToken,
		},
	}

	return ret, nil
}

func (p *Client) assumeRole() (*sts.AssumeRoleResponse, error) {
	roleArn := "acs:ram::" + p.account + ":role/aliyunosstokengeneratorrole"
	req := p.createAssumeRoleRequest(roleArn)
	stsClient := sts.NewClient(p.key, p.secret)
	resp, err := stsClient.AssumeRole(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &resp, nil
}

func (p *Client) createAssumeRoleRequest(roleArn string) sts.AssumeRoleRequest {
	document, _ := json.Marshal(p.createPolicyDocument(p.opt.Bucket))
	return sts.AssumeRoleRequest{
		RoleArn:         roleArn,
		RoleSessionName: "AliyunOSSTokenGeneratorRole",
		DurationSeconds: 900, // aliyun 要求最小15分钟
		Policy:          string(document),
	}
}

func (p *Client) createPolicyDocument(bucket string) ram.PolicyDocument {
	resource := "acs:oss:*:*:" + bucket + "/*"
	policyDocument := ram.PolicyDocument{
		Statement: []ram.PolicyItem{
			{
				Action:   "oss:GetObject",
				Effect:   "Allow",
				Resource: resource,
			},
			{
				Action:   "oss:PutObject",
				Effect:   "Allow",
				Resource: resource,
			},
			{
				Action:   "oss:ListObjects",
				Effect:   "Allow",
				Resource: resource,
			},
			{
				Action:   "oss:ListParts",
				Effect:   "Allow",
				Resource: resource,
			},
			{
				Action:   "oss:AbortMultipartUpload",
				Effect:   "Allow",
				Resource: resource,
			},
		},
		Version: "1",
	}
	return policyDocument
}

func (p *Client) fmtGmtIso8601(expire_end int64) string {
	var tokenExpire = time.Unix(expire_end, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

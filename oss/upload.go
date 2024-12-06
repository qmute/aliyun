package oss

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/denverdino/aliyungo/oss"
	"github.com/google/uuid"
)

const (
	defaultContentType = "application/octet-stream"
)

type ACL string

const (
	Private           = ACL("private")
	PublicRead        = ACL("public-read")
	PublicReadWrite   = ACL("public-read-write")
	AuthenticatedRead = ACL("authenticated-read")
	BucketOwnerRead   = ACL("bucket-owner-read")
	BucketOwnerFull   = ACL("bucket-owner-full-control")
)

type UploadInfo struct {
	Payload      io.Reader
	OriginalName string // 包含后缀
	ContentType  string
	Dir          string // 上传目录， 空时，是默认目录
	Size         int64  // 文件大小
	Acl          ACL    // 权限
}

func (p *UploadInfo) Valid() error {
	if p.OriginalName == "" {
		return ErrEmptyOriginalName
	}

	return nil
}

type FileInfo struct {
	Filename    string
	Size        int64
	DownloadUrl string
	ContentType string
}

func (p *Client) Upload(ctx context.Context, info *UploadInfo) (*FileInfo, error) {
	if err := info.Valid(); err != nil {
		return nil, err
	}

	if info.Acl == "" {
		info.Acl = Private
	}

	filename, err := p.getUploadFilename(info.OriginalName)
	if err != nil {
		return nil, err
	}

	opt := p.buildOptions(info)

	contentType := info.ContentType
	if contentType == "" {
		contentType = p.guessContextType(info.OriginalName)
	}

	dir := p.opt.Dir
	if info.Dir != "" {
		dir = fmt.Sprintf("%s%s", dir, info.Dir)
	}

	path := dir + filename

	err = p.bucket.PutReader(path, info.Payload, info.Size, contentType, oss.ACL(info.Acl), opt)
	if err != nil {
		return nil, err
	}

	ossFile := &FileInfo{
		Filename:    filename,
		ContentType: contentType,
		Size:        info.Size,
		DownloadUrl: p.ObjectURL(dir, filename),
	}
	return ossFile, nil

}

func (p *Client) GenerateObjectKey() (string, error) {
	UUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return UUID.String(), nil
}

func (p *Client) ObjectURL(dir, filename string) string {
	return fmt.Sprintf("%v%v", p.objectBaseURL(dir), filename)
}

// return http(s)://f.xxx.com/ || http(s)://f.xxx.com/mt/
func (p *Client) objectBaseURL(dir string) string {
	if !strings.HasPrefix(dir, "/") {
		dir = "/" + dir
	}
	return fmt.Sprintf("%v%v", p.opt.Base, dir)
}

func (p *Client) getUploadFilename(originalName string) (string, error) {
	filename, err := p.GenerateObjectKey()
	if err != nil {
		return "", err
	}

	return filename + p.getSuffix(originalName), nil
}

func (p *Client) GetFilename(originalName string) string {
	name, err := p.getUploadFilename(originalName)
	if err != nil {
		panic(err)
	}

	return name
}

func (p *Client) getSuffix(originalName string) string {
	idx := strings.LastIndex(originalName, ".")
	if idx == -1 {
		return ""
	}

	suffix := originalName[idx:]

	return strings.ToLower(suffix)
}

func (p *Client) guessContextType(originalName string) string {
	contentType := defaultContentType

	suffix := p.getSuffix(originalName)

	switch suffix {
	case ".txt":
		contentType = "text/plain"
	case ".gif":
		contentType = "image/gif"
	case ".png":
		contentType = "image/png"
	case ".jpeg", ".jpg":
		contentType = "image/jpeg"
	case ".pdf":
		contentType = "application/pdf"
	case ".doc", ".docx":
		contentType = "application/msword"
	case ".xls", ".xlsx":
		contentType = "application/vnd.ms-excel"
	case ".csv":
		contentType = "text/csv"
	case ".mp4":
		contentType = "video/mp4"
	case ".mp3":
		contentType = "audio/mpeg"
	}

	return contentType
}

func (p *Client) buildOptions(up *UploadInfo) oss.Options {
	opt := oss.Options{}
	if up.OriginalName != "" {
		opt.ContentDisposition = fmt.Sprintf("filename=%s", up.OriginalName)
	}

	return opt
}

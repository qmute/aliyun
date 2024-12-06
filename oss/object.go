package oss

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ImageInfo struct {
	FileSize    Value `json:"FileSize"`
	Format      Value `json:"Format"`
	ImageHeight Value `json:"ImageHeight"`
	ImageWidth  Value `json:"ImageWidth"`
}

func (p *ImageInfo) GetFileSize() int {
	if p.FileSize.Value == "" {
		return 0
	}

	return p.toInt(p.FileSize.Value)
}

func (p *ImageInfo) GetImageHeight() int {
	if p.ImageHeight.Value == "" {
		return 0
	}

	return p.toInt(p.ImageHeight.Value)
}

func (p *ImageInfo) GetImageWidth() int {
	if p.ImageWidth.Value == "" {
		return 0
	}

	return p.toInt(p.ImageWidth.Value)
}

func (p *ImageInfo) toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type Value struct {
	Value string `json:"value"`
}

// GetImageInfo 获取图片信息
// https://cdn.qian.fm/t/upload/manage/20210820/3a7c040e-019b-11ec-8aee-00163e106f77.jpeg?x-oss-process=image/info
func GetImageInfo(url string) (ImageInfo, error) {
	info := ImageInfo{}

	wrapUrl := fmt.Sprintf("%s?x-oss-process=image/info", url)
	clt := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, wrapUrl, nil)
	if err != nil {
		return info, err
	}

	resp, err := clt.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return info, errors.New(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}

	if err = json.Unmarshal(b, &info); err != nil {
		return info, errors.WithStack(err)
	}

	return info, nil
}

package sts

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"time"
)

type Resp struct {
	AccessID string `json:"access_id"`
	Secret   string `json:"secret"`
	Token    string `json:"token"`
	Expire   string `json:"expire"`
}

func createClient(akID, akSecret, endpoint string) (*sts.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(akID),
		AccessKeySecret: tea.String(akSecret),
	}

	config.Endpoint = tea.String(endpoint)

	result, err := sts.NewClient(config)

	return result, err
}

func GenStsToken(akID, akSecret, endpoint, arn string, duration int64) (*Resp, error) {
	if duration < 900 || duration > 3600 {
		return nil, errors.New("duration must be between 900 and 3600")
	}

	client, err := createClient(akID, akSecret, endpoint)
	if err != nil {
		return nil, err
	}

	request := &sts.AssumeRoleRequest{
		DurationSeconds: tea.Int64(duration),
		RoleArn:         tea.String(arn),
		RoleSessionName: tea.String("gin.session." + time.Now().Format("20060102150405")),
	}

	resp, err := client.AssumeRoleWithOptions(request, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}

	cred := resp.Body.Credentials
	return &Resp{
		AccessID: tea.StringValue(cred.AccessKeyId),
		Secret:   tea.StringValue(cred.AccessKeySecret),
		Token:    tea.StringValue(cred.SecurityToken),
		Expire:   tea.StringValue(cred.Expiration),
	}, nil
}

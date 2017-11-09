package awsiot

import (
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"bytes"
	"net/url"
	"fmt"
	"time"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
)

type AwsIot struct {
	clientConfig client.Config
}

const serviceName = "iotdevicegateway"

func New(p client.ConfigProvider, cfgs ...*aws.Config) *AwsIot {
	return &AwsIot{
		clientConfig: p.ClientConfig(serviceName, cfgs...),
	}
}

func (a *AwsIot) WebsocketUrl(endpoint string) (string, error) {
	signer := v4.NewSigner(a.clientConfig.Config.Credentials)
	duration := time.Duration(5)
	body := bytes.NewReader([]byte{})

	originalUrl, err := url.Parse(fmt.Sprintf("wss://%s/mqtt", endpoint))
	if err != nil {
		return "", err
	}

	req := &http.Request{
		Method: "GET",
		URL:    originalUrl,
	}
	_, err = signer.Presign(req, body, a.clientConfig.SigningName, a.clientConfig.SigningRegion, duration, time.Now())

	if err != nil {
		return "", err
	}

	return req.URL.String(), nil
}

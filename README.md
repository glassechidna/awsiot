## Why

The AWS SDK for Go doesn't yet support bidirectional message sending for the AWS
IoT service. As such, I created a _tiny_ library to fill the gap (at least as far
as my needs extended) until AWS provides this support.

## How

This library's usage is modelled on how AWS API clients are used in the official
library. It should feel familiar to developers who have used those before.

```go
package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/glassechidna/awsiot"
)

func main() {
	sessOpts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}

	sess := session.Must(session.NewSessionWithOptions(sessOpts))
	iot := awsiot.New(sess)
	theUrl, _ := iot.WebsocketUrl("a1kxjqeyezkt7")

	opts := MQTT.NewClientOptions().AddBroker(theUrl)
	opts.SetClientID("clientid")
	client := MQTT.NewClient(opts)
}
``` 

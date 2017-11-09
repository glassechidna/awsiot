## Why

The AWS SDK for Go doesn't yet support bidirectional message sending for the AWS
IoT service - see [#706][gh706], [#820][gh820], [#1304][gh1304]. As such, I created 
a _tiny_ library to fill the gap (at least as far as my needs extended) until AWS 
provides this support.

[gh706]: https://github.com/aws/aws-sdk-go/issues/706
[gh820]: https://github.com/aws/aws-sdk-go/issues/820
[gh1304]: https://github.com/aws/aws-sdk-go/issues/1304

## How

This library's usage is modelled on how AWS API clients are used in the official
library. It should feel familiar to developers who have used those before. It 
returns a URL string, which can be used to initialise a client in the 
`eclipse/paho.mqtt.golang` MQTT library.

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

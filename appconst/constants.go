package appconst

import "fmt"

var (
	PUBSUB_NAME = "apppubsub"
	TOPIC_NAME  = "thetopic"
)

func MakeDAPRSidecarPubURL(pubsubName, topicName, daprPort string, isRawMessage bool) string {
	if isRawMessage {
		return fmt.Sprintf("http://localhost:%s/v1.0/publish/%s/%s?metadata.rawPayload=true",
			daprPort, PUBSUB_NAME, TOPIC_NAME)
	}
	return fmt.Sprintf("http://localhost:%s/v1.0/publish/%s/%s", daprPort, PUBSUB_NAME, TOPIC_NAME)
}

package marathon

import (
	"log"
	"time"
)

//TODO: we should have separate event structs for each message type
type Event struct {
	EventType string `json:"eventType"`
	Timestamp string `json:"timestamp"`
	AppId     string `json:"appId,omitempty"`
	TaskId    string `json:"taskId,omitempty"`
}

/*
API Request: Fired every time marathon receives an API request that modifies an app
Deployment: Fired whenever a deployment is started or stopped

type DeploymentEvent struct {
	eventType string `json:"eventType"`
	timestamp string `json:"timestamp"`
	groupId   string `json:"timestamp"`
	version   string `json:"version"`
}

*/

func (e *Event) Time() time.Time {
	t, err := time.Parse(time.RFC3339, e.Timestamp)
	if err != nil {
		log.Printf(err.Error())
		return time.Time{}
	}
	return t
}

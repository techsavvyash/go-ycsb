package stylusdb

import (
	"encoding/json"
	"fmt"
)

type Command struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type DataItem struct {
	Command Command `json:"command"`
}

type Payload struct {
	Task string     `json:"task"`
	Data []DataItem `json:"data"`
}

func givePayloadSkeleton(op, key, val string) Payload {
	return Payload{
		Task: op,
		Data: []DataItem{
			{
				Command: Command{
					Key: key,
					Val: val,
				},
			},
		},
	}
}

func GenerateEvents(length int) []Payload {
	var events []Payload
	for i := 1; i <= length; i++ {
		events = append(events, givePayloadSkeleton("SET", fmt.Sprintf("key_%d", i), fmt.Sprintf("%d", i)))
	}
	for i := 1; i <= length; i++ {
		events = append(events, givePayloadSkeleton("GET", fmt.Sprintf("key_%d", i), fmt.Sprintf("%d", i)))
	}
	return events
}

func ConvertToQuery(uuid string, payload Payload) string {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err) // Handle the error gracefully in your application
	}
	return fmt.Sprintf("%s|%s\n", uuid, string(jsonData))
}

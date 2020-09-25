package backend

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func parseJSON() (map[string]interface{}, error) {
	f, err := os.Open("./events.json")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	fileContents, _ := ioutil.ReadAll(f)
	var jsonContents map[string]interface{}
	json.Unmarshal([]byte(fileContents), &jsonContents)

	return jsonContents, nil
}

func GetEventsFromConfig() (map[string]*Event, error) {
	jsonContents, err := parseJSON()
	if err != nil {
		return nil, err
	}

	eventsJSON := jsonContents["events"].([]interface{})
	events := make(map[string]*Event)

	for _, item := range eventsJSON {
		v := item.(map[string]interface{})
		event := Event{v["id"].(string), v["name"].(string), v["description"].(string)}
		events[v["id"].(string)] = &event
	}

	return events, nil
}

func GetCommentsFromConfig() (map[string]*Comment, error) {
	jsonContents, err := parseJSON()
	if err != nil {
		return nil, err
	}

	commentsInterface := jsonContents["comments"].([]interface{})
	comments := make(map[string]*Comment)

	for _, item := range commentsInterface {
		v := item.(map[string]interface{})
		comment := Comment{v["id"].(string), v["commenter"].(string), v["content"].(string), v["event_id"].(string)}
		comments[v["id"].(string)] = &comment
	}

	return comments, nil
}

package slack

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/raybarrera/webservice/pkg/response"
)

type SlackEvent struct {
	Token     string       `json:"token"`
	TeamID    string       `json:"team_id"`
	APIAppID  string       `json:"api_app_id"`
	Event     SlackMessage `json:"event"`
	Type      string       `json:"type"`
	EventID   string       `json:"event_id"`
	EventTime int64        `json:"event_time"`
	Challenge string       `json:"challenge"`
}

type SlackMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	TS      string `json:"ts"`
}

func SlackVerificationHandler(writer http.ResponseWriter, request *http.Request) {
	var slackEvent SlackEvent
	err := json.NewDecoder(request.Body).Decode(&slackEvent)
	encoder := json.NewEncoder(writer)
	if err != nil {
		slog.Error(err.Error())
		response := response.Payload{
			Status: "ERROR",
			Code:   http.StatusBadRequest,
			Errors: []response.ApiError{
				{
					ReferenceCode: "",
					Message:       err.Error(),
				},
			},
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		_ = encoder.Encode(response)
		return
	}
	switch slackEvent.Type {
	case "url_verification":
		response := slackEvent.Challenge
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte(response))
		return
	default:
		slog.Info("Slack Message Received", "message", slackEvent)
		response := response.Payload{
			Status:   "OK",
			Code:     http.StatusOK,
			Messages: []string{"OK"},
		}
		writer.WriteHeader(http.StatusOK)
		encoder.Encode(response)
		return
	}
}

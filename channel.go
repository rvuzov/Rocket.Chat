package rocketchat

var Channel channel

type channel struct{}

type (
	apiPostMessageRequest struct {
		RoomId string `json:"roomId"`
		Text   string `json:"text"`
	}

	apiPostMessageResponse struct {
		apiResponse
	}
)

func (channel) PostMessage(roomId string, text string) (err error) {
	defer makeError(&err, "PostMessage", "roomId", roomId, "text", text)

	request := apiPostMessageRequest{
		RoomId: roomId,
		Text:   text,
	}

	var response apiPostMessageResponse
	err = apiPost(kapiChannelPostMessage, request, &response)
	return
}

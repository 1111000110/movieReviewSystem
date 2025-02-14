package hub

import "movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/types"

type Client interface {
	GetClientId() int64
	WritePump()
	ReadPump()
	GetSendBuffer() chan types.Message
}

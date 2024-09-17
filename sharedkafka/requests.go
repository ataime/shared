package sharedkafka

import (
	"time"

	"github.com/ataime/shared/shareddto"
)

// chat->admin 举报消息（明文）
type KafkaComplaintChatToCensorRequest struct {
	ComplainID string               `json:"complain_id"` // 举报人
	RespondID  string               `json:"respond_id"`  // 被举报人
	ChatMsgID  string               `json:"chat_msg_id"`
	Content    shareddto.OutChatMsg `json:"content"`
	ProduceAt  time.Time            `json:"produce_at"`
	ConsumeAt  *time.Time           `json:"consume_at"`
}

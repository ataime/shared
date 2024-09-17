package shareddto

import (
	"time"

	"github.com/gocql/gocql"
)

type OutChatMsg struct {
	ID             gocql.UUID    `json:"id"`
	RoomID         string        `json:"room_id"`
	RoomName       string        `json:"room_name"`
	Sender         ShortUserInfo `json:"sender"`
	MentionedUsers []string      `json:"mentioned_users"` // 被@的人
	Content        ChatContent   `json:"content"`
	SendAt         time.Time     `json:"send_at"`
	ReplyTo        *gocql.UUID   `json:"reply_to"` // 回复的消息ID
}

type ShortUserInfo struct {
	UserID       string `json:"user_id"`       // 用户ID
	AnonID       string `json:"anon_id"`       // 用户匿名ID标识（对外隐藏UserID，保护用户隐私）
	CustomID     string `json:"custom_id"`     // 用户自定义ID
	Nickname     string `json:"nickname"`      // 昵称
	Avatar       string `json:"avatar"`        // 头像
	Gender       string `json:"gender"`        // 性别（同 realname 一样限制逻辑，不允许修改）
	GenderHidden bool   `json:"gender_hidden"` // 性别隐藏
}

type ChatContent struct {
	Text     *TextContent     `json:"text,omitempty"`
	Image    *ImageContent    `json:"image,omitempty"`
	Video    *VideoContent    `json:"video,omitempty"`
	File     *FileContent     `json:"file,omitempty"`
	Link     *LinkContent     `json:"link,omitempty"`
	Card     *CardContent     `json:"card,omitempty"`
	Music    *MusicContent    `json:"music,omitempty"`
	Location *LocationContent `json:"location,omitempty"`
	Sticker  *StickerContent  `json:"sticker,omitempty"`
	Tip      *TipContent      `json:"tip,omitempty"`
}

type TextContent struct {
	Text string `json:"text"` // 文本内容
}
type ImageContent struct {
	Url       string `json:"url"`       // 图片URL
	Thumbnail string `json:"thumbnail"` // 缩略图URL
	Width     int    `json:"width"`     // 图片宽度
	Height    int    `json:"height"`    // 图片高度
}
type VideoContent struct {
	URL         string `json:"url"`         // 视频URL
	Thumbnail   string `json:"thumbnail"`   // 视频缩略图URL
	Title       string `json:"title"`       // 视频标题
	Description string `json:"description"` // 视频描述
	Duration    int    `json:"duration"`    // 视频时长（秒）
	Width       int    `json:"width"`       // 视频宽度
	Height      int    `json:"height"`      // 视频高度
}
type FileContent struct {
	FileName string `json:"file_name"` // 文件名
	FileSize int    `json:"file_size"` // 文件大小
	URL      string `json:"url"`       // 文件URL
}
type LinkContent struct {
	URL           string   `json:"url"`
	Title         string   `json:"title,omitempty"`         // 链接的标题
	Description   string   `json:"description,omitempty"`   // 链接的描述
	ImageURL      string   `json:"imageUrl,omitempty"`      // 链接相关的图片URL
	Author        string   `json:"author,omitempty"`        // 链接内容的作者
	PublishedDate string   `json:"publishedDate,omitempty"` // 链接内容的发布时间
	ContentType   string   `json:"contentType,omitempty"`   // e.g., "article", "video", "image", etc.
	Tags          []string `json:"tags,omitempty"`          // 链接的标签，可以用来区分业务
	Source        string   `json:"source,omitempty"`        // e.g., "website name"
}
type CardContent struct {
	URL          string   `json:"url,omitempty"`
	UserID       string   `json:"user_id"`      // 用户ID
	Avatar       string   `json:"avatar"`       // 头像
	Nickname     string   `json:"nickname"`     // 昵称
	Gender       string   `json:"gender"`       // 性别
	Name         string   `json:"name"`         // 名字
	Organization string   `json:"organization"` // 组织
	Tags         []string `json:"tags"`         // 标签
}
type MusicContent struct {
	Artist   string `json:"artist"`   // 艺术家
	Title    string `json:"title"`    // 标题
	Album    string `json:"album"`    // 专辑
	CoverURL string `json:"coverUrl"` // 封面URL
	MusicURL string `json:"musicUrl"` // 音乐URL
}
type LocationContent struct {
	Latitude  float64 `json:"latitude"`  // 纬度
	Longitude float64 `json:"longitude"` // 经度
	Name      string  `json:"name"`      // 位置名称
	Address   string  `json:"address"`   // 位置地址
}
type StickerContent struct {
	StickerID  string `json:"stickerId"`  // 表情ID
	StickerURL string `json:"stickerUrl"` // 表情URL
}

type TipContent struct {
	Msg string `json:"msg"`
}

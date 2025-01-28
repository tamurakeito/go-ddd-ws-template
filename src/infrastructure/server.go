package infrastructure

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Server構造体: WebSocketサーバーの主要なデータと操作を管理
type Server struct {
	repository *MessageRepository
	Upgrader   websocket.Upgrader // HTTPをWebSocket接続にアップグレードするための設定
}

// NewServer関数: 新しいServerインスタンスを作成して返す
func NewServer(repo *MessageRepository) *Server {
	return &Server{
		repository: repo,
		Upgrader: websocket.Upgrader{
			// CheckOrigin関数: オリジン制限を設定（セキュリティ対策として必要に応じて修正）
			CheckOrigin: func(r *http.Request) bool {
				return true // 現在はすべてのオリジンを許可
			},
		},
	}
}

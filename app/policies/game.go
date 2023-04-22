package policies

import (
	"competition-backend/app/http/models/game"
	"github.com/gin-gonic/gin"
)

// CanModifyGame 是否可以修改数据
// Status == 0 表示比赛未开始，则可以修改，反之
func CanModifyGame(c *gin.Context, _game game.Game) bool {
	return _game.Status == 0
}

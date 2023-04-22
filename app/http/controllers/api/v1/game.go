package v1

import (
	"competition-backend/app/http/models/game"
	"competition-backend/app/requests"
	"competition-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type GameAPIController struct {
	BaseAPIController
}

func (gc *GameAPIController) Create(c *gin.Context) {
	request := requests.GameRequest{}
	if ok := requests.Validate(c, &request, requests.CreateGame); !ok {
		return
	}
	// 创建成功返回一个 SFCode，用户需使用 SFCode 进入比赛
	gameModel := game.Game{
		Title:  request.Title,
		Desc:   request.Desc,
		Status: 0,
	}
	if gameModel.ID > 0 {
		response.Created(c, gameModel)
	} else {
		response.Abort500(c, "创建失败，请稍后重试")
	}
}

func (gc *GameAPIController) Index(c *gin.Context) {

}

func (gc *GameAPIController) Update(c *gin.Context) {

}

func (gc *GameAPIController) Delete(c *gin.Context) {
}

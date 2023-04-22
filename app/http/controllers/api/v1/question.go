package v1

import (
	"competition-backend/app/http/models/question"
	"competition-backend/app/requests"
	"competition-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type QuestionAPIController struct {
	BaseAPIController
}

func (qc *QuestionAPIController) Index(c *gin.Context) {

}

func (qc *QuestionAPIController) Create(c *gin.Context) {
	req := requests.QuestionRequest{}
	if ok := requests.Validate(c, &req, requests.CreateQuestion); !ok {
		return
	}
	qModel := question.Question{
		Title:  req.Title,
		Desc:   req.Desc,
		Answer: req.Answer,
		Score:  req.Score,
	}
	qModel.Create()
	if qModel.ID > 0 {
		response.Created(c, qModel)
	} else {
		response.Abort500(c, "创建失败，请稍后重试")
	}
}

func (qc *QuestionAPIController) Delete(c *gin.Context) {
	qModel := question.Get(c.Param("id"))
	if qModel.ID == 0 {
		response.Abort404(c)
		return
	}
	ra := qModel.Delete()
	if ra > 0 {
		response.Success(c)
		return
	}
	response.Abort500(c, "删除失败，请稍后重试")
}

func (qc *QuestionAPIController) Update(c *gin.Context) {
	qModel := question.Get(c.Param("id"))
	if qModel.ID == 0 {
		response.Abort404(c)
		return
	}
	req := requests.QuestionRequest{}
	if ok := requests.Validate(c, &req, requests.CreateQuestion); !ok {
		return
	}
	qModel.Title = req.Title
	qModel.Desc = req.Desc
	qModel.Answer = req.Answer
	qModel.Score = req.Score
	ra := qModel.Save()
	if ra > 0 {
		response.Data(c, qModel)
	} else {
		response.Abort500(c, "更新失败，请稍后重试")
	}
}

func (qc *QuestionAPIController) Show(c *gin.Context) {
	qModel := question.Get(c.Param("id"))
	if qModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, qModel)
}

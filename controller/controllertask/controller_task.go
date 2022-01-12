package controllertask

import (
	"net/http"
	"p03_kanban_board/helper"
	"p03_kanban_board/model/modeltask"
	"p03_kanban_board/service/servicetask"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerTask interface {
	Create(ctx *gin.Context)
	Gets(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	srv servicetask.ServiceTask
}

func (c *controller) Delete(ctx *gin.Context) {
	taskID, _ := strconv.ParseUint(ctx.Param("taskID"), 10, 64)
	UserID := ctx.MustGet("user_id").(uint)
	err := c.srv.Delete(uint(taskID), UserID)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, "Task has been deleted", nil))
}

func (c *controller) UpdateStatus(ctx *gin.Context) {
	request := new(modeltask.RequestUpdateStatus)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	request.ID, _ = strconv.ParseUint(ctx.Param("taskID"), 10, 64)
	request.UserID = ctx.MustGet("user_id").(uint)
	resp, err := c.srv.UpdateStatus(*request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, resp, nil))
	return
}

func (c *controller) UpdateCategory(ctx *gin.Context) {
	request := new(modeltask.RequestUpdateCategory)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	request.ID, _ = strconv.ParseUint(ctx.Param("taskID"), 10, 64)
	request.UserID = ctx.MustGet("user_id").(uint)
	resp, err := c.srv.UpdateCategory(*request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, resp, nil))
	return
}

func (c *controller) Update(ctx *gin.Context) {
	request := new(modeltask.RequestUpdate)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	request.ID, _ = strconv.ParseUint(ctx.Param("taskID"), 10, 64)
	request.UserID = ctx.MustGet("user_id").(uint)
	resp, err := c.srv.Update(*request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, resp, nil))
	return
}

func (c *controller) Gets(ctx *gin.Context) {
	response, err := c.srv.Gets()
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (c *controller) Create(ctx *gin.Context) {
	request := new(modeltask.Request)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	request.UserID = ctx.MustGet("user_id").(uint)

	response, err := c.srv.Create(*request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusCreated, response, nil))
}

func New(srv servicetask.ServiceTask) ControllerTask {
	return &controller{srv: srv}
}

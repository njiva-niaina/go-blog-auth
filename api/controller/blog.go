package controller

import (
	"blog/api/service"
	"blog/models"
	"blog/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service service.PostService
}

func NewPostController(s service.PostService) PostController {
	return PostController{
		service: s,
	}
}

func (p PostController) GetPosts(ctx *gin.Context) {
	var posts models.Post
	keyword := ctx.Query("keyword")
	data, total, err := p.service.FindAll(posts, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0)
	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Post result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		},
	})
}

func (p PostController) AddPost(ctx *gin.Context) {
	var post models.Post
	ctx.ShouldBind(&post)
	if post.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required ")
		return
	}
	if post.Body == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required ")
		return
	}
	err := p.service.Save(post)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create post")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Post")
}

func (p PostController) GetPost(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var post models.Post
	post.ID = id
	foundPost, err := p.service.Find(post)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Post")
		return
	}
	response := foundPost.ResponseMap()
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Post",
		Data:    &response,
	})
}

func (p PostController) DeletePost(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(id)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete post")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted successfully",
	}
	ctx.JSON(http.StatusOK, response)
}

func (p PostController) UpdatePost(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var post models.Post
	post.ID = id
	postRecord, err := p.service.Find(post)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Post with given id not found")
		return
	}
	ctx.ShouldBind(&postRecord)
	if postRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required ")
		return
	}
	if postRecord.Body == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required ")
		return
	}
	if err := p.service.Update(postRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Post")
		return
	}
	response := postRecord.ResponseMap()
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Post",
		Data:    response,
	})
}

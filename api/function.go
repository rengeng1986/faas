package restful

import (
	"context"
	"github.com/gin-gonic/gin"
	error2 "github.com/quanxiang-cloud/cabin/error"
	ginheader "github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/cabin/tailormade/resp"
	"github.com/quanxiang-cloud/faas/internal/logic"
	"github.com/quanxiang-cloud/faas/pkg/code"
	"github.com/quanxiang-cloud/faas/pkg/config"
	"github.com/quanxiang-cloud/faas/pkg/k8s"
	"gorm.io/gorm"
)

// Function Function api
type Function struct {
	fn logic.Function
}

// NewFunctionAPI new
func NewFunctionAPI(c context.Context, conf *config.Config, db *gorm.DB, kc k8s.Client) *Function {
	return &Function{
		fn: logic.NewFunction(c, db, *conf, kc),
	}
}

// Create create
func (f *Function) Create(c *gin.Context) {
	r := &logic.CreateFunctionRequest{}
	err := c.ShouldBind(r)
	if err != nil {
		resp.Format(nil, error2.New(code.InvalidParams)).Context(c)
		return
	}
	res, err := f.fn.Create(ginheader.MutateContext(c), r)
	if err != nil {
		resp.Format(nil, err).Context(c)
		return
	}
	buildFunctionRequest := &logic.BuildFunctionRequest{ID: res.ID}
	_, err = f.fn.Build(ginheader.MutateContext(c), buildFunctionRequest)
	if err != nil {
		resp.Format(res, err).Context(c)
	}
	resp.Format(res, nil).Context(c)
}

// UpdateStatus UpdateStatus
func (f *Function) UpdateStatus(c *gin.Context) {
	r := &logic.UpdateFunctionRequest{}
	err := c.ShouldBind(r)
	if err != nil {
		resp.Format(nil, error2.New(code.InvalidParams)).Context(c)
		return
	}
	response, err := f.fn.UpdateStatus(ginheader.MutateContext(c), r)
	if err != nil {
		resp.Format(nil, err).Context(c)
		return
	}
	_, err = f.fn.DelFunction(c, &logic.DelBuildFunctionRequest{
		ID: response.ID,
	})
	if err != nil {
		resp.Format(nil, err).Context(c)
		return
	}
	resp.Format(response, nil).Context(c)
}

// Delete delete
func (f *Function) Delete(c *gin.Context) {
	r := &logic.DeleteFunctionRequest{}
	err := c.ShouldBind(r)
	if err != nil {
		resp.Format(nil, error2.New(code.InvalidParams)).Context(c)
		return
	}
	resp.Format(f.fn.Delete(ginheader.MutateContext(c), r)).Context(c)
}

// Get get
func (f *Function) Get(c *gin.Context) {
	r := &logic.GetFunctionRequest{}
	err := c.ShouldBind(r)
	if err != nil {
		resp.Format(nil, error2.New(code.InvalidParams)).Context(c)
		return
	}
	resp.Format(f.fn.Get(ginheader.MutateContext(c), r)).Context(c)
}

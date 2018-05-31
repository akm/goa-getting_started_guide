// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Contexts
//
// Command:
// $ goagen
// --design=goa-getting_started_guide/design
// --out=$(GOPATH)/src/goa-getting_started_guide
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	BottleID int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(r *GoaExampleBottle) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.bottle+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

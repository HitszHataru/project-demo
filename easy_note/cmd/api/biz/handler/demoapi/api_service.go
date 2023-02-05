// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by hertz generator.

package demoapi

import (
	"context"

	"github.com/cloudwego/biz-demo/easy_note/cmd/api/biz/model/demoapi"
	"github.com/cloudwego/biz-demo/easy_note/cmd/api/biz/mw"
	"github.com/cloudwego/biz-demo/easy_note/cmd/api/biz/rpc"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demonote"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demouser"
	"github.com/cloudwego/biz-demo/easy_note/pkg/consts"
	"github.com/cloudwego/biz-demo/easy_note/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// CreateUser .
// @router /v1/user/register [POST]
func CreateUser(_ context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	err = rpc.CreateUser(context.Background(), &demouser.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// CheckUser .
// @router /v1/user/login [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// CreateNote .
// @router /v1/note [POST]
func CreateNote(_ context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.CreateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	err = rpc.CreateNote(context.Background(), &demonote.CreateNoteRequest{
		Title:   req.Title,
		Content: req.Content,
		UserId:  v.(*demoapi.User).UserID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// QueryNote .
// @router /v1/note/query [GET]
func QueryNote(_ context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.QueryNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	notes, total, err := rpc.QueryNotes(context.Background(), &demonote.QueryNoteRequest{
		UserId:    v.(*demoapi.User).UserID,
		SearchKey: req.SearchKey,
		Offset:    req.Offset,
		Limit:     req.Limit,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, utils.H{
		consts.Total: total,
		consts.Notes: notes,
	})
}

// UpdateNote .
// @router /v1/note/:note_id [PUT]
func UpdateNote(_ context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.UpdateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	err = rpc.UpdateNote(context.Background(), &demonote.UpdateNoteRequest{
		NoteId:  req.NoteID,
		UserId:  v.(*demoapi.User).UserID,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// DeleteNote .
// @router /v1/note/:note_id [DELETE]
func DeleteNote(_ context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.DeleteNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	err = rpc.DeleteNote(context.Background(), &demonote.DeleteNoteRequest{
		NoteId: req.NoteID,
		UserId: v.(*demoapi.User).UserID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// Package human_authentication code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkhuman_authentication

import (
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

const (
	UserIdTypeOpenId  = "open_id"  // 以open_id来识别用户
	UserIdTypeUserId  = "user_id"  // 以user_id来识别用户
	UserIdTypeUnionId = "union_id" // 以union_id来识别用户
)

type Identity struct {
	IdentityName *string `json:"identity_name,omitempty"` // user identity name
	IdentityCode *string `json:"identity_code,omitempty"` // user identity code
	Mobile       *string `json:"mobile,omitempty"`        // user mobile
}

type IdentityBuilder struct {
	identityName     string // user identity name
	identityNameFlag bool
	identityCode     string // user identity code
	identityCodeFlag bool
	mobile           string // user mobile
	mobileFlag       bool
}

func NewIdentityBuilder() *IdentityBuilder {
	builder := &IdentityBuilder{}
	return builder
}

// user identity name
//
// 示例值：
func (builder *IdentityBuilder) IdentityName(identityName string) *IdentityBuilder {
	builder.identityName = identityName
	builder.identityNameFlag = true
	return builder
}

// user identity code
//
// 示例值：
func (builder *IdentityBuilder) IdentityCode(identityCode string) *IdentityBuilder {
	builder.identityCode = identityCode
	builder.identityCodeFlag = true
	return builder
}

// user mobile
//
// 示例值：
func (builder *IdentityBuilder) Mobile(mobile string) *IdentityBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}

func (builder *IdentityBuilder) Build() *Identity {
	req := &Identity{}
	if builder.identityNameFlag {
		req.IdentityName = &builder.identityName

	}
	if builder.identityCodeFlag {
		req.IdentityCode = &builder.identityCode

	}
	if builder.mobileFlag {
		req.Mobile = &builder.mobile

	}
	return req
}

type CreateIdentityReqBodyBuilder struct {
	identityName     string // 姓名
	identityNameFlag bool
	identityCode     string // 身份证号
	identityCodeFlag bool
	mobile           string // 手机号
	mobileFlag       bool
}

func NewCreateIdentityReqBodyBuilder() *CreateIdentityReqBodyBuilder {
	builder := &CreateIdentityReqBodyBuilder{}
	return builder
}

// 姓名
//
//示例值：张三
func (builder *CreateIdentityReqBodyBuilder) IdentityName(identityName string) *CreateIdentityReqBodyBuilder {
	builder.identityName = identityName
	builder.identityNameFlag = true
	return builder
}

// 身份证号
//
//示例值：4xxxxxxxx
func (builder *CreateIdentityReqBodyBuilder) IdentityCode(identityCode string) *CreateIdentityReqBodyBuilder {
	builder.identityCode = identityCode
	builder.identityCodeFlag = true
	return builder
}

// 手机号
//
//示例值：13xxxxxxx
func (builder *CreateIdentityReqBodyBuilder) Mobile(mobile string) *CreateIdentityReqBodyBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}

func (builder *CreateIdentityReqBodyBuilder) Build() *CreateIdentityReqBody {
	req := &CreateIdentityReqBody{}
	if builder.identityNameFlag {
		req.IdentityName = &builder.identityName
	}
	if builder.identityCodeFlag {
		req.IdentityCode = &builder.identityCode
	}
	if builder.mobileFlag {
		req.Mobile = &builder.mobile
	}
	return req
}

type CreateIdentityPathReqBodyBuilder struct {
	identityName     string // 姓名
	identityNameFlag bool
	identityCode     string // 身份证号
	identityCodeFlag bool
	mobile           string // 手机号
	mobileFlag       bool
}

func NewCreateIdentityPathReqBodyBuilder() *CreateIdentityPathReqBodyBuilder {
	builder := &CreateIdentityPathReqBodyBuilder{}
	return builder
}

// 姓名
//
// 示例值：张三
func (builder *CreateIdentityPathReqBodyBuilder) IdentityName(identityName string) *CreateIdentityPathReqBodyBuilder {
	builder.identityName = identityName
	builder.identityNameFlag = true
	return builder
}

// 身份证号
//
// 示例值：4xxxxxxxx
func (builder *CreateIdentityPathReqBodyBuilder) IdentityCode(identityCode string) *CreateIdentityPathReqBodyBuilder {
	builder.identityCode = identityCode
	builder.identityCodeFlag = true
	return builder
}

// 手机号
//
// 示例值：13xxxxxxx
func (builder *CreateIdentityPathReqBodyBuilder) Mobile(mobile string) *CreateIdentityPathReqBodyBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}

func (builder *CreateIdentityPathReqBodyBuilder) Build() (*CreateIdentityReqBody, error) {
	req := &CreateIdentityReqBody{}
	if builder.identityNameFlag {
		req.IdentityName = &builder.identityName
	}
	if builder.identityCodeFlag {
		req.IdentityCode = &builder.identityCode
	}
	if builder.mobileFlag {
		req.Mobile = &builder.mobile
	}
	return req, nil
}

type CreateIdentityReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *CreateIdentityReqBody
}

func NewCreateIdentityReqBuilder() *CreateIdentityReqBuilder {
	builder := &CreateIdentityReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 用户的唯一标识（使用的ID类型见下一参数描述，不同ID类型的区别和获取，参考文档：[如何获得 User ID、Open ID 和 Union ID？](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)）
//
// 示例值：ou_2eb5483cb377daa5054bc6f86e2089a5
func (builder *CreateIdentityReqBuilder) UserId(userId string) *CreateIdentityReqBuilder {
	builder.apiReq.QueryParams.Set("user_id", fmt.Sprint(userId))
	return builder
}

// 用户ID类型 open_id/user_id/union_id
//
// 示例值：
func (builder *CreateIdentityReqBuilder) UserIdType(userIdType string) *CreateIdentityReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

// 该接口用于录入实名认证的身份信息，在唤起有源活体认证前，需要使用该接口进行实名认证。
func (builder *CreateIdentityReqBuilder) Body(body *CreateIdentityReqBody) *CreateIdentityReqBuilder {
	builder.body = body
	return builder
}

func (builder *CreateIdentityReqBuilder) Build() *CreateIdentityReq {
	req := &CreateIdentityReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.body
	return req
}

type CreateIdentityReqBody struct {
	IdentityName *string `json:"identity_name,omitempty"` // 姓名
	IdentityCode *string `json:"identity_code,omitempty"` // 身份证号
	Mobile       *string `json:"mobile,omitempty"`        // 手机号
}

type CreateIdentityReq struct {
	apiReq *larkcore.ApiReq
	Body   *CreateIdentityReqBody `body:""`
}

type CreateIdentityRespData struct {
	VerifyUid *string `json:"verify_uid,omitempty"` // 用户绑定实名身份的uid
}

type CreateIdentityResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *CreateIdentityRespData `json:"data"` // 业务数据
}

func (resp *CreateIdentityResp) Success() bool {
	return resp.Code == 0
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"golang.org/x/crypto/pbkdf2"
)

var (
	_ function.Function = Pbkdf2Sha512Function{}
)

func NewPbkdf2Sha512Function() function.Function {
	return Pbkdf2Sha512Function{}
}

type Pbkdf2Sha512Function struct{}

func (r Pbkdf2Sha512Function) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "pbkdf2_sha512"
}

func (r Pbkdf2Sha512Function) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Example function",
		MarkdownDescription: "Echoes given argument as result",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "String to echo",
			},
			function.StringParameter{
				Name:                "salt",
				MarkdownDescription: "String to echo",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r Pbkdf2Sha512Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string
	var salt string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data, &salt))

	if resp.Error != nil {
		return
	}

	saltBytes, err := base64.StdEncoding.DecodeString(salt)

	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Given salt could not be decoded into bytes: "+err.Error()))
		return
	}

	// https://docs.python.org/3/library/hashlib.html#key-derivation to use the digest size of the hashing algorithm
	keyLen := sha512.Size

	dk := pbkdf2.Key([]byte(data), saltBytes, 101, keyLen, sha512.New)
	output := fmt.Sprintf("$%s$%d$%s$%s", "7", 101, base64.StdEncoding.EncodeToString(saltBytes), base64.StdEncoding.EncodeToString(dk))

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, output))
}

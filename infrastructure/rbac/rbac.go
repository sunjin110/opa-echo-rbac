package rbac

import (
	"context"
	"opa-echo-test/infrastructure/opa"
)

const (
	fileName = "rbac.rego"
	query    = "data.rbac.authz.allow"
)

// 後でopa.goから持ってくる

// Input .
type Input struct {
	User              string   `opa:"user" json:"user"`                       // user名
	Roles             []string `opa:"roles" json:"roles"`                     // 所持しているrole
	AllowResourceList []string `opa:"allow_resources" json:"allow_resources"` // 許可されているresource 正規表現も可

	Method         string `opa:"method" json:"method"`                   // 今回アクセスするmethod
	Path           string `opa:"path" json:"path"`                       // 今回アクセスするpath
	AccessResource string `opa:"access_resource" json:"access_resource"` // 今回アクセスするresource
}

func Setup(modle []byte) {
	opa.Setup(fileName, modle)
}

func Eval(ctx context.Context, input Input) bool {
	return opa.EvalAllowed(ctx, fileName, query, input, nil)
}

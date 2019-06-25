package validate

import (
	"echo_shop/pkg/utils"
	"strings"
)

type (
	// ValidatorFunc 验证器函数 (返回的 string 不为空即表示验证失败)
	ValidatorFunc func() (msg string)
	// Validators 验证器数组 map
	Validators map[string][]ValidatorFunc
	// Messages 错误信息
	Messages map[string][]string

	// Validater -
	Validater interface {
		// IsStrict : 严格模式时，第一个验证出错时，即会停止其他验证
		IsStrict() bool
		// Validators : 注册验证器 map
		Validators() Validators
		// RegisterMessages : 注册错误信息 map
		Messages() Messages
	}
)

// Error error interface
func (msg Messages) Error() string {
	var val strings.Builder
	for k, v := range msg {
		val.WriteString(k + ": " + strings.Join(v, ",") + "\n")
	}

	return val.String()
}

// BaseValidate -
type BaseValidate struct{}

// IsStrict 是否为严格模式
func (*BaseValidate) IsStrict() bool {
	return true
}

// Validators : 注册验证器
// 验证器数组按顺序验证，一旦验证没通过，即结束该字段的验证
func (*BaseValidate) Validators() Validators {
	return Validators{}
}

// Messages 注册错误信息
func (*BaseValidate) Messages() Messages {
	return Messages{}
}

// Run 执行验证
func Run(v Validater) (errMap Messages, ok bool) {
	strict := v.IsStrict()
	validatorMap := v.Validators()
	messageMap := v.Messages()

	errMap = make(Messages)
	ok = true

	for key, validators := range validatorMap {
		customMsgArr := messageMap[key] // 自定义的错误信息
		customMsgArrLen := len(customMsgArr)

		for i, fn := range validators {
			errMsg := fn() // 执行验证函数
			if errMsg != "" {
				ok = false

				if i < customMsgArrLen && customMsgArr[i] != "" {
					// 采用自定义的错误信息输出
					errMsg = customMsgArr[i]
				} else {
					// 采用默认的错误信息输出，错误信息中可使用 $name 替换字段名
					errMsg = utils.ParseEasyTemplate(errMsg, map[string]string{
						"$name": key,
					})
				}

				if errMap[key] == nil {
					errMap[key] = make([]string, 0)
				}
				errMap[key] = append(errMap[key], errMsg)

				if strict {
					return // 严格模式: 结束所有验证
				}
			}
		}
	}

	return
}

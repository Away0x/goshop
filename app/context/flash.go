package context

// FlashUserVerification 用户需要激活
func (a *AppContext) FlashUserVerification() {
	a.FlashWarningMessage("您的账号未激活，请检查邮箱中的注册邮件进行激活或重新发送注册邮件")
}

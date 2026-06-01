package command

// GPG 插件相关命令
// gpg 插件用于对构件进行 GPG 签名，发布到 Maven Central 的必要步骤

// GpgSign 对构件进行 GPG 签名（mvn gpg:sign）
func GpgSign(executable string) (string, error) {
	return ExecForStdout(executable, "gpg:sign")
}

package command

// EffectivePom 显示当前项目的有效 POM 配置，合并了所有激活的 profile
func EffectivePom(executable string) (string, error) {
	return ExecForStdout(executable, "help:effective-pom")
}

// EffectiveSettings 显示当前有效的 Maven 设置，合并了全局和用户级 settings.xml
func EffectiveSettings(executable string) (string, error) {
	return ExecForStdout(executable, "help:effective-settings")
}

// ActiveProfiles 显示当前项目中所有激活的 Maven profile
func ActiveProfiles(executable string) (string, error) {
	return ExecForStdout(executable, "help:active-profiles")
}

// DescribePlugin 描述指定 Maven 插件的目标详细信息
func DescribePlugin(executable string, plugin string) (string, error) {
	return ExecForStdout(executable, "help:describe", "-Dplugin="+plugin)
}
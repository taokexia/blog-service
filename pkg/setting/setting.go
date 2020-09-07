package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}
// 用于初始化本项目配置的基础属性
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config") // 设定配置文件的名称为config
    vp.AddConfigPath("configs/") // 设置其配置路径为相对路径configs/
	vp.SetConfigType("yaml") // 配置类型为yaml
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

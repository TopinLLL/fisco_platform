package global

import "context"

var (
	GlobalConfig Global
	BackGround   = context.Background()
)

type Global struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
}

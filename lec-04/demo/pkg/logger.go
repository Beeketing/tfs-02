package pkg

import "go.uber.org/zap"

func NewFileLogger(filepath string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	if filepath != "" {
		cfg.OutputPaths = []string{
			filepath,
		}
	}
	return cfg.Build()
}

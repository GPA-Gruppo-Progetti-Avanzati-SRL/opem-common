package linkedservices

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-common/linkedservices/hermodr"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-aws-common/s3/awss3lks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
)

type Config struct {
	Mongo            []mongolks.Config `mapstructure:"mongo-db,omitempty" json:"mongo-db,omitempty" yaml:"mongo-db,omitempty"`
	S3               []awss3lks.Config `mapstructure:"aws-s3,omitempty" json:"aws-s3,omitempty" yaml:"aws-s3,omitempty"`
	HermodrClientCfg *hermodr.Config   `mapstructure:"hermodr,omitempty" json:"hermodr,omitempty" yaml:"hermodr,omitempty"`
}

func (c *Config) PostProcess() error {
	return nil
}

package linkedservices

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-common/linkedservices/hermodr"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-aws-common/s3/awss3lks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-kafka-common/kafkalks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"

	"github.com/rs/zerolog/log"
)

type ServiceRegistry struct {
	HermodrLks *hermodr.LinkedService
}

var registry ServiceRegistry

func InitRegistry(cfg *Config) error {

	registry = ServiceRegistry{}
	log.Info().Msg("initialize services registry")

	_, err := mongolks.Initialize(cfg.Mongo)
	if err != nil {
		return err
	}

	_, err = awss3lks.Initialize(cfg.S3)
	if err != nil {
		return err
	}

	_, err = kafkalks.Initialize(cfg.Kafka)
	if err != nil {
		return err
	}

	err = initializeHermodrClientLinkedService(cfg.HermodrClientCfg)
	if err != nil {
		return err
	}

	return nil
}

/*
 * TokensApiClient Initialization
 */

func initializeHermodrClientLinkedService(cfg *hermodr.Config) error {
	const semLogContext = "service-registry::initialize-hermodr-client-provider"
	log.Info().Msg(semLogContext)
	if cfg != nil {
		lks, err := hermodr.NewInstanceWithConfig(cfg)
		if err != nil {
			return err
		}

		registry.HermodrLks = lks
	}

	return nil
}

func NewHermodrClient() (*hermodr.Client, error) {
	return registry.HermodrLks.NewClient()
}

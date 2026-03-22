package hermodr

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-http-client/restclient"
	"github.com/rs/zerolog/log"
)

type LinkedService struct {
	Cfg           *Config
	httpClientLks *restclient.LinkedService
}

func NewInstanceWithConfig(cfg *Config) (*LinkedService, error) {
	var err error
	lks := &LinkedService{Cfg: cfg}

	lks.httpClientLks, err = restclient.NewInstanceWithConfig(cfg.HtttpClient)
	return lks, err
}

func (lks *LinkedService) NewClient() (*Client, error) {
	const semLogContext = "cobol-parser-lks::new-client"
	httpCli, err := lks.httpClientLks.NewClient()
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	return &Client{cfg: lks.Cfg, httpClient: httpCli}, nil

}

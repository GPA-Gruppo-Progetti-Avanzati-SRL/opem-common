package hermodr

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-http-client/restclient"
)

type LogoutHandlingConfig struct {
	Url string `json:"url,omitempty" yaml:"url,omitempty" mapstructure:"url,omitempty"`
}

type Config struct {
	HtttpClient    *restclient.Config   `yaml:"http-client,omitempty" mapstructure:"http-client,omitempty" json:"http-client,omitempty"`
	LogoutHandling LogoutHandlingConfig `yaml:"logout,omitempty" mapstructure:"logout,omitempty" json:"logout,omitempty"`
}

type Client struct {
	cfg        *Config
	httpClient *restclient.Client
}

func (c Client) Close() {
	if c.httpClient != nil {
		c.httpClient.Close()
		c.httpClient = nil
	}
}

/*
func (c *Client) Post(req *Request) (int, *Response, error) {
	const semLogContext = "cob-parse-client::post"

	urlBuilder := har.UrlBuilder{}
	urlBuilder.WithScheme(c.cfg.HttpScheme)
	urlBuilder.WithHostname(c.cfg.HostName)
	urlBuilder.WithPort(c.cfg.ServerPort)
	urlBuilder.WithPath(c.cfg.Url)

	reqBody := req.MustToJson()

	reqHeaders := []har.NameValuePair{{Name: "Content-type", Value: "application/json"}, {Name: "Accept", Value: "application/json"}}
	request, err := c.httpClient.NewRequest(http.MethodPost, urlBuilder.Url(), reqBody, reqHeaders, nil)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return http.StatusInternalServerError, nil, err
	}

	harEntry, err := c.httpClient.Execute(request)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return http.StatusInternalServerError, nil, err
	}

	if harEntry != nil && harEntry.Response != nil && harEntry.Response.Content != nil {
		var resp *Response
		resp, err = DeserializeJson(harEntry.Response.Content.Data)
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return http.StatusInternalServerError, nil, err
		}
		return harEntry.Response.Status, resp, nil
	}

	log.Error().Msg(semLogContext + " - no response content")
	return http.StatusInternalServerError, nil, nil
}
*/

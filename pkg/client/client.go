package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ogre0403/iscsi-target-client/pkg/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	CreateVolumeEndpoint = "/createVol"
	AttachLunEndpoint    = "/attachLun"
	DeleteTargetEndpoint = "/deleteTar"
	DeleteVolumeEndpoint = "/deleteVol"
)

type Client struct {
	client       *http.Client
	serverIP     string
	serverConfig *model.ServerCfg
}

func NewClient(server string, serverconf *model.ServerCfg) *Client {
	return &Client{
		client:       &http.Client{},
		serverIP:     server,
		serverConfig: serverconf,
	}
}

func (c *Client) AttachLun(lun *model.Lun) error {
	return c.request(http.MethodPost, AttachLunEndpoint, lun)
}

func (c *Client) CreateVolume(vol *model.Volume) error {
	return c.request(http.MethodPost, CreateVolumeEndpoint, vol)
}

func (c *Client) DeleteVolume(vol *model.Volume) error {
	return c.request(http.MethodDelete, DeleteVolumeEndpoint, vol)
}

func (c *Client) DeleteTarget(target *model.Target) error {
	return c.request(http.MethodDelete, DeleteTargetEndpoint, target)
}

func (c *Client) getEndpoint(endpoint string) string {
	return fmt.Sprintf("http://%s:%s%s", c.serverIP, strconv.Itoa(c.serverConfig.Port), endpoint)
}

func (c *Client) request(method, endpoint string, config interface{}) error {
	b, err := json.Marshal(config)

	if err != nil {
		return nil
	}

	req, err := http.NewRequest(
		method, c.getEndpoint(endpoint), bytes.NewBuffer(b))

	if err != nil {
		return err
	}

	req.SetBasicAuth(c.serverConfig.Username, c.serverConfig.Password)

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("request is not authorized")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	r := model.Response{}
	err = json.Unmarshal(body, &r)

	if err != nil {
		return err
	}

	if r.Error != false {
		return errors.New(r.Message)
	}
	return nil
}

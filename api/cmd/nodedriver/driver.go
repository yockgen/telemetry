package main

import (
	"fmt"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
)

type Driver struct {
	*drivers.BaseDriver
	client      *Client
	AccessToken string
	URL         string
	NodeID      string
}

func NewDriver(hostName, storePath string) *Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			EnvVar: "EDGEIAAS_ACCESS_TOKEN",
			Name:   "edgeiaas-access-token",
			Usage:  "Edge IaaS access token",
		},
		mcnflag.StringFlag{
			EnvVar: "EDGEIAAS_URL",
			Name:   "edgeiaas-url",
			Usage:  "Edge IaaS URL",
		},
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "edgeiaas"
}

func (d *Driver) GetState() (state.State, error) {
	c := d.getClient()
	node, err := c.GetNodeByID(d.NodeID)
	if err != nil {
		return state.None, err
	}
	switch node.State {
	case "online", "up", "on":
		return state.Running, nil
	case "offline", "down", "off":
		return state.Stopped, nil
	case "Unknown", "unknown":
		return state.None, nil
	default:
		return state.Running, nil
	}
}

func (d *Driver) Create() error {
	c := d.getClient()
	node, err := c.AcquireNode(Filter{})
	if err != nil {
		return err
	}
	d.IPAddress = node.IP
	d.NodeID = node.ID
	return nil
}

func (d *Driver) Start() error {
	c := d.getClient()
	err := c.Action(d.NodeID, "PowerOn")
	return err
}

func (d *Driver) Stop() error {
	c := d.getClient()
	err := c.Action(d.NodeID, "PowerOff")
	return err
}

func (d *Driver) Restart() error {
	if err := d.Stop(); err != nil {
		return err
	}
	if err := d.Start(); err != nil {
		return err
	}
	return nil
}

func (d *Driver) Kill() error {
	return d.Stop()
}

func (d *Driver) Remove() error {
	if err := d.Stop(); err != nil {
		return err
	}
	c := d.getClient()
	if err := c.ReleaseNode(d.NodeID); err != nil {
		return err
	}
	return nil
}

func (d *Driver) Upgrade() error {
	return nil
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.AccessToken = flags.String("edgeiaas-access-token")
	d.URL = flags.String("edgeiaas-url")
	return nil
}

func (d *Driver) GetURL() (string, error) {
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}
	if ip == "" {
		return "", nil
	}
	return fmt.Sprintf("tcp://%s:2376", ip), nil
}

func (d *Driver) GetMachineName() string {
	return d.MachineName
}

func (d *Driver) GetIP() (string, error) {
	if d.IPAddress == "" {
		return "", fmt.Errorf("IP address is not set")
	}
	return d.IPAddress, nil
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

func (d *Driver) GetSSHKeyPath() string {
	return d.SSHKeyPath
}

func (d *Driver) GetSSHPort() (int, error) {
	return 0, nil
}

func (d *Driver) GetSSHUsername() string {
	return ""
}

func (d *Driver) getClient() *Client {
	if d.client == nil {
		d.client = NewClient(d.AccessToken, d.URL)
	}
	return d.client
}

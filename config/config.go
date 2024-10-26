package config

import (
	yaml "gopkg.in/yaml.v2"
)

// Config represents the configuration for the exporter
type Config struct {
	Devices  []Device `yaml:"devices"`
	Features Features `yaml:"features,omitempty"`
}

type Features struct {
	BGP       bool `yaml:"bgp,omitempty"`
	Conntrack bool `yaml:"conntrack,omitempty"`
	DHCP      bool `yaml:"dhcp,omitempty"`
	DHCPL     bool `yaml:"dhcpl,omitempty"`
	DHCPv6    bool `yaml:"dhcpv6,omitempty"`
	Firmware  bool `yaml:"firmware,omitempty"`
	Health    bool `yaml:"health,omitempty"`
	Routes    bool `yaml:"routes,omitempty"`
	POE       bool `yaml:"poe,omitempty"`
	Pools     bool `yaml:"pools,omitempty"`
	Optics    bool `yaml:"optics,omitempty"`
	W60G      bool `yaml:"w60g,omitempty"`
	WlanSTA   bool `yaml:"wlansta,omitempty"`
	Capsman   bool `yaml:"capsman,omitempty"`
	WlanIF    bool `yaml:"wlanif,omitempty"`
	Monitor   bool `yaml:"monitor,omitempty"`
	Ipsec     bool `yaml:"ipsec,omitempty"`
	Lte       bool `yaml:"lte,omitempty"`
	Netwatch  bool `yaml:"netwatch,omitempty"`
}

// Device represents a target device
type Device struct {
	Name     string    `yaml:"name"`
	Address  string    `yaml:"address,omitempty"`
	Srv      SrvRecord `yaml:"srv,omitempty"`
	User     string    `yaml:"user"`
	Password string    `yaml:"password"`
	Port     string    `yaml:"port"`
}

type SrvRecord struct {
	Record string    `yaml:"record"`
	Dns    DnsServer `yaml:"dns,omitempty"`
}

type DnsServer struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

// Load reads YAML from reader and unmashals in Config
func Load(data []byte) (*Config, error) {
	c := &Config{}

	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}

	return c, nil
}

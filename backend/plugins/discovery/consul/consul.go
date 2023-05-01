package consul

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	consul "github.com/hashicorp/consul/api"
	sctx "github.com/viettranx/service-context"
	"supertruyen/plugins/discovery"
)

type consulComponent struct {
	id          string
	serviceName string
	logger      sctx.Logger
	client      *consul.Client
	instanceID  string
	host        string
}

func NewConsulComponent(id string, serviceName string) *consulComponent {
	return &consulComponent{id: id, serviceName: serviceName}
}

func (c *consulComponent) ID() string {
	return c.id
}

func (c *consulComponent) InitFlags() {
	flag.StringVar(&c.host, "consul_host", "localhost:8500", "consult host, should be localhost:8500")
}

func (c *consulComponent) Activate(sc sctx.ServiceContext) error {
	c.logger = sctx.GlobalLogger().GetLogger(c.id)

	config := consul.DefaultConfig()
	config.Address = c.host
	client, err := consul.NewClient(config)
	if err != nil {
		return err
	}

	c.client = client

	c.instanceID = discovery.GenerateInstanceID(c.serviceName)
	if err = c.Register(context.Background(), c.instanceID, c.serviceName, c.host); err != nil {
		return err
	}

	go func() {
		for {
			if err := c.ReportHealthyState(c.instanceID, c.serviceName); err != nil {
				c.logger.Errorln("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return err
}

func (c *consulComponent) Stop() error {
	fmt.Println("STOPPPPPPPPPPPPPP")
	return c.Deregister(context.Background(), c.instanceID, c.serviceName)
}

func (c *consulComponent) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("hostPort must be in a form of <host>:<port>, example: localhost:8081")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	return c.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: parts[0],
		ID:      instanceID,
		Name:    serviceName,
		Port:    port,
		Check:   &consul.AgentServiceCheck{CheckID: instanceID, TTL: "5s"},
	})
}

func (c *consulComponent) Deregister(ctx context.Context, instanceID string, serviceName string) error {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	err := c.client.Agent().ServiceDeregister(instanceID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("<<<<<<<<<<<<<<<<<<")
	return err
}

func (c *consulComponent) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := c.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string
	for _, e := range entries {
		res = append(res, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port))
	}
	return res, nil
}

func (c *consulComponent) ReportHealthyState(instanceID string, serviceName string) error {
	return c.client.Agent().PassTTL(instanceID, "")
}

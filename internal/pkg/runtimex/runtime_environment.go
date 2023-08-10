package runtimex

import (
	"errors"

	"github.com/samber/lo"

	"cms-be/internal/pkg/runtimex/ip"
)

type RuntimeEnvironment struct {
	IP IP
}

func Load() (RuntimeEnvironment, error) {
	ipInfo, err := loadIP()
	if err != nil {
		return lo.Empty[RuntimeEnvironment](), errors.Join(err, errors.New("load IP"))
	}

	re := RuntimeEnvironment{
		IP: ipInfo,
	}

	return re, nil
}

type IP struct {
	V4 string
	V6 string
}

func loadIP() (IP, error) {
	ipv4, ipv6, err := ip.Get()
	if err != nil {
		return lo.Empty[IP](), errors.Join(err, errors.New("get IP"))
	}

	return IP{
		V4: ipv4,
		V6: ipv6,
	}, nil
}

package mob

import (
	"context"
	"sync"
)

var (
	MobSession mobSession
	Mutex      sync.Mutex
)

type mobSession struct {
	SessionName   string   `json:"session_name"`
	Mobbers       []string `json:"mobbers"`
	State         string   `json:"state"`
	GitRepo       string   `json:"git_repo"`
	Driver        string   `json:"driver"`
	Navigator     string   `json:"navigator"`
	Duration      int      `json:"duration"`
	Retro         bool     `json:"retro"`
	LastDriver    string   `json:"last_driver"`
	LastNavigator string   `json:"last_naviagtor"`
	RetroCounter  int      `json:"retro_counter"`
}

var CTX = context.Background()

type NewMobSession struct {
	Mobber   string `json:"mobber"`
	GitRepo  string `json:"git_repo"`
	Duration int    `json:"duration"`
	Retro    bool   `json:"retro"`
}

type JoinMob struct {
	Mobber string `json:"mobber"`
}

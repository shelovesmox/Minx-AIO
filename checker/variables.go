package checker

import (
	"sync"
	"time"
)

// Global variables
var (
	Bad            uint64
	Good           uint64
	Custom         uint64
	Cpm            uint64
	CpmAverages    []uint64
	Errors         uint64
	Checking       bool
	Stopping       bool
	Proxies        []string
	LockedProxies  []string
	BadProxies     []string
	Accounts       []string
	Remaining      []string
	TotalProxies   uint64
	TotalAccounts  uint64
	CurrentTime    time.Time
	DiscordWebhook string
	LockProxies    bool
	Cui            bool
	Retries        uint64
	Timeout        uint64
	Threads        uint64
	ProxyType      string
	Pool           interface{}
	ProxyLock      sync.Mutex
	SaveLock       sync.Mutex
	PrintLock      sync.Mutex
)

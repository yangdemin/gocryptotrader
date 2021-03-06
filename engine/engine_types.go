package engine

import "time"

// Settings stores engine params
type Settings struct {
	ConfigFile            string
	DataDir               string
	MigrationDir          string
	LogFile               string
	GoMaxProcs            int
	CheckParamInteraction bool

	// Core Settings
	EnableDryRun                bool
	EnableAllExchanges          bool
	EnableAllPairs              bool
	EnableCoinmarketcapAnalysis bool
	EnablePortfolioManager      bool
	EnableGRPC                  bool
	EnableGRPCProxy             bool
	EnableWebsocketRPC          bool
	EnableDeprecatedRPC         bool
	EnableCommsRelayer          bool
	EnableExchangeSyncManager   bool
	EnableDepositAddressManager bool
	EnableEventManager          bool
	EnableOrderManager          bool
	EnableConnectivityMonitor   bool
	EnableDatabaseManager       bool
	EnableNTPClient             bool
	EnableWebsocketRoutine      bool
	EventManagerDelay           time.Duration
	Verbose                     bool

	// Exchange syncer settings
	EnableTickerSyncing    bool
	EnableOrderbookSyncing bool
	EnableTradeSyncing     bool
	SyncWorkers            int
	SyncContinuously       bool
	SyncTimeout            time.Duration

	// Forex settings
	EnableCurrencyConverter bool
	EnableCurrencyLayer     bool
	EnableFixer             bool
	EnableOpenExchangeRates bool

	// Exchange tuning settings
	EnableExchangeHTTPRateLimiter  bool
	EnableExchangeHTTPDebugging    bool
	EnableExchangeVerbose          bool
	ExchangePurgeCredentials       bool
	EnableExchangeAutoPairUpdates  bool
	DisableExchangeAutoPairUpdates bool
	EnableExchangeRESTSupport      bool
	EnableExchangeWebsocketSupport bool
	MaxHTTPRequestJobsLimit        int
	RequestTimeoutRetryAttempts    int

	// Global HTTP related settings
	GlobalHTTPTimeout   time.Duration
	GlobalHTTPUserAgent string
	GlobalHTTPProxy     string

	// Exchange HTTP related settings
	ExchangeHTTPTimeout   time.Duration
	ExchangeHTTPUserAgent string
	ExchangeHTTPProxy     string

	// Dispatch system settings
	EnableDispatcher        bool
	DispatchMaxWorkerAmount int
	DispatchJobsLimit       int
}

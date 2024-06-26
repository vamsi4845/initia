package genesis_markets

const GenesisMarkets = `[
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "BNB",
		  "Quote": "USD"
		},
		"decimals": 7,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "BNB-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "BNB-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "BNBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "binance_api",
		  "off_chain_ticker": "BNBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "BNBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "BNB_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "BTC",
		  "Quote": "USD"
		},
		"decimals": 5,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "BTCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "BTCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "BTC-USD"
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "btcusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "XXBTZUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "BTC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "BTCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "BTC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "DOGE",
		  "Quote": "USD"
		},
		"decimals": 11,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "DOGEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "DOGEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "DOGE-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "DOGE_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "dogeusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "XDGUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "DOGE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "DOGEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "DOGE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "EOS",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "EOSUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "EOSUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "EOS-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "EOS_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "EOSUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "EOS-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "EOS-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "EOSUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ICP",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ICPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ICPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "ICP-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "ICPUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ICP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ICP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ICPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "NTRN",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 2
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "NTRNUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "NTRN_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "NTRN-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "NTRN-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "PEPE",
		  "Quote": "USD"
		},
		"decimals": 16,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "PEPEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "PEPEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "PEPE_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "PEPEUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "PEPE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "PEPEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "PEPE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "SHIB",
		  "Quote": "USD"
		},
		"decimals": 15,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "SHIBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "SHIBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "SHIB-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "SHIB_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "SHIBUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "SHIB-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "SHIBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "SHIB-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "PYTH",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "PYTHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "PYTHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "PYTH_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "PYTH-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "PYTH-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "PYTHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "RNDR",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "RNDRUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "RNDR-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "RNDRUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "RNDR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "RNDR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "RNDRUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "STRK",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "STRKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "STRKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "STRKUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "STRK-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "STRK-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "STRK_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ALGO",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ALGOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ALGOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "ALGO-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "ALGOUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ALGO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ALGOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ALGO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "BCH",
		  "Quote": "USD"
		},
		"decimals": 7,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "BCHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "BCHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "BCH-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "BCH_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "bchusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "BCHUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "BCH-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "BCHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "BCH-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "DYDX",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "DYDXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "DYDXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "DYDX_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "DYDX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "DYDXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "DYDX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "HBAR",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "HBARUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bitstamp_ws",
		  "off_chain_ticker": "hbarusd"
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "HBARUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "HBAR-USD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "HBAR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "HBARUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "HBAR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "OP",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "OPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "OP-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "OP_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "OP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "OPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "OP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ORDI",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ORDIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ORDIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "ORDI_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "ordiusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ORDI-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ORDI-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ORDIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "TRX",
		  "Quote": "USD"
		},
		"decimals": 11,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "TRXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "TRXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "TRX_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "trxusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "TRXUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "TRX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "TRXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "TRX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "AAVE",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "AAVEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "aaveusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "AAVEUSD"
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "AAVE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "AAVE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "binance_api",
		  "off_chain_ticker": "AAVEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "AAVE-USD"
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "BLUR",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "BLUR-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "BLUR_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "BLURUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "BLUR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "BLURUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "BLUR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "UNI",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "UNIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "UNIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "UNI-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "UNI_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "UNIUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "UNI-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "UNI-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ARKM",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ARKMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ARKMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "ARKM_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ARKM-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ARKMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "JUP",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "JUP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "JUP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "binance_api",
		  "off_chain_ticker": "JUPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "JUPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "JUP_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "JUPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "LDO",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "LDOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "LDO-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "LDOUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "LDO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "LDOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "LDO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "WLD",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "WLDUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "WLDUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "WLD_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "wldusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "WLD-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "WLDUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "WLD-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "DYM",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "DYMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "DYMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "DYM_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "DYM-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "DYMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "COMP",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "COMPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "COMP-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "COMP_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "COMPUSD"
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "COMPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "COMP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "DOT",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "DOTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "DOTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "DOT-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "DOT_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "DOTUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "DOT-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "DOTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "DOT-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ETC",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ETCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "ETC-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "ETC_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "etcusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ETC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ETCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ETC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "LINK",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "LINKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "LINKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "LINK-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "LINKUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "LINK-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "LINKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "LINK-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "USDT",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 1
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "USDCUSDT",
		  "invert": true
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "USDCUSDT",
		  "invert": true
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "USDT-USD"
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "ethusdt",
		  "normalize_by_pair": {
			"Base": "ETH",
			"Quote": "USD"
		  },
		  "invert": true
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "USDTZUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "BTC-USDT",
		  "normalize_by_pair": {
			"Base": "BTC",
			"Quote": "USD"
		  },
		  "invert": true
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "USDC-USDT",
		  "invert": true
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "AVAX",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "AVAXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "AVAXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "AVAX-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "AVAX_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "avaxusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "AVAXUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "AVAX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "AVAX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "BONK",
		  "Quote": "USD"
		},
		"decimals": 14,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "BONKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "BONKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "BONK-USD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "BONK-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "BONK-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "BONKUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "CRV",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "CRVUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "CRV-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "CRV_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "CRVUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "CRV-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "CRVUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "CRV-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "IMX",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "IMXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "IMX-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "IMXUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "IMX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "IMXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "IMX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "SEI",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "SEIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "SEIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "SEI-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "SEI_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "seiusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "SEI-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "SEIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "SUI",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "SUIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "SUIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "SUI-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "SUI_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "suiusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "SUI-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "SUIUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "SUI-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "GRT",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "GRTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "GRTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "GRT-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "GRT_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "GRTUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "GRT-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "GRTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "GRT-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "MATIC",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "MATICUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "MATICUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "MATIC-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "MATIC_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "maticusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "MATICUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "MATIC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "MATICUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "MATIC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "FET",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "FETUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "FET-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "FETUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "FET-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "FET-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "FETUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "LTC",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "LTCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "LTCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "LTC-USD"
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "ltcusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "XLTCZUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "LTC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "LTCUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "LTC-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "APE",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "APEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "APE-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "APE_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "APEUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "APE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "APEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "APE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "JTO",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "JTO-USD"
		},
		{
		  "name": "binance_api",
		  "off_chain_ticker": "JTOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "JTOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "JTOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "JTO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "JTO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "SNX",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "SNXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "SNXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "SNX-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "SNXUSD"
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "SNXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "SNX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "SOL",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "SOLUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "SOLUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "SOL-USD"
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "solusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "SOLUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "SOL-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "SOLUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "SOL-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "AEVO",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "AEVOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "AEVOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "AEVO_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "AEVOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "AEVO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "AGIX",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "AGIXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "AGIXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "AGIX_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "AGIX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "AGIX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "AGIXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "FIL",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "FILUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "FIL-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "FIL_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "filusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "FILUSD"
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "FILUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "FIL-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "INJ",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "INJUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "INJUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "INJ-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "INJUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "INJ-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "INJ-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "INJUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "TIA",
		  "Quote": "USD"
		},
		"decimals": 8,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "TIAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "TIAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "TIA-USD"
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "tiausdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "TIAUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "TIA-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "TIAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "TIA-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "WOO",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "WOOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "WOO_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "WOO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "WOO-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "WOOUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "APT",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "APTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "APTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "APT-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "APT_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "aptusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "APT-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "APTUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "APT-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ARB",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ARBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ARBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "ARB-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "ARB_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "arbusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ARB-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ARBUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ARB-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ATOM",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ATOMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ATOMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "ATOM-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "ATOM_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "ATOMUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ATOM-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ATOMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ATOM-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "MKR",
		  "Quote": "USD"
		},
		"decimals": 6,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "MKRUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "MKR-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "MKRUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "MKR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "MKRUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "MKR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "STX",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "STXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "STXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "STX-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "STX_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "STXUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "STX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "STX-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "STXUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "XRP",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "XRPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "XRPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "XRP-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "XRP_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "xrpusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "XXRPZUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "XRP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "XRPUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "XRP-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ASTR",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ASTRUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "ASTR_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "ASTRUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ASTR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ASTRUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ASTR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ETH",
		  "Quote": "USD"
		},
		"decimals": 6,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ETHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ETHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "ETH-USD"
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "ethusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "XETHZUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ETH-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ETHUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ETH-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "MANA",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "MANAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "MANA-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "MANA_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "MANAUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "MANA-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "MANAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "MANA-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "AXL",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "AXLUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "AXLUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "AXL-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "WAXL_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "WAXLUSD"
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "NEAR",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "NEARUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "NEAR-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "NEAR_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "nearusdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "NEAR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "NEARUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "NEAR-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "XLM",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "XLMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "XLMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "XLM-USD"
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "XXLMZUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "XLM-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "XLMUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "XLM-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "ADA",
		  "Quote": "USD"
		},
		"decimals": 10,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "ADAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "bybit_ws",
		  "off_chain_ticker": "ADAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "coinbase_api",
		  "off_chain_ticker": "ADA-USD"
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "ADA_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "huobi_ws",
		  "off_chain_ticker": "adausdt",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "ADAUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "ADA-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "ADAUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "okx_ws",
		  "off_chain_ticker": "ADA-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	},
	{
	  "ticker": {
		"currency_pair": {
		  "Base": "RUNE",
		  "Quote": "USD"
		},
		"decimals": 9,
		"min_provider_count": 3
	  },
	  "provider_configs": [
		{
		  "name": "binance_api",
		  "off_chain_ticker": "RUNEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "gate_ws",
		  "off_chain_ticker": "RUNE_USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "kraken_api",
		  "off_chain_ticker": "RUNEUSD"
		},
		{
		  "name": "kucoin_ws",
		  "off_chain_ticker": "RUNE-USDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		},
		{
		  "name": "mexc_ws",
		  "off_chain_ticker": "RUNEUSDT",
		  "normalize_by_pair": {
			"Base": "USDT",
			"Quote": "USD"
		  }
		}
	  ]
	}
  ]`

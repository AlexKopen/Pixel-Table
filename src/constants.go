package main

var BotParameters = Parameters{
	ChangeThresholdPercentage: 0.0035,
	LossSellPercentage:        0.015,
	GainSellPercentage:        0.0015,
}

const (
	Purchase MarketOrderAction = iota
	Sell
	Wait
)

var Symbols = [...]string{
	"ADA",
	"AION",
	"ALGO",
	"ANKR",
	"ANT",
	"ARDR",
	"ARPA",
	"ATOM",
	"BAL",
	"BAND",
	"BAT",
	"BCH",
	"BEAM",
	"BEL",
	"BLZ",
	"BNB",
	"BNT",
	"BTC",
	"BTS",
	"BTT",
	"BZRX",
	"CELR",
	"CHR",
	"CHZ",
	"COCOS",
	"COMP",
	"COS",
	"COTI",
	"CRV",
	"CTSI",
	"CTXC",
	"CVC",
	"DASH",
	"DATA",
	"DCR",
	"DENT",
	"DGB",
	"DIA",
	"DOCK",
	"DOGE",
	"DOT",
	"DREP",
	"DUSK",
	"EGLD",
	"ENJ",
	"EOS",
	"ETC",
	"ETH",
	"FET",
	"FIO",
	"FTM",
	"FTT",
	"FUN",
	"GTO",
	"GXS",
	"HBAR",
	"HC",
	"HIVE",
	"HOT",
	"ICX",
	"IOST",
	"IOTA",
	"IOTX",
	"IRIS",
	"JST",
	"KAVA",
	"KEY",
	"KMD",
	"KNC",
	"KSM",
	"LEND",
	"LINK",
	"LRC",
	"LSK",
	"LTC",
	"LTO",
	"LUNA",
	"MANA",
	"MATIC",
	"MBL",
	"MCO",
	"MDT",
	"MFT",
	"MITH",
	"MKR",
	"MTL",
	"NANO",
	"NEO",
	"NKN",
	"NMR",
	"NPXS",
	"NULS",
	"OCEAN",
	"OGN",
	"OMG",
	"ONE",
	"ONG",
	"ONT",
	"OXT",
	"PAXG",
	"PERL",
	"PNT",
	"QTUM",
	"REN",
	"REP",
	"RLC",
	"RSR",
	"RUNE",
	"RVN",
	"SAND",
	"SC",
	"SNX",
	"SOL",
	"SRM",
	"STMX",
	"STORJ",
	"STPT",
	"STRAT",
	"STX",
	"SUSHI",
	"SXP",
	"TCT",
	"TFUEL",
	"THETA",
	"TOMO",
	"TRB",
	"TROY",
	"TRX",
	"UMA",
	"UNI",
	"VET",
	"VITE",
	"VTHO",
	"WAN",
	"WAVES",
	"WIN",
	"WNXM",
	"WRX",
	"WTC",
	"XLM",
	"XMR",
	"XRP",
	"XTZ",
	"XZC",
	"YFI",
	"YFII",
	"ZEC",
	"ZEN",
	"ZIL",
	"ZRX",
}

package ccl

import . "github.com/timbeurskens/goparselib"

var (
	engineType = Union{
		CTerminal("MySQL"),
		CTerminal("PostgreSQL"),
		CTerminal("MariaDB"),
		CTerminal("Oracle"),
		CTerminal("SQL Server"),
	}
	osType = Union{
		CTerminal("Linux"),
		CTerminal("Red Hat Enterprise"),
		CTerminal("Ubuntu Server"),
		CTerminal("Windows Server 2019"),
	}
	storageType = Union{
		CTerminal("BLS"),
		CTerminal("SSD"),
	}
	usWest = Union{
		CTerminal("US West"),
		CTerminal("California"),
	}
	africa = Union{
		CTerminal("Africa"),
		CTerminal("Cape Town"),
	}
	europe = Union{
		CTerminal("Europe"),
		CTerminal("Frankfurt"),
	}
	southAmerica = Union{
		CTerminal("South America"),
		CTerminal("Bogota"),
	}
	asiaPacific = Union{
		CTerminal("Asia Pacific"),
		CTerminal("Seoul"),
	}

	regionProperty         = Concat{CTerminal("region\\:"), BlankOpt, Union{usWest, africa, europe, southAmerica, asiaPacific}}
	osProperty             = Concat{CTerminal("OS\\:"), BlankOpt, osType}
	ipv6Property           = Concat{CTerminal("IPV6\\:"), BlankOpt, Union{CTerminal("yes"), CTerminal("no")}}
	storageProperty        = Concat{CTerminal("storage\\:"), BlankOpt, storageType, Blank, CTerminal("of"), Blank, Natural, BlankOpt, CTerminal("GB")}
	cpuProperty            = Concat{CTerminal("CPU\\:"), BlankOpt, Natural, BlankOpt, CTerminal("cores")}
	memoryProperty         = Concat{CTerminal("memory\\:"), BlankOpt, Natural, BlankOpt, CTerminal("GB")}
	engineProperty         = Concat{CTerminal("engine\\:"), BlankOpt, engineType}
	storageBlockProperty   = Union{regionProperty, engineProperty, cpuProperty, memoryProperty, ipv6Property, storageProperty}
	computingBlockProperty = Union{regionProperty, osProperty, ipv6Property, storageProperty, cpuProperty, memoryProperty}
	blankList              = Plus(Union{EOL, Blank})
	storageBody            = List(storageBlockProperty, Comma, blankList)
	computingBody          = List(computingBlockProperty, Comma, blankList)
	storage                = Concat{CTerminal("storage"), Blank, Ident, BlankOpt, LBracket, blankList, storageBody, blankList, RBracket}
	computing              = Concat{CTerminal("computing"), Blank, Ident, BlankOpt, LBracket, blankList, computingBody, blankList, RBracket}
	resourceBody           = List(Union{computing, storage, Ident}, Comma, blankList)
	Root                   = Concat{CTerminal("resource"), Blank, Ident, BlankOpt, LBracket, blankList, resourceBody, blankList, RBracket}
)

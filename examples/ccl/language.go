package ccl

import . "goparselib"

var (
	us_west       = Union{CTerminal("US West"), CTerminal("California")}
	africa        = Union{CTerminal("Africa"), CTerminal("Cape Town")}
	europe        = Union{CTerminal("Europe"), CTerminal("Frankfurt")}
	south_america = Union{CTerminal("South America"), CTerminal("Bogota")}
	asia_pacific  = Union{CTerminal("Asia Pacific"), CTerminal("Seoul")}

	regionProperty = Concat{CTerminal("region\\:"), BlankOpt, Union{us_west, africa, europe, south_america, asia_pacific}}

	storageProperty   = Union{regionProperty}
	computingProperty = Union{regionProperty}

	storageBody   = new(Symbol)
	computingBody = new(Symbol)

	storage      = Concat{CTerminal("storage"), Blank, Ident, BlankOpt, LBracket, R(storageBody), RBracket}
	computing    = Concat{CTerminal("computing"), Blank, Ident, BlankOpt, LBracket, R(computingBody), RBracket}
	resourceBody = new(Symbol)

	root = Concat{CTerminal("resource"), Blank, Ident, BlankOpt, LBracket, R(resourceBody), RBracket}
)

func init() {
	Define(resourceBody,
		Union{nil,
			Concat{EOL, R(resourceBody)},
			Concat{storage, Union{nil, Concat{Comma, R(resourceBody)}}},
			Concat{computing, Union{nil, Concat{Comma, R(resourceBody)}}},
		})
}

const (
	complete = `resource my_cluster {
    storage my_db {
        region: Bogota,
        engine: MySQL,
        CPU: 2 cores,
        memory: 2 GB,
        IPV6: no,
        storage: BLS of 16 GB
    },
    computing my_server {
        region: Bogota,
        OS: Linux,
        IPV6: yes,
        storage: SSD of 256 GB,
        CPU: 4 cores,
        memory: 8 GB
    },
    my_server
}`
)

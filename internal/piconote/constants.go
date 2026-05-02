package piconote

// Valid sub-commands allowed by this program
var Commands = []string{
	"view",
	"list",
	"write",
	"remove",
}

// Sub-commands that do not need the file argument; set as a map to optimize processing
var BypassFileCommands = map[string]bool{
	"list": true,
}

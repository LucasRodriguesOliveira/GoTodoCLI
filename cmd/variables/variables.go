package variables

import (
	c "GoTodoCLI/cmd/config"
)

var HomeDir string = c.HomeDir()
var CfgFile string
var Db string
var Version string = "v1.0"

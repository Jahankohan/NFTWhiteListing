package deploy

import (
	"fmt"
	"os/exec"
)

func Compile() string {
	// This Method Compile NFTWhiteList.sol to abi and bin and then build a go interface for communicating with it.
	cmd, err := exec.Command("/bin/sh", "./compile.sh").Output()
    if err != nil {
    fmt.Printf("error %s", err)
    }
    output := string(cmd)
    return output
}
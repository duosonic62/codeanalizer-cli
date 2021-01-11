package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Use: "codeanalizer",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Usage()
			if err != nil {
				Exit(err, 1)
			}
		},
	}
)

func Run()  {
	err := RootCmd.Execute()
	if err != nil {
		Exit(err, 1)
	}
}

func Exit(err error, codes ...int)  {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}

	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
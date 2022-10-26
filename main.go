package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

var root = &cobra.Command{
	Use:     "jaya [flags] class [args]",
	Short:   "jaya [flags] class [args]",
	Version: VERSION,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			_ = cmd.Help()
			return
		}
		flags.class = args[0]
		flags.args = args[1:]
		if flags.class == "" {
			_ = cmd.Help()
			return
		}
		startJYM(flags)
	},
}
var flags = new(Flags)

func main() {
	root.Flags().StringVarP(&flags.cpOption, "cp", "", "", "class path")
	root.Flags().StringVarP(&flags.cpOption, "classpath", "", "", "class path")
	root.Flags().StringVarP(&flags.class, "class", "", "", "目标启动类")
	_ = root.Execute()
}

func startJYM(flags *Flags) {
	fmt.Printf("classpath: %s class: %s args:%v\n", flags.cpOption, flags.class, flags.args)
}

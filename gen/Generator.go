package gen

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// EnumCommand go run .\main.go enum --name="-e=category -f=种类:dept-1-部门,company-2-公司,team-3-团队,group-4-组" --out=app\enums
var EnumCommand = &cobra.Command{
	Use:     "enum",
	Short:   "Enum构建工具",
	Long:    "Enum构建工具",
	Example: "enum --name=bin\\enum_cmd.md或--name=\"-e=state -f=状态:issue-1-发布,draft-2-草稿\" --out=app\\Enums",
	Run: func(cmd *cobra.Command, args []string) {
		NewEnum(name, out).Generate()
	},
}

func Execute() {
	InitEnum()
	if err := EnumCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

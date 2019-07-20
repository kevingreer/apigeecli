package getfh

import (
	"../../shared"
	"github.com/spf13/cobra"
	"net/url"
	"path"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get a flowhook",
	Long:  "Get a flowhook",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "environments", shared.RootArgs.Env, "flowhooks", name)
		_ = shared.HttpClient(u.String())
	},
}

var name string

func init() {

	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "Flowhook point")

	_ = Cmd.MarkFlagRequired("name")
}


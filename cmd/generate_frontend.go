package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/mjibson/esc/embed"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(genFrontend)
}

var genFrontend = &cobra.Command{
	Use:   "gen-frontend",
	Short: "Starts all services",
	Long:  `Starts all services.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conf := &embed.Config{
			Invocation: strings.Join(os.Args[1:], " "),
		}

		conf.OutputFile = "services/frontend/static.go"
		conf.Package = "frontend"
		conf.Prefix = "services/frontend/web_assets/build"
		conf.Files = []string{"services/frontend/web_assets/build"}

		var err error
		out := os.Stdout
		if conf.OutputFile != "" {
			if out, err = os.Create(conf.OutputFile); err != nil {
				log.Fatal(err)
			}
			defer out.Close()
		}
		if err = embed.Run(conf, out); err != nil {
			return err
		}

		return nil
	},
}

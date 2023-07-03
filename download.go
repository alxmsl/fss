package main

import (
	"github.com/alxmsl/fss/dispatcher/inmemory"
	"github.com/alxmsl/fss/separator"
	"github.com/alxmsl/fss/service"
	"github.com/alxmsl/fss/storage"
	"github.com/alxmsl/fss/storage/local"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var (
	downloadName   string
	downloadOutput string

	downloadCmd = &cobra.Command{
		Use: "download",
		Short: `
Store CLI tool stores file in the Storage Service

Tool automatically finds Storage Service and stores file into the storage 
`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				//@todo: pass storages via CLI
				storages = []storage.Interface{
					local.NewStorage("storage_1", "/tmp/storage_1", time.Second),
					local.NewStorage("storage_2", "/tmp/storage_2", time.Second),
					local.NewStorage("storage_3", "/tmp/storage_3", time.Second),
					local.NewStorage("storage_4", "/tmp/storage_4", time.Second),
					local.NewStorage("storage_5", "/tmp/storage_5", time.Second),
					local.NewStorage("storage_6", "/tmp/storage_6", time.Second),
				}
				svc = service.New(inmemory.NewDispatcher(storages), separator.DefaultSeparator)
			)

			f, err := os.OpenFile(downloadOutput, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatalln(err)
			}
			err = svc.Download(downloadName, f)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
)

func init() {
	downloadCmd.PersistentFlags().StringVar(&downloadName, "object", "", "")
	downloadCmd.PersistentFlags().StringVar(&downloadOutput, "output", "", "")

	rootCmd.AddCommand(downloadCmd)
}

package main

import (
	"fmt"
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
	uploadName string

	uploadCmd = &cobra.Command{
		Use: "upload",
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
				svc     = service.New(inmemory.NewDispatcher(storages), separator.DefaultSeparator)
				fi, err = os.Stat(uploadName)
			)
			if err != nil {
				log.Fatalln(err)
			}

			f, err := os.Open(uploadName)
			if err != nil {
				log.Fatalln(err)
			}

			res, err := svc.Upload(f, fi.Size())
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(res)
		},
	}
)

func init() {
	uploadCmd.PersistentFlags().StringVar(&uploadName, "file", "", "file")

	rootCmd.AddCommand(uploadCmd)
}

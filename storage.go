package main

import (
	"github.com/alxmsl/fss/storage/local"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	storageName string

	storageCmd = &cobra.Command{
		Use: "storage",
		Short: `
Store CLI tool stores file in the Storage Service

Tool automatically finds Storage Service and stores file into the storage 
`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				storage = local.NewStorage(storageName, "/tmp/"+storageName, time.Second)
				router  = mux.NewRouter()
				web     = http.Server{
					Addr:    storageName,
					Handler: router,
				}
			)

			router.HandleFunc("/Size", func(w http.ResponseWriter, r *http.Request) {
				var size = storage.Size()
				_, _ = w.Write([]byte(strconv.Itoa(int(size))))
			}).Methods(http.MethodGet)

			router.HandleFunc("/GetObject/{name}", func(w http.ResponseWriter, r *http.Request) {
				var err = storage.Get(mux.Vars(r)["name"], w)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}).Methods(http.MethodGet)

			router.HandleFunc("/PutObject/{name}", func(w http.ResponseWriter, r *http.Request) {
				var err = storage.Put(mux.Vars(r)["name"], r.Body)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}).Methods(http.MethodPost)

			log.Println(storageName, "starting server...")
			if err := web.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalln(err)
			}
		},
	}
)

func init() {
	storageCmd.PersistentFlags().StringVar(&storageName, "name", "", "")

	rootCmd.AddCommand(storageCmd)
}

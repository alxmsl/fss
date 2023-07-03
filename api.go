package main

import (
	"github.com/alxmsl/fss/dispatcher/inmemory"
	"github.com/alxmsl/fss/separator"
	"github.com/alxmsl/fss/service"
	"github.com/alxmsl/fss/storage"
	"github.com/alxmsl/fss/storage/remote"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"strconv"
)

var (
	apiAddr string

	apiCmd = &cobra.Command{
		Use: "api",
		Short: `
Store CLI tool stores file in the Storage Service

Tool automatically finds Storage Service and stores file into the storage 
`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				//@todo: pass storages via CLI
				storages = []storage.Interface{
					remote.NewStorage("remote_storage_1", "http://127.0.0.1:8081"),
					remote.NewStorage("remote_storage_2", "http://127.0.0.1:8082"),
					remote.NewStorage("remote_storage_3", "http://127.0.0.1:8083"),
					remote.NewStorage("remote_storage_4", "http://127.0.0.1:8084"),
					remote.NewStorage("remote_storage_5", "http://127.0.0.1:8085"),
					remote.NewStorage("remote_storage_6", "http://127.0.0.1:8086"),
				}
				svc = service.New(inmemory.NewDispatcher(storages), separator.DefaultSeparator)

				router = mux.NewRouter()
				web    = http.Server{
					Addr:    apiAddr,
					Handler: router,
				}
			)

			router.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
				var err = svc.Download(mux.Vars(r)["name"], w)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}).Methods(http.MethodGet)

			router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				var (
					contentLength = r.Header.Get("Content-Length")
					size, err     = strconv.Atoi(contentLength)
				)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				res, err := svc.Upload(r.Body, int64(size))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				_, _ = w.Write([]byte(res))
			}).Methods(http.MethodPost)

			log.Println(apiAddr, "starting server...")
			if err := web.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalln(err)
			}
		},
	}
)

func init() {
	apiCmd.PersistentFlags().StringVar(&apiAddr, "addr", "", "")

	rootCmd.AddCommand(apiCmd)
}

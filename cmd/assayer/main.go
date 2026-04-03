package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-assayer/internal/server";"github.com/stockyard-dev/stockyard-assayer/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9700"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./assayer-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("assayer: %v",err)};defer db.Close();srv:=server.New(db,server.DefaultLimits())
fmt.Printf("\n  Assayer — Self-hosted revenue and subscription tracker\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("assayer: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}

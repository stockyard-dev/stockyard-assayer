package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-assayer/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list,_:=s.db.List();if list==nil{list=[]store.RevenueEntry{}};writeJSON(w,200,list)}
func(s *Server)handleRecord(w http.ResponseWriter,r *http.Request){var e store.RevenueEntry;json.NewDecoder(r.Body).Decode(&e);if e.Period==""{writeError(w,400,"period required");return};s.db.Record(&e);writeJSON(w,201,e)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}

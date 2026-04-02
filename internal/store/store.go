package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Subscription struct {
	ID string `json:"id"`
	CustomerName string `json:"customer_name"`
	Plan string `json:"plan"`
	MRR int `json:"mrr"`
	Status string `json:"status"`
	StartDate string `json:"start_date"`
	RenewalDate string `json:"renewal_date"`
	Notes string `json:"notes"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"assayer.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS subscriptions(id TEXT PRIMARY KEY,customer_name TEXT NOT NULL,plan TEXT DEFAULT '',mrr INTEGER DEFAULT 0,status TEXT DEFAULT 'active',start_date TEXT DEFAULT '',renewal_date TEXT DEFAULT '',notes TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Subscription)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO subscriptions(id,customer_name,plan,mrr,status,start_date,renewal_date,notes,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.CustomerName,e.Plan,e.MRR,e.Status,e.StartDate,e.RenewalDate,e.Notes,e.CreatedAt);return err}
func(d *DB)Get(id string)*Subscription{var e Subscription;if d.db.QueryRow(`SELECT id,customer_name,plan,mrr,status,start_date,renewal_date,notes,created_at FROM subscriptions WHERE id=?`,id).Scan(&e.ID,&e.CustomerName,&e.Plan,&e.MRR,&e.Status,&e.StartDate,&e.RenewalDate,&e.Notes,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Subscription{rows,_:=d.db.Query(`SELECT id,customer_name,plan,mrr,status,start_date,renewal_date,notes,created_at FROM subscriptions ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Subscription;for rows.Next(){var e Subscription;rows.Scan(&e.ID,&e.CustomerName,&e.Plan,&e.MRR,&e.Status,&e.StartDate,&e.RenewalDate,&e.Notes,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM subscriptions WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM subscriptions`).Scan(&n);return n}

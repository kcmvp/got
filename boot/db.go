// nolint:all
package boot

const (
	DataSource = "datasource"
)

type dataSource struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Db       string
}

func setupDb() {
	var ds []dataSource
	if err := cfg.UnmarshalKey(DataSource, &ds); err == nil {
		if len(ds) > 0 {
			// 1: setup datasource
			// 2: inject datasource into container
			// 3: execute init.sql
		}
	} else {
		//@todo log message
	}

	//db, err := sql.Open("mysql", "user:password@/dbname")
	//if err != nil {
	//	panic(err)
	//}
	//db.Close()
	//// See "Important settings" section.
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	//db.Stats()
	//db.Close()
	//
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//defer conn.Close(context.Background())
	//
	//var name string
	//var weight int64
	//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(name, weight)
}

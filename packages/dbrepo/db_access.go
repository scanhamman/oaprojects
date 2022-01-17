package dbrepo

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "WinterIsComing!"
	dbname   = "context"
)

func ProcessProjectData(dp Project,
	dfs []Funding,
	dhs []H2020,
	dss []PSubject) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// close database
	defer db.Close()

	// store project and its details
	StoreProjectsSQL := `insert into openaire.projects (projectid, oaid, title, acronym, summary, 
		keywords, websiteurl, callidentifier, code, startdate, enddate,
		grantedcurr, grantedamount, grantedtotcost, openaccessdataset, openaccesspubs) 
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
	StoreFundingsSQL := `insert into  openaire.projectfundings 
	           (projectid, description, streamid, jurisdiction, name, shortname) values($1, $2, $3, $4, $5, $6)`
	StoreH2020sSQL := `insert into openaire.projecth2020s (projectid, code, description) values($1, $2, $3)`
	StoreSubjectsSQL := `insert into openaire.projectsubjects (projectid, subject) values($1, $2)`

	if dp.Projectid < 100020 {
		fmt.Printf("%+v\n\n", dp) // for testing
	}

	_, err = db.Exec(StoreProjectsSQL, dp.Projectid, dp.OAID, dp.Title, dp.Acronym, dp.Summary,
		dp.Keywords, dp.Websiteurl, dp.Callidentifier, dp.Code, dp.Startdate, dp.Enddate,
		dp.GrantedCurr, dp.GrantedAmount, dp.GrantedTotCost, dp.Openaccessdataset, dp.Openaccesspubs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(dfs) > 0 {
		for _, f := range dfs {
			//fmt.Printf("%+v\n\n", f)
			_, err = db.Exec(StoreFundingsSQL, f.Projectid, f.Description, f.StreamID,
				f.Jurisdiction, f.Name, f.ShortName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	if len(dhs) > 0 {
		for _, h := range dhs {
			//fmt.Printf("%+v\n\n", h)
			_, err = db.Exec(StoreH2020sSQL, h.Projectid, h.Code, h.Description)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	if len(dss) > 0 {
		for _, s := range dss {
			//fmt.Printf("%+v\n\n", s)
			_, err = db.Exec(StoreSubjectsSQL, s.Projectid, s.Subject)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

func GetMaxPiD() uint64 {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// close database
	defer db.Close()

	var maxid uint64
	var rowcount int
	SelectCountSQL := `select count(*) as rowcount from openaire.projects`
	SelectMaxSQL := `select max(projectid) as maxid from openaire.projects`
	row := db.QueryRow(SelectCountSQL)
	if err := row.Scan(&rowcount); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if rowcount == 0 {
		return 100000
	} else {
		row := db.QueryRow(SelectMaxSQL)
		switch err := row.Scan(&maxid); err {
		case sql.ErrNoRows:
			return 100000
		case nil:
			return maxid
		default:
			panic(err)
		}
	}
}

func TruncateTables() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// close database
	defer db.Close()

	// do the truncations
	TruncateProjectsSQL := `truncate table openaire.projects`
	TruncateFundingsSQL := `truncate table openaire.projectfundings`
	TruncateH2020sSQL := `truncate table openaire.projecth2020s`
	TruncateSubjectsSQL := `truncate table openaire.projectsubjects`

	db.Exec(TruncateProjectsSQL)
	db.Exec(TruncateFundingsSQL)
	db.Exec(TruncateH2020sSQL)
	db.Exec(TruncateSubjectsSQL)
}

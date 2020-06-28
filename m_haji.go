package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Book struct (Model)
type Mhaji struct {
	HajID            string `json:"HajID"`
	namalengkap      string `json:"namalengkap"`
	jeniskelamin     string `json:"jeniskelamin"`
	noktp            string `json:"noktp"`
	tempatlahir      string `json:"tempatlahir"`
	tgllahir         string `json:"tgllahir"`
	namaibukandung   string `json:"namaibukandung"`
	nopaspor         string `json:"nopaspor"`
	berlakupaspor    string `json:"berlakupaspor"`
	nohp             string `json:"nohp"`
	alamatemail      string `json:"alamatemail"`
	statusperkawinan string `json:"statusperkawinan"`
	alamat           string `json:"alamat"`
	mktime           string `json:"mktime"`
}

// Get all orders

/* func prefetchImages() error {

	cmd := exec.Command("glance-cache-prefetcher")
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("glance-cache-prefetcher failed to execute properly: %v", err)
	}

	return nil
}
*/

func getm_haji(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m_haji []Mhaji

	sql := `SELECT
				HajID,
				IFNULL(namalengkap,''),
				IFNULL(jeniskelamin,'') jeniskelamin,
				IFNULL(noktp,'') noktp,
				IFNULL(tempatlahir,'') tempatlahir,
				IFNULL(tgllahir,'') tgllahir,
				IFNULL(namaibukandung,'') namaibukandung,
				IFNULL(nopaspor,'') nopaspor,
				IFNULL(berlakupaspor,'') berlakupaspor,
				IFNULL(nohp,'') nohp,
				IFNULL(alamatemail,'') alamatemail,
				IFNULL(statusperkawinan,'') statusperkawinan,
				IFNULL(alamat,'') alamat,
				IFNULL(mktime,'') mktime
			FROM m_haji`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var mhaji Mhaji
		err := result.Scan(&mhaji.HajID, &mhaji.namalengkap, &mhaji.jeniskelamin,
			&mhaji.noktp, &mhaji.tempatlahir, &mhaji.tgllahir, &mhaji.namaibukandung,
			&mhaji.nopaspor, &mhaji.berlakupaspor, &mhaji.nohp, &mhaji.alamatemail,
			&mhaji.statusperkawinan, &mhaji.alamat, &mhaji.mktime)

		if err != nil {
			panic(err.Error())
		}
		m_haji = append(m_haji, mhaji)
	}

	json.NewEncoder(w).Encode(m_haji)
}

func createMhaji(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		HajID := r.FormValue("HajID")
		namalengkap := r.FormValue("namalengkap")
		jeniskelamin := r.FormValue("jeniskelamin")
		noktp := r.FormValue("noktp")
		tempatlahir := r.FormValue("tempatlahir")
		tgllahir := r.FormValue("tgllahir")
		namaibukandung := r.FormValue("namaibukandung")
		nopaspor := r.FormValue("nopaspor")
		berlakupaspor := r.FormValue("berlakupaspor")
		nohp := r.FormValue("nohp")
		alamatemail := r.FormValue("alamatemail")
		statusperkawinan := r.FormValue("statusperkawinan")
		alamat := r.FormValue("alamat")
		mktime := r.FormValue("mktime")

		stmt, err := db.Prepare("INSERT INTO m_haji (HajID,namalengkap,jeniskelamin,noktp,tempatlahir,tgllahir,namaibukandung,nopaspor,berlakupaspor,nohp,alamatemail,statusperkawinan,alamat,mktime) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(HajID, namalengkap, jeniskelamin, noktp, tempatlahir, tgllahir, namaibukandung, nopaspor, berlakupaspor, nohp, alamatemail, statusperkawinan, alamat, mktime)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Date Created")
		}

	}
}

func getMhaji(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m_haji []Mhaji
	params := mux.Vars(r)

	sql := `SELECT
				HajID,
				IFNULL(namalengkap,''),
				IFNULL(jeniskelamin,'') jeniskelamin,
				IFNULL(noktp,'') noktp,
				IFNULL(tempatlahir,'') tempatlahir,
				IFNULL(tgllahir,'') tgllahir,
				IFNULL(namaibukandung,'') namaibukandung,
				IFNULL(nopaspor,'') nopaspor,
				IFNULL(berlakupaspor,'') berlakupaspor,
				IFNULL(nohp,'') nohp,
				IFNULL(alamatemail,'') alamatemail,
				IFNULL(statusperkawinan,'') statusperkawinan,
				IFNULL(alamat,'') alamat,
				IFNULL(mktime,'') mktime
			FROM m_haji WHERE HajID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var mhaji Mhaji

	for result.Next() {

		err := result.Scan(&mhaji.HajID, &mhaji.namalengkap, &mhaji.jeniskelamin,
			&mhaji.noktp, &mhaji.tempatlahir, &mhaji.tgllahir, &mhaji.namaibukandung,
			&mhaji.nopaspor, &mhaji.berlakupaspor, &mhaji.nohp, &mhaji.alamatemail,
			&mhaji.statusperkawinan, &mhaji.alamat, &mhaji.mktime)

		if err != nil {
			panic(err.Error())
		}

		m_haji = append(m_haji, mhaji)
	}

	json.NewEncoder(w).Encode(m_haji)
}

func updateMhaji(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newnamalengkap := r.FormValue("namalengkap")

		stmt, err := db.Prepare("UPDATE m_haji SET namalengkap = ? WHERE HajID = ?")

		_, err = stmt.Exec(newnamalengkap, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Mhaji with HajID = %s was updated", params["id"])
	}
}

func deleteMhaji(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM m_haji WHERE HajID = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Mhaji with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var m_haji []Mhaji

	HajID := r.FormValue("hajID")
	namalengkap := r.FormValue("namalengkap")

	sql := `SELECT
				HajID,
				IFNULL(namalengkap,''),
				IFNULL(jeniskelamin,'') jeniskelamin,
				IFNULL(noktp,'') noktp,
				IFNULL(tempatlahir,'') tempatlahir,
				IFNULL(tgllahir,'') tgllahir,
				IFNULL(namaibukandung,'') namaibukandung,
				IFNULL(nopaspor,'') nopaspor,
				IFNULL(berlakupaspor,'') berlakupaspor,
				IFNULL(nohp,'') nohp,
				IFNULL(alamatemail,'') alamatemail,
				IFNULL(statusperkawinan,'') statusperkawinan,
				IFNULL(alamat,'') alamat,
				IFNULL(mktime,'') mktime
			FROM m_haji WHERE HajID = ? AND namalengkap = ?`

	result, err := db.Query(sql, HajID, namalengkap)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var mhaji Mhaji

	for result.Next() {

		err := result.Scan(&mhaji.HajID, &mhaji.namalengkap, &mhaji.jeniskelamin,
			&mhaji.noktp, &mhaji.tempatlahir, &mhaji.tgllahir, &mhaji.namaibukandung,
			&mhaji.nopaspor, &mhaji.berlakupaspor, &mhaji.nohp, &mhaji.alamatemail,
			&mhaji.statusperkawinan, &mhaji.alamat, &mhaji.mktime)

		if err != nil {
			panic(err.Error())
		}

		m_haji = append(m_haji, mhaji)
	}

	json.NewEncoder(w).Encode(m_haji)

}

// Main function (127.0.0.1)
func main() {

	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/m_haji", getm_haji).Methods("GET")
	r.HandleFunc("/m_haji/{id}", getMhaji).Methods("GET")
	r.HandleFunc("/m_haji", createMhaji).Methods("POST")
	r.HandleFunc("/m_haji/{id}", updateMhaji).Methods("PUT")
	r.HandleFunc("/m_haji/{id}", deleteMhaji).Methods("DELETE")

	// new
	r.HandleFunc("/getmhaji", getPost).Methods("POST")
	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))
}

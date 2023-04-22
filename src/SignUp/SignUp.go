package SignUp

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Password string
}

type data struct {
	Message string
}

var db *sql.DB
var err error

func openDB() {
	db, err = sql.Open("mysql", "root:namanh202@tcp(localhost:3306)/webhoctructuyen")
	if err != nil {
		log.Fatal(err)
	}
}

func SignUpHandler() {

	http.HandleFunc("/html/sign-up", func(w http.ResponseWriter, r *http.Request) {
		label := data{
			Message: "",
		}

		if r.Method == "GET" {
			t, err := template.ParseFiles("html/sign-up/sign-up.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.Execute(w, label)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")
			hoten := r.FormValue("hoten")
			role := r.FormValue("role")
			// println(username)

			if checkUserExists(username) {
				label := data{
					Message: "Tài khoản đã tồn tại!",
				}
				// Hiện thông báo
				t, err := template.ParseFiles("html/sign-up/sign-up.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				err = t.Execute(w, label)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				// adduser
				addUser(username, password, hoten, role, w)

				label := data{
					Message: "Đăng ký thành công!",
				}
				//Hiện thông báo
				tmpl, err := template.ParseFiles("html/sign-up/sign-up.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err = tmpl.Execute(w, label)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			// http.Redirect(w, r, "/html/sign-up", http.StatusFound)

		}

	})
}

func checkUserExists(username string) bool {
	openDB()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username=?", username).Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	db.Close()

	if count > 0 {
		return true
	}
	return false
}

func addUser(username string, password string, hoten string, role string, w http.ResponseWriter) {
	openDB()

	_, err = db.Exec("INSERT INTO users (username, password, hoten, role) VALUES (?, ?, ?, ?)", username, password, hoten, role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Close()

}

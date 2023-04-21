package SignUp

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Password string
}

func SignUpHandler() {
	db, err := sql.Open("mysql", "root:namanh202@tcp(localhost:3306)/webhoctructuyen")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/html/sign-up", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.ParseFiles("html/sign-up/sign-up.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			t.Execute(w, nil)
		} else if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")
			hoten := r.FormValue("hoten")
			// println(username)

			_, err = db.Exec("INSERT INTO users (username, password, hoten) VALUES (?, ?, ?)", username, password, hoten)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := struct {
				Message string
			}{
				Message: "Đăng ký thành công!",
			}

			tmpl, err := template.ParseFiles("html/sign-up/sign-up.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
	})

	http.ListenAndServe(":8080", nil)
}

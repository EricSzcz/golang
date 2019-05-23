package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//Usuario
type Usuario struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para a função adequada
func usuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)
	name := strings.TrimPrefix(r.URL.Path, "/usuarios/")

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	case r.Method == "POST":
		insereUsuario(w, r, name)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "Application/json")
	fmt.Fprintf(w, string(json))

}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}
	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "Application/json")
	fmt.Fprintf(w, string(json))
}

func insereUsuario(w http.ResponseWriter, r *http.Request, name string) {
	db, err := sql.Open("mysql", "root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuario Usuario
	usuario.Nome = name
	// decoder := json.NewDecoder(r.Body)
	// err = decoder.Decode(&usuario)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into usuarios(nome) values (?)")
	res, erro := stmt.Exec(usuario.Nome)

	if err != nil {
		tx.Rollback()
		log.Fatal(erro)
	}
	tx.Commit()
	usuario.ID, _ = res.LastInsertId()

	json, _ := json.Marshal(usuario)

	w.Header().Set("Content-Type", "Application/json")
	fmt.Fprintf(w, string(json))

}

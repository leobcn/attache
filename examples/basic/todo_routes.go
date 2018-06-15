package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/mccolljr/attache"
	"github.com/mccolljr/attache/examples/basic/models"
)

func (c *Ctx) GET_TodoNew(r *http.Request) ([]byte, error) {
	return c.Views.Render("todo.create", nil)
}

func (c *Ctx) GET_TodoList(r *http.Request) ([]byte, error) {
	all, err := c.DB.All(func() attache.Storable { return new(models.Todo) })
	if err != nil {
		attache.ErrorFatal(err)
	}

	return c.Views.Render("todo.list", all)
}

func (c *Ctx) GET_Todo(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	target := new(models.Todo)
	if err := c.DB.Find(target, id); err != nil {
		if err == sql.ErrNoRows {
			attache.Error(404)
		}

		attache.ErrorFatal(err)
	}

	data, err := c.Views.Render("todo.update", &target)
	if err != nil {
		attache.ErrorFatal(err)
	}

	w.Header().Set("content-type", "text/html")
	w.Write(data)
}

func (c *Ctx) POST_TodoNew(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		attache.ErrorFatal(err)
	}

	target := new(models.Todo)

	if err := attache.FormDecode(target, r.Form); err != nil {
		attache.ErrorFatal(err)
	}

	if err := c.DB.Insert(target); err != nil {
		attache.ErrorFatal(err)
	}

	attache.RedirectPage(fmt.Sprintf("/todo?id=%d", target.ID))
}

func (c *Ctx) POST_Todo(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	target := new(models.Todo)
	if err := c.DB.Find(target, id); err != nil {
		if err == sql.ErrNoRows {
			attache.Error(404)
		}

		attache.ErrorFatal(err)
	}

	if err := attache.FormDecode(target, r.Form); err != nil {
		attache.ErrorFatal(err)
	}

	if err := c.DB.Update(target); err != nil {
		attache.ErrorFatal(err)
	}

	attache.RedirectPage(fmt.Sprintf("/todo?id=%d", target.ID))
}

func (c *Ctx) DELETE_Todo(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	target := new(models.Todo)
	if err := c.DB.Find(target, id); err != nil {
		if err == sql.ErrNoRows {
			attache.Success()
		}

		attache.ErrorFatal(err)
	}

	if err := c.DB.Delete(target); err != nil {
		attache.ErrorFatal(err)
	}

	w.WriteHeader(200)
}
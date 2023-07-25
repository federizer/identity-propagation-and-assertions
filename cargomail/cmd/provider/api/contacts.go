package api

import (
	"cargomail/cmd/provider/api/helper"
	"cargomail/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

type ContactsApi struct {
	contacts repository.ContactsRepository
}

func (api *ContactsApi) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(repository.UserContextKey).(*repository.User)
		if !ok {
			helper.ReturnErr(w, repository.ErrMissingUserContext, http.StatusInternalServerError)
			return
		}

		var contact *repository.Contact

		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		contact, err = api.contacts.Create(user, contact)
		if err != nil {
			log.Println(err)
			return
		}

		helper.SetJsonResponse(w, http.StatusCreated, contact)
	})
}

func (api *ContactsApi) GetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(repository.UserContextKey).(*repository.User)
		if !ok {
			helper.ReturnErr(w, repository.ErrMissingUserContext, http.StatusInternalServerError)
			return
		}

		contacts, err := api.contacts.GetAll(user)
		if err != nil {
			log.Println(err)
			return
		}

		helper.SetJsonResponse(w, http.StatusCreated, contacts)
	})
}

func (api *ContactsApi) GetHistory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(repository.UserContextKey).(*repository.User)
		if !ok {
			helper.ReturnErr(w, repository.ErrMissingUserContext, http.StatusInternalServerError)
			return
		}

		var history *repository.History

		err := json.NewDecoder(r.Body).Decode(&history)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		contactHistory, err := api.contacts.GetHistory(user, history)
		if err != nil {
			log.Println(err)
			return
		}

		helper.SetJsonResponse(w, http.StatusCreated, contactHistory)
	})
}

func (api *ContactsApi) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(repository.UserContextKey).(*repository.User)
		if !ok {
			helper.ReturnErr(w, repository.ErrMissingUserContext, http.StatusInternalServerError)
			return
		}

		var contact *repository.Contact

		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		contact, err = api.contacts.Update(user, contact)
		if err != nil {
			log.Println(err)
			return
		}

		helper.SetJsonResponse(w, http.StatusCreated, contact)
	})
}

func (api *ContactsApi) TrashByUuidList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(repository.UserContextKey).(*repository.User)
		if !ok {
			helper.ReturnErr(w, repository.ErrMissingUserContext, http.StatusInternalServerError)
			return
		}

		var list []string

		err := json.NewDecoder(r.Body).Decode(&list)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = api.contacts.TrashByUuidList(user, list)
		if err != nil {
			log.Println(err)
			return
		}

		helper.SetJsonResponse(w, http.StatusOK, map[string]string{"status": "OK"})
	})
}
package controller

import (
	"encoding/json"
	"github.com/go-starter/domain"
	"github.com/gorilla/mux"

	"go.uber.org/zap"
	"net/http"
)

type CustomerController struct {
	Store  domain.CustomerStore
	Logger *zap.Logger // Uber's Zap Logger
}

func (ctl CustomerController) Post(w http.ResponseWriter, r *http.Request)  {
	defer ctl.Logger.Sync()
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url",r.URL.String()),
		)
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	// Create customer
	if err:= ctl.Store.Create(customer); err!=nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url",r.URL.String()),
		)
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	ctl.Logger.Info("created note",
		zap.String("url",r.URL.String()),
	)
	w.WriteHeader(http.StatusCreated)
}

func (ctl CustomerController) Put(w http.ResponseWriter, r *http.Request) {
	defer ctl.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil{
		ctl.Logger.Error(err.Error(),
			zap.String("url",r.URL.String()),
		)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	//Update
	if err:= ctl.Store.Update(id,customer); err!=nil{
		ctl.Logger.Error(err.Error(),
			zap.String("customer ID", id),
			zap.String("url",r.URL.String()),
		)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	ctl.Logger.Info("update customer info",
		zap.String("customer ID",id),
		zap.String("url",r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

func (ctl CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	defer ctl.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]

	//delete
	if err:= ctl.Store.Delete(id);err!=nil{
		ctl.Logger.Error(err.Error(),
			zap.String("customer ID",id),
			zap.String("URL",r.URL.String()),
		)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	ctl.Logger.Info("deleted cusomter info",
		zap.String("cutomer ID",id),
		zap.String("URL", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

func (ctl CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	defer ctl.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	//Get by id
	if customer, err := ctl.Store.GetById(id);err!=nil{
		ctl.Logger.Error(err.Error(),
			zap.String("Cutomer ID",id),
			zap.String("URL",r.URL.String()),
		)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}else {
		w.Header().Set("Content-Type","application/json")
		j,err := json.Marshal(customer)
		if err!=nil{
			ctl.Logger.Error(err.Error(),
				zap.String("customer id",id),
				zap.String("url",r.URL.String()),
			)
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func (ctl CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	defer ctl.Logger.Sync()
	if customers, err := ctl.Store.GetAll(); err!=nil{
		ctl.Logger.Error(err.Error(),
			zap.String("url",r.URL.String()),
		)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}else {
		j,err := json.Marshal(customers)
		if err!=nil{
			ctl.Logger.Error(err.Error(),
				zap.String("url",r.URL.String()),
			)
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		w.Header().Set("content-type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
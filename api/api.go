package api

import (
	"encoding/json"
	"net/http"
  "StatusCheck/database"
  // "fmt"
  "github.com/gorilla/mux"
  // "io/ioutil"

)

type API struct {
  Database database.DB
}

type Live struct{
  Live bool
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
  w.Header().Set("Content-Type", "application/json")
  b, err := json.MarshalIndent(data, "", " ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return err
  }
  w.Write(b)
  return nil
}

func (api *API) Live(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Cache-Control", "no-cache")
    url := mux.Vars(r);
    myurl := "https://" + url["url"];
    resp, err := http.Get(myurl)
    live := Live{Live: true}
    if(err != nil){
      live.Live = false;
    }else{
      if(resp.StatusCode != 200){
        live.Live = false;
      }
    }
    WriteJSON(w,live);
  }

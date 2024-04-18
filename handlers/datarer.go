package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abdirahman04/jama/service"
	"github.com/Abdirahman04/jama/types"
)

func Datarer(w http.ResponseWriter, r *http.Request) {
  var data types.DataOpts
  
  if err := json.NewDecoder(r.Body).Decode(&data);err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("error decoding body"))
    return
  }

  newData := service.GetObjs(data)

  jsn, err := json.Marshal(newData)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("error marshalling response"))
    return
  }

  w.Write(jsn)
}

package server

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2/client"
	demoSrv "micro-kit/apps/srv/demo/api/demo"
	"micro-kit/global/reg"
	"net/http"
	"time"
)

func DemoCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	demoSrvMeta := reg.NewSrvDemo()
	// call the backend service
	demoClient := demoSrv.NewDemoService(demoSrvMeta.GetName(), client.DefaultClient)
	rsp, err := demoClient.Call(context.TODO(), &demoSrv.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

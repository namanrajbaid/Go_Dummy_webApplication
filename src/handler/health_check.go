package handler

import (
	"net/http"
	"time"
	"utils"
	log "github.com/sirupsen/logrus"
)

// This struct to store all input parameters
type healthCheckParams struct{
	// Add HTTP Request Object if needed
	HTTPRequest *http.Request `json:"-"`
}

// This function takes the input parameter from request and creates the struct
func NewHealthCheckParams(r *http.Request) (*healthCheckParams, error){
	params := new(healthCheckParams)
	params.HTTPRequest = r
	return params, nil
}

/*
	This function calls New_Params function to get input parameters and after
 	that it calls the Impl function to execute the busniess logic and at the end adds
 	the response with httpreseponseWriter.
	This Function will not have any business logic
*/
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.Debugf("Got request to HealthCheck")

	params, err := NewHealthCheckParams(r)
	if err != nil {
		utils.NewCreateResponseWithParams(http.StatusBadRequest, nil, err.Error()).WriteResponse(w)
		return
	}
	response := HealthCheckImpl(params)

	log.Infof("RE_PERF : time_taken:%v, status_code:%d, error_message:%s",
		time.Now().Sub(startTime).Seconds()*1000, response.StatusCode, response.ErrorMessage)
	response.WriteResponse(w)
}

/*
	This function will have business logic
*/ 
func HealthCheckImpl(params *healthCheckParams) *utils.CreateResponse{
	resp := make(map[string]interface{})
	resp["success"] = true
	return utils.NewCreateResponseWithParams(http.StatusNotFound, resp, "")
}

func init() {
	SubRouter.HandleFunc("/health_check", HealthCheck).Methods(http.MethodGet)
}
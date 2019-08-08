package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Waziup/wazigate-edge/edge"
	"github.com/Waziup/wazigate-edge/tools"
	"github.com/globalsign/mgo/bson"
	routing "github.com/julienschmidt/httprouter"
)

// Actuator represents a Waziup actuator
type Actuator struct {
	ID       string      `json:"id" bson:"id"`
	Name     string      `json:"name" bson:"name"`
	Modified time.Time   `json:"modified" bson:"modified"`
	Created  time.Time   `json:"created" bson:"created"`
	Time     time.Time   `json:"time" bson:"time"`
	Value    interface{} `json:"value" bson:"value"`
}

// GetDeviceActuator implements GET /devices/{deviceID}/actuators/{actuatorID}
func GetDeviceActuator(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	getDeviceActuator(resp, params.ByName("device_id"), params.ByName("actuator_id"))
}

// GetActuator implements GET /actuators/{actuatorID}
func GetActuator(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	getDeviceActuator(resp, edge.LocalID(), params.ByName("actuator_id"))
}

// GetDeviceActuators implements GET /devices/{deviceID}/actuators
func GetDeviceActuators(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	getDeviceActuators(resp, params.ByName("device_id"))
}

// GetActuators implements GET /actuators
func GetActuators(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	getDeviceActuators(resp, edge.LocalID())
}

// PostDeviceActuator implements POST /devices/{deviceID}/actuators
func PostDeviceActuator(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	postDeviceActuator(resp, req, params.ByName("device_id"))
}

// PostActuator implements POST /actuators
func PostActuator(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	postDeviceActuator(resp, req, edge.LocalID())
}

// DeleteDeviceActuator implements DELETE /devices/{deviceID}/actuators/{actuatorID}
func DeleteDeviceActuator(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	deleteDeviceActuator(resp, params.ByName("device_id"), params.ByName("actuator_id"))
}

// DeleteActuator implements DELETE /actuators/{actuatorID}
func DeleteActuator(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	deleteDeviceActuator(resp, edge.LocalID(), params.ByName("actuator_id"))
}

// PostDeviceActuatorName implements POST /devices/{deviceID}/actuators/{actuatorID}/name
func PostDeviceActuatorName(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	postDeviceActuatorName(resp, req, params.ByName("device_id"), params.ByName("actuator_id"))
}

// PostActuatorName implements POST /actuators/{actuatorID}/name
func PostActuatorName(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	postDeviceActuatorName(resp, req, edge.LocalID(), params.ByName("actuator_id"))
}

////////////////////

func getDeviceActuator(resp http.ResponseWriter, deviceID string, actuatorID string) {

	actuator, err := edge.GetActuator(deviceID, actuatorID)
	if err != nil {
		serveError(resp, err)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(actuator)
	resp.Write(data)
}

func getDeviceActuators(resp http.ResponseWriter, deviceID string) {

	device, err := edge.GetDevice(deviceID)
	if err != nil {
		serveError(resp, err)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(device.Actuators)
	resp.Write(data)
}

func postDeviceActuator(resp http.ResponseWriter, req *http.Request, deviceID string) {

	var actuator edge.Actuator
	if err := getReqActuator(req, &actuator); err != nil {
		http.Error(resp, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := edge.PostActuator(deviceID, &actuator); err != nil {
		serveError(resp, err)
		return
	}

	log.Printf("[DB   ] Actuator %s/%s created.\n", deviceID, actuator.ID)

	resp.Header().Set("Content-Type", "application/json")
	resp.Write([]byte{'"'})
	resp.Write([]byte(actuator.ID))
	resp.Write([]byte{'"'})
}

func postDeviceActuatorName(resp http.ResponseWriter, req *http.Request, deviceID string, actuatorID string) {

	body, err := tools.ReadAll(req.Body)
	if err != nil {
		http.Error(resp, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	var name string
	err = json.Unmarshal(body, &name)
	if err != nil {
		http.Error(resp, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = edge.SetActuatorName(deviceID, actuatorID, name)
	if err != nil {
		serveError(resp, err)
		return
	}

	log.Printf("[DB   ] Actuator %s/%s name changed: %q", deviceID, actuatorID, name)
}

func deleteDeviceActuator(resp http.ResponseWriter, deviceID string, actuatorID string) {

	points, err := edge.DeleteActuator(deviceID, actuatorID)
	if err != nil {
		serveError(resp, err)
		return
	}

	log.Printf("[DB   ] Actuator %s/%s removed. (%d values)\n", deviceID, actuatorID, points)
}

////////////////////

func getReqActuator(req *http.Request, actuator *edge.Actuator) error {
	body, err := tools.ReadAll(req.Body)
	if err != nil {
		return err
	}

	now := time.Now()
	actuator.Time = now
	actuator.Modified = now
	actuator.Created = now

	err = json.Unmarshal(body, &actuator)
	if err != nil {
		return err
	}

	if actuator.ID == "" {
		actuator.ID = bson.NewObjectId().String()
	}
	return nil
}

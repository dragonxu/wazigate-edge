package main

import (
	api "github.com/Waziup/waziup-edge/api"
	routing "github.com/julienschmidt/httprouter"
)

var router = routing.New()

func init() {

	// Device Endpoints

	router.GET("/devices", api.GetDevices)
	router.GET("/devices/:device_id", api.GetDevice)
	router.POST("/devices", api.PostDevice)
	router.DELETE("/devices/:device_id", api.DeleteDevice)

	// Sensor Endpoints

	router.GET("/devices/:device_id/sensors", api.GetDeviceSensors)
	router.GET("/devices/:device_id/sensors/:sensor_id", api.GetDeviceSensor)

	router.POST("/devices/:device_id/sensors", api.PostDeviceSensor)
	router.DELETE("/devices/:device_id/sensors/:sensor_id", api.DeleteDeviceSensor)

	router.GET("/devices/:device_id/sensors/:sensor_id/value", api.GetDeviceSensorValue)
	router.GET("/devices/:device_id/sensors/:sensor_id/values", api.GetDeviceSensorValues)

	router.POST("/devices/:device_id/sensors/:sensor_id/value", api.PostDeviceSensorValue)
	router.POST("/devices/:device_id/sensors/:sensor_id/values", api.PostDeviceSensorValues)

	// Actuator Endpoints

	router.GET("/devices/:device_id/actuators", api.GetDeviceActuators)
	router.GET("/devices/:device_id/actuators/:actuator_id", api.GetDeviceActuator)

	router.POST("/devices/:device_id/actuator", api.PostDeviceActuator)
	router.DELETE("/devices/:device_id/actuator/:actuator_id", api.DeleteDeviceActuator)

	router.GET("/devices/:device_id/actuators/:actuator_id/value", api.GetDeviceActuatorValue)
	router.GET("/devices/:device_id/actuators/:actuator_id/values", api.GetDeviceActuatorValues)

	router.POST("/devices/:device_id/actuators/:actuator_id/value", api.PostDeviceActuatorValue)
	router.POST("/devices/:device_id/actuators/:actuator_id/values", api.PostDeviceActuatorValues)

	// Shortcut Endpoints (equals device_id = current device ID)

	router.GET("/device", api.GetCurrentDevice)

	router.GET("/sensors", api.GetSensors)
	router.GET("/sensors/:sensor_id", api.GetSensor)
	router.DELETE("/sensors/:sensor_id", api.DeleteSensor)
	router.GET("/sensors/:sensor_id/value", api.GetSensorValue)
	router.GET("/sensors/:sensor_id/values", api.GetSensorValues)

	router.GET("/actuators", api.GetActuators)
	router.GET("/actuators/:actuator_id", api.GetActuator)
	router.DELETE("/actuators/:actuator_id", api.DeleteActuator)
	router.GET("/actuators/:actuator_id/value", api.GetActuatorValue)
	router.GET("/actuators/:actuator_id/values", api.GetActuatorValues)

	router.POST("/sensors", api.PostSensor)

	router.POST("/sensors/:sensor_id/value", api.PostSensorValue)
	router.POST("/sensors/:sensor_id/values", api.PostSensorValues)

	router.POST("/actuators/:actuator_id/value", api.PostSensorValue)
	router.POST("/actuators/:actuator_id/values", api.PostSensorValues)
}
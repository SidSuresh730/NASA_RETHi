package router

import (
	"../controller"
	"github.com/AmyangXYZ/sweetygo"
)

// SetRouter .
func SetRouter(app *sweetygo.SweetyGo) {
	app.GET("/", controller.Index)
	app.GET("/static/*files", controller.Static)

	app.GET("/api/eclss/wrs", controller.GetWRSData) // ECLSS Water Recovery System
	app.GET("/api/eclss/ogs", controller.GetOGSData) // ECLSS Oxygen Generation System
	app.GET("/api/eclss/fms", controller.GetFMSData) // Food Management System
	app.GET("/api/c2/decisions", controller.GetHMCtrlDecisions) // Health Management System Control Decisions // ERROR
	app.GET("/api/c2/thermflux", controller.GetHMThermFlux) // Health Management Thermal Flux Sensors // ERROR
	app.GET("/api/human/agents", controller.GetHumans) // Human Agents // ERROR
	app.GET("/api/human/setpoints", controller.GetHumanSetPts) // Human Agent Set Point Data
	app.GET("/api/interior/temperature", controller.GetIETemperatureSensors) // Interior Environment Temperature Sensors
	app.GET("/api/interior/humidity", controller.GetIEHumiditySensors) // Interior Environment Humidity Sensors
	app.GET("/api/interior/pressure", controller.GetIEPressureSensors) // Interior Environment Pressure Sensors
	app.GET("/api/structural/temperature", controller.GetSSTemperatureSensors) // Structural System Temperature Sensors
	app.GET("/api/structural/strain", controller.GetSSStrainSensors) // Structural System Strain Sensors
	app.GET("/api/structural/pressure", controller.GetSSPressureSensors) // Structural System Pressure Sensors
	app.GET("/api/structural/acceleration", controller.GetSSAcclSensors) // Structural System Acceleration Sensors
	app.GET("/api/structural/displacement", controller.GetSSDispSensors) // Structural Systm Displacement Sensors
	app.GET("/api/structural/damageinfo", controller.GetSSDamageInfo) // Structural System Damage Information
	app.GET("/api/structural/visualassessment", controller.GetSSCameras) // Structural System Visual Assessment
	app.GET("/api/robot/agents", controller.GetRobots) // Robot Agents
	app.GET("/api/robot/interventions", controller.GetRobotInterventions) // Robot Interventions
	app.GET("/api/power/powergeneration", controller.GetSolarPanels) // Power System Power Generation (Solar Panels)
	app.GET("/api/power/powerconsumption", controller.GetPowerConsumers) // Power System Power Consumption
	app.GET("/api/power/storage", controller.GetBatteries) // Power System Storage (Batteries)
	app.GET("/api/power/health", controller.GetPwrHealthStates) // Power System Component Health States
	app.GET("/api/inventory", controller.GetInventory) // Inventory and Dumping Site
	app.GET("/api/external/disturbances", controller.GetEEDisturbances) // External Environment Disturbances
}

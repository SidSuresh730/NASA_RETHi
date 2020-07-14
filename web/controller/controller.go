package controller

import (
	"net/http"
	//"time"

	"../model"
	"github.com/AmyangXYZ/sweetygo"
)
/*
var lastBootTime int64 = 0

func init() {
	go func() {
		for {
			lastBootTime = .GetLastBootTime()
			time.Sleep(5 * time.Minute)
		}
	}()

}
*/

// Index page handler.
func Index(ctx *sweetygo.Context) error {
	return ctx.Render(200, "index")
}

// Static files handler.
func Static(ctx *sweetygo.Context) error {
	staticHandle := http.StripPrefix("/static",
		http.FileServer(http.Dir("./static")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
	return nil
}

// GetWRSData handles GET /api/eclss/wrs
func GetWRSData(ctx *sweetygo.Context) error {
	wrsList, err := model.GetWRSData()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(wrsList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", wrsList)
}

// GetWRSData handles GET /api/eclss/ogs
func GetOGSData(ctx *sweetygo.Context) error {
	ogsList, err := model.GetOGSData()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(ogsList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", ogsList)
}

// GetWRSData handles GET /api/eclss/fms
func GetFMSData(ctx *sweetygo.Context) error {
	fmsList, err := model.GetFMSData()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(fmsList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", fmsList)
}

// GetHMCtrlDecisions handles GET /api/healthmanagement/controldecisions
func GetHMCtrlDecisions(ctx *sweetygo.Context) error {
	decisionsList, err := model.GetHMCtrlDecisions()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(decisionsList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", decisionsList)
}
/*
// GetHMThermFlux handles GET /api/healthmanagement/thermflux
func GetHMThermFlux(ctx *sweetygo.Context) error {
	tfList, err := model.GetHMThermFlux()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(tfList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", tfList)
}
*/
// GetHumans handles GET /api/humans
func GetHumans(ctx *sweetygo.Context) error {
	humans, err := model.GetHumans()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(humans) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", humans)
}

// GetHumanSetPts handles GET /api/humans/setpoints
func GetHumanSetPts(ctx *sweetygo.Context) error {
	setpts, err := model.GetHumanSetPts()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(setpts) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", setpts)
}

// GetIETemperatureSensors handles GET /api/interiorenvironment/temperature
func GetIETemperatureSensors(ctx *sweetygo.Context) error {
	temps, err := model.GetIETemperatureSensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(temps) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", temps)
}

// GetIEHumiditySensors handles GET /api/interiorenvironment/humidity
func GetIEHumiditySensors(ctx *sweetygo.Context) error {
	hums, err := model.GetIEHumiditySensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(hums) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", hums)
}

// GetIEPressureSensors handles GET /api/interiorenvironment/pressure
func GetIEPressureSensors(ctx *sweetygo.Context) error {
	pressures, err := model.GetIEPressureSensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(pressures) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", pressures)
}

// GetSSPressureSensors handles GET /api/structuralsystem/pressure
func GetSSPressureSensors(ctx *sweetygo.Context) error {
	pressures, err := model.GetSSPressureSensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(pressures) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", pressures)
}

// GetSSTemperatureSensors handles GET /api/structuralsystem/temperature
func GetSSTemperatureSensors(ctx *sweetygo.Context) error {
	temps, err := model.GetSSTemperatureSensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(temps) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", temps)
}

// GetSSStrainSensors handles GET /api/structuralsystem/strain
func GetSSStrainSensors(ctx *sweetygo.Context) error {
	strains, err := model.GetSSAcclSensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(strains) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", strains)
}

// GetSSAcclSensors handles GET /api/structuralsystem/acceleration
func GetSSAcclSensors(ctx *sweetygo.Context) error {
	accls, err := model.GetSSAcclSensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(accls) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", accls)
}

// GetSSDispSensors handles GET /api/structuralsystem/displacement
func GetSSDispSensors(ctx *sweetygo.Context) error {
	disps, err := model.GetSSDispSensors()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(disps) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", disps)
}
/*
// GetSSCameras handles GET /api/structuralsystem/visualassessment
func GetSSCameras(ctx *sweetygo.Context) error {
	cams, err := model.GetSSCameras()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(cams) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", cams)
}
*/
// GetSSDamageInfo handles GET /api/structuralsystem/damageinfo
func GetSSDamageInfo(ctx *sweetygo.Context) error {
	dList, err := model.GetSSDamageInfo()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(dList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", dList)
}

// GetRobots handles GET /api/robots
func GetRobots(ctx *sweetygo.Context) error {
	robots, err := model.GetRobots()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(robots) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", robots)
}

// GetRobotInterventions handles GET /api/robots/interventions
func GetRobotInterventions(ctx *sweetygo.Context) error {
	iList, err := model.GetRobotInterventions()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(iList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", iList)
}

// GetSolarPanels handles GET /api/powersystem/powergeneration
func GetSolarPanels(ctx *sweetygo.Context) error {
	panels, err := model.GetSolarPanels()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(panels) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", panels)
}

// GetPowerConsumers handles GET /api/powersystem/powerconsumption
func GetPowerConsumers(ctx *sweetygo.Context) error {
	cons, err := model.GetPowerConsumers()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(cons) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", cons)
}

// GetBatteries handles GET /api/powersystem/storage
func GetBatteries(ctx *sweetygo.Context) error {
	bList, err := model.GetBatteries()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(bList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", bList)
}

// GetPwrHealthStates handles GET /api/powersystem/health
func GetPwrHealthStates(ctx *sweetygo.Context) error {
	states, err := model.GetPwrHealthStates()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(states) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", states)
}

// GetInventory handles GET /api/inventory
func GetInventory(ctx *sweetygo.Context) error {
	items, err := model.GetInventory()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(items) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", items)
}

// GetEEDisturbances handles GET /api/externalenvironment/disturbances
func GetEEDisturbances(ctx *sweetygo.Context) error {
	dList, err := model.GetEEDisturbances()
	if err != nil {
		return ctx.JSON(500, 0, err.Error(), nil)
	}
	if len(dList) == 0 {
		return ctx.JSON(200, 0, "no result found", nil)
	}
	return ctx.JSON(200, 1, "success", dList)
}

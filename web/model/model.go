package model

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	//"strconv"
	//"strings"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func init() {
	dbAddr := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/mcvt", os.Getenv("DBPasswd"))
	db, _ = sql.Open("mysql", dbAddr)
	for {
		if err := db.Ping(); err != nil {
			fmt.Println(err, ", retry in 10s...")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	// https://github.com/go-sql-driver/mysql/issues/674
	db.SetMaxIdleConns(0)
}
/*
type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Temperature 			int 		`json:"temperature"`
}

type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Pressure 					int 		`json:"pressure"`
}

type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Humidity 					int 		`json:"humidity"`
}

type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Strain 						float64	`json:"strain"`
}

type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Acceleration			struct {
		X float64 `json:"X"`
		Y	float64	`json:"Y"`
		Z float64 `json:"Z"`
	} `json:"acceleration"`
}

type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Displacement			struct {
		X float64 `json:"X"`
		Y	float64	`json:"Y"`
		Z float64 `json:"Z"`
	} `json:"displacement"`
}

type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Visual 						float64	`json:"visual"`
}

type Sensor struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	DamageInfo 						float64	`json:"DamageInfo"`
}
*/

type ECLSSwrs struct {
	UnitID 						string 		`json:"unit_id"`
	LastUpdate 				int64 	`json:"last_update"`
	HealthStatus 			string 	`json:"health_status"`
	OperationStatus 	string	`json:"operation_status"`
	Humidity					float64 `json:"humidity"`
	WaterLevel				float64	`json:"water_level"`
}

type ECLSSogs struct {
	UnitID 						string 		`json:"unit_id"`
	LastUpdate 				int64 	`json:"last_update"`
	HealthStatus 			string 	`json:"health_status"`
	OperationStatus 	string	`json:"operation_status"`
	OxygenLevel 			float64 `json:"oxygen_level"`
	CO2Level					float64	`json:"co2_level"`
}

type ECLSSfms struct {
	UnitID 						string 		`json:"unit_id"`
	LastUpdate 				int64 	`json:"last_update"`
	HealthStatus 			string 	`json:"health_status"`
	OperationStatus 	string	`json:"operation_status"`
	RefrigeratorTemp	float64 `json:"refrigerator_temp"`
	Moisture 					float64	`json:"moisture"`
	Circulation 			float64 `json:"circulation"`
}

type EEDisturbance struct {
	SenderID 					string 		`json:"sender_id"`
	LastUpdate 				int64 	`json:"last_update"`
	ReceiverID 				string 		`json:"receiver_id"`
	SenderPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Disturbance 			string 	`json:"disturbance"`
}

type HMCtrlDecision struct {
	//TimeStamp 				int 		`json:"time_stamp"`
	SenderID 					string	`json:"sender_id"`
	LastUpdate 				int64		`json:"last_update"`
	ReceiverID 				string 	`json:"receiver_id"`
	Command 		 			string 	`json:"command"`
}
/*
type HMThermFlux struct {
	SensorID 					int 		`json:"sensor_id"`
	LastUpdate 				int64		`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`json:"X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	ThermFlux					float64	`json:"thermflux"`
}
*/
type Human struct {
	HumanID 					string 		`json:"human_id"`
	LastUpdate 				int64 	`json:"last_update"`
	HumanStatus 			string 	`json:"human_status"`
	HR 								float64 `json:"heart_rate"`
	BP 								float64	`json:"blood_pressure"`
	O2Sat 						float64 `json:"o2_saturation"`
}

type HumanSetPt struct {
	SenderID 				string 			`json:"sender_id"`
	LastUpdate 			int64 		`json:"last_update"`
	ReceiverID			string 			`json:"receiver_id"`
	SetPoint 				int 			`json:"set_point"`
}

type Sensor struct {
	SensorID 					string 		`json:"sensor_id"`
	LastUpdate 				int64 	`json:"last_update"`
	SensorStatus 			string	`json:"sensor_status"`
	PowerConsumption	float64	`json:"power_consumption"`
	SensorPosition		struct {
		X 	int 	`X"`
		Y		int		`json:"Y"`
		Z		int 	`json:"Z"`
	} `json:"position"`
	Value 						string 	`json:"value"`
}

type InventoryItem struct {
	ItemID 					string			`json:"item_id"`
	LastUpdate      int64 		`json:"last_update"`
	Quantity 				int 			`json:"quantity"`
	RestockOrder 		string 		`json:"restock_order"`
}

type PowerConsumer struct {
	ConsumerID 					string 		`json:"consumer_id"`
	LastUpdate     			int64 	`json:"last_update"`
	PowerConsumption		float64	`json:"power_consumption"`
}

type SolarPanel struct {
	SolarPanelID 					string 		`json:"solar_panel_id"`
	LastUpdate     				int64 	`json:"last_update"`
	Status 								string 	`json:"solar_panel_status"`
	Efficiency 						int 		`json:"solar_panel_efficiency"`
	SolarPanelPosition 		struct {
		X	int `json:"X"`
		Y int `json:"Y"`
		Z int `json:"Z"`
	} `json:"position"`
	PowerGeneration 			float64 `json:"power_generation"`
	HeatGeneration 				float64 `json:"heat_generation"`
}

type HealthState struct {
	OperationUnitID 			string 		`json:"operation_unit_id"`
	LastUpdate     				int64		`json:"last_update"`
	HealthState 					string 	`json:"health_state"`
}

type Battery struct {
	BatteryID 						string 		`json:"battery_id"`
	LastUpdate     				int64 	`json:"last_update"`
	BatteryStatus 				string 	`json:"battery_status"`
	BatteryEfficiency 		int 		`json:"battery_id_efficiency"`
	BatteryPercent 				float64	`json:"battery_percent"`
	Voltage 							float64	`json:"voltage"`
}

type Robot struct {
  RobotID 					string 		`json:"robot_id"`
	LastUpdate     		int64 	`json:"last_update"`
	RobotStatus 			string 	`json:"robot_status"`
	RobotHealth 			string  `json:"robot_health"`
	Battery 					float64	`json:"battery"`
}

type RobotIntervention struct {
  SenderID 					string 		`json:"sender_id"`
	LastUpdate     		int64 	`json:"last_update"`
  ReceiverID 				string 		`json:"receiver_id"`
  Intervention 		  string 	`json:"intervention"`
}

func GetWRSData() ([]ECLSSwrs, error) {
	var unit ECLSSwrs
	var rows *sql.Rows
	uList := make([]ECLSSwrs, 0)
	rows, err := db.Query(`SELECT HEX(t.id), t.timestamp, t.HEALTH_STATUS,
		 										t.OPERATION_STATUS, t.HUMIDITY, t.WATER_LEVEL
												FROM ECLSS_RT_OPERATION
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM ECLSS_RT_OPERATION GROUP BY id) max
												ON t.timestamp = max.timestamp
												WHERE TYPE="WRS"
												ORDER BY t.id `)
	if err != nil {
		return uList, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&unit.UnitID, &unit.LastUpdate, &unit.HealthStatus, &unit.OperationStatus, &unit.Humidity, &unit.WaterLevel)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetOGSData() ([]ECLSSogs, error) {
	var unit ECLSSogs
	uList := make([]ECLSSogs, 0)

	rows, err := db.Query(`SELECT HEX(t.id), t.timestamp, t.HEALTH_STATUS,
		 										t.OPERATION_STATUS, t.OXYGEN_LEVEL, t.CO2_LEVEL
												FROM ECLSS_RT_OPERATION
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM ECLSS_RT_OPERATION GROUP BY id) max
												ON t.timestamp = max.timestamp
												WHERE TYPE="OGS"
												ORDER BY t.id `)
	if err != nil {
		return uList, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&unit.UnitID, &unit.LastUpdate, &unit.HealthStatus, &unit.OperationStatus, &unit.OxygenLevel, &unit.CO2Level)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetFMSData() ([]ECLSSfms, error) {
	var unit ECLSSfms
	uList := make([]ECLSSfms, 0)

	rows, err := db.Query(`SELECT HEX(t.id), t.timestamp, t.HEALTH_STATUS,
		 										t.OPERATION_STATUS, t.REFRIGERATOR_TEMP, t.MOISTURE, t.CIRCULATION
												FROM ECLSS_RT_OPERATION
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM ECLSS_RT_OPERATION GROUP BY id) max
												ON t.timestamp = max.timestamp
												WHERE TYPE="FMS"
												ORDER BY t.id `)
	if err != nil {
		return uList, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&unit.UnitID, &unit.LastUpdate, &unit.HealthStatus, &unit.OperationStatus, &unit.RefrigeratorTemp, &unit.Moisture, &unit.Circulation)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetEEDisturbances() ([]EEDisturbance, error) {
	var unit EEDisturbance
	uList := make([]EEDisturbance, 0)
	rows, err := db.Query(`SELECT HEX(t.sender_id), t.timestamp, HEX(t.receiver_id), t.sender_x,
												t.sender_y, t.sender_z, t.disturbance
												FROM ee_disturbance_detection
												t JOIN(SELECT MAX(timestamp) AS timestamp FROM ee_disturbance_detection
												GROUP BY sender_id) max ON t.timestamp = max.timestamp
												ORDER BY t.sender_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&unit.SenderID, &unit.LastUpdate, &unit.ReceiverID, &unit.SenderPosition.X, &unit.SenderPosition.Y, &unit.SenderPosition.Z, &unit.Disturbance)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetHMCtrlDecisions() ([]HMCtrlDecision, error) {
	var unit HMCtrlDecision
	uList := make([]HMCtrlDecision, 0)
	rows, err := db.Query(`SELECT HEX(t.sender_id), t.timestamp, HEX(t.receiver_id),
		 										t.command
												FROM HM_CTRLDECSN
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM HM_CTRLDECSN GROUP BY sender_id) max
												ON t.timestamp = max.timestamp
												ORDER BY t.sender_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SenderID, &unit.LastUpdate, &unit.ReceiverID, &unit.Command)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetHMThermFlux() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.sensor_id), t.timestamp, t.sensor_status,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.value FROM HM_RT_THERMFLUX
												t JOIN (SELECT MAX(timestamp) as timestamp
												FROM HM_RT_THERMFLUX GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
			&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetHumans() ([]Human, error) {
	var unit Human
	uList := make([]Human, 0)
	rows, err := db.Query(`SELECT HEX(t.human_id), t.timestamp, t.HUMAN_STATUS,
		 										t.HR, t.BP, t.O2_SAT FROM HUMAN_RT_UPDATES
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM HUMAN_RT_UPDATES GROUP BY human_id) max
												ON t.timestamp = max.timestamp
												ORDER BY t.human_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.HumanID, &unit.LastUpdate, &unit.HumanStatus, &unit.HR, &unit.BP, &unit.O2Sat)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetHumanSetPts() ([]HumanSetPt, error) {
	var unit HumanSetPt
	uList := make([]HumanSetPt, 0)
	rows, err := db.Query(`SELECT HEX(t.sender_id), t.timestamp, HEX(t.RECEIVER_ID),
		 										t.SET_POINT FROM HUMAN_SETPTCTRL
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM HUMAN_SETPTCTRL GROUP BY sender_id) max
												ON t.timestamp = max.timestamp
												ORDER BY t.sender_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SenderID, &unit.LastUpdate, &unit.ReceiverID, &unit.SetPoint)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetIETemperatureSensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM IE_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM IE_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "TEMP"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSSTemperatureSensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM SS_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM SS_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "TEMP"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetIEPressureSensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM IE_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM IE_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "PRESSURE"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSSPressureSensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM SS_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM SS_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "PRESSURE"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSSStrainSensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM SS_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM SS_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "Strain"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetIEHumiditySensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM IE_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM IE_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "HUMIDITY"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSSAcclSensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM SS_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM SS_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "ACCL"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSSDispSensors() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM SS_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM SS_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "DISP"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSSCameras() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM SS_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM SS_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "VISUAL"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSSDamageInfo() ([]Sensor, error) {
	var unit Sensor
	uList := make([]Sensor, 0)
	rows, err := db.Query(`SELECT HEX(t.SENSOR_ID), t.timestamp, t.SENSOR_STATUS,
		 										t.SENSOR_POWER_CONSUME, t.SENSOR_X, t.SENSOR_Y, t.SENSOR_Z,
												t.VALUE FROM SS_RT_SENSORS
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM SS_RT_SENSORS GROUP BY sensor_id) max
												ON t.timestamp = max.timestamp
												WHERE SENSOR_TYPE = "DAMAGE"
												ORDER BY t.sensor_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SensorID, &unit.LastUpdate, &unit.SensorStatus, &unit.PowerConsumption, &unit.SensorPosition.X,
		&unit.SensorPosition.Y, &unit.SensorPosition.Z, &unit.Value)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetInventory() ([]InventoryItem, error) {
	var unit InventoryItem
	uList := make([]InventoryItem, 0)
	rows, err := db.Query(`SELECT HEX(t.item_id), t.timestamp, t.QUANTITY,
		 										t.RESTOCK_ORDER FROM INVENTORY
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM INVENTORY GROUP BY item_id) max
												ON t.timestamp = max.timestamp
												ORDER BY t.ITEM_ID`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.ItemID, &unit.LastUpdate, &unit.Quantity, &unit.RestockOrder)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetPowerConsumers() ([]PowerConsumer, error) {
	var unit PowerConsumer
	uList := make([]PowerConsumer, 0)
	rows, err := db.Query(`SELECT HEX(t.CONSUMER_ID), t.timestamp, t.POWER_CONS
												FROM  PWR_CONSUMPTION
												t JOIN (SELECT MAX(timestamp) as timestamp
												FROM PWR_CONSUMPTION GROUP BY CONSUMER_ID) max
												ON t.timestamp = max.timestamp
												ORDER BY t.consumer_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.ConsumerID, &unit.LastUpdate, &unit.PowerConsumption)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetSolarPanels() ([]SolarPanel, error) {
	var unit SolarPanel
	uList := make([]SolarPanel, 0)
	rows, err := db.Query(`SELECT HEX(t.SOLAR_PANEL_ID), t.timestamp, t.SOLAR_PANEL_STATUS,
		 										t.SOLAR_PANEL_EFFICIENCY, t.PANEL_X, t.PANEL_Y, t.PANEL_Z,
												t.POWER_GEN, t.HEAT_GEN FROM PWR_GENERATION
												t JOIN (SELECT MAX(timestamp) as timestamp
												FROM PWR_GENERATION GROUP BY solar_panel_id) max
												ON t.timestamp = max.timestamp
												ORDER BY t.solar_panel_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SolarPanelID, &unit.LastUpdate, &unit.Status, &unit.Efficiency, &unit.SolarPanelPosition.X,
		&unit.SolarPanelPosition.Y, &unit.SolarPanelPosition.Z, &unit.PowerGeneration, &unit.HeatGeneration)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetPwrHealthStates() ([]HealthState, error) {
	var unit HealthState
	uList := make([]HealthState, 0)
	rows, err := db.Query(`SELECT HEX(t.OPERATIONUNIT_ID), t.timestamp, t.OPERATIONUNIT_HEALTHSTATE
		 										FROM PWR_HEALTHSTATE
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM PWR_HEALTHSTATE GROUP BY OPERATIONUNIT_ID) max
												ON t.timestamp = max.timestamp
												ORDER BY t.OPERATIONUNIT_ID`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.OperationUnitID, &unit.LastUpdate, &unit.HealthState)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetBatteries() ([]Battery, error) {
	var unit Battery
	uList := make([]Battery, 0)
	rows, err := db.Query(`SELECT HEX(t.BATTERY_ID), t.timestamp, t.BATTERY_STATUS,
		 										t.BATTERY_EFFICIENCY, t.BATTERY_PERCENT, t.VOLTAGE
												FROM PWR_STORAGE
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM PWR_STORAGE GROUP BY BATTERY_ID) max
												ON t.timestamp = max.timestamp
												ORDER BY t.battery_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.BatteryID, &unit.LastUpdate, &unit.BatteryStatus, &unit.BatteryEfficiency, &unit.BatteryPercent, &unit.Voltage)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetRobots() ([]Robot, error) {
	var unit Robot
	uList := make([]Robot, 0)
	rows, err := db.Query(`SELECT HEX(t.ROBOT_ID), t.timestamp, t.ROBOT_STATUS,
		 										t.ROBOT_HEALTH, t.BATTERY
												FROM ROBOT_RT_UPDATES
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM ROBOT_RT_UPDATES GROUP BY ROBOT_ID) max
												ON t.timestamp = max.timestamp
												ORDER BY t.ROBOT_ID`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.RobotID, &unit.LastUpdate, &unit.RobotStatus, &unit.RobotHealth, &unit.Battery)
		uList = append(uList, unit)
	}
	return uList, nil
}

func GetRobotInterventions() ([]RobotIntervention, error) {
	var unit RobotIntervention
	uList := make([]RobotIntervention, 0)
	rows, err := db.Query(`SELECT HEX(t.SENDER_ID), t.timestamp, t.RECEIVER_ID,
		 										t.INTERVENTION
												FROM ROBOT_INTERVENTIONCTRL
												t JOIN(SELECT MAX(timestamp) as timestamp
												FROM ROBOT_INTERVENTIONCTRL GROUP BY SENDER_ID) max
												ON t.timestamp = max.timestamp
												ORDER BY t.sender_id`)
	if err != nil {
		return uList, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unit.SenderID, &unit.LastUpdate, &unit.ReceiverID, &unit.Intervention)
		uList = append(uList, unit)
	}
	return uList, nil
}

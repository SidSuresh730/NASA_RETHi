package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	"os"
	"strconv"
	"time"
  "math/rand"
  "strings"
  //"../web/model"
	_ "github.com/go-sql-driver/mysql"
)

var (
  db *sql.DB
  err error
  Info        *log.Logger
  Error       *log.Logger
  infoHandle  = os.Stdout
  errorHandle = os.Stdout
  eclss_type = [3]string{"WRS", "OGS", "FMS"}
  eclss_health = [2]string{"Healthy", "Damaged"}
  eclss_operation = [3]string{"Active", "Standby", "Sleep"}
  ee_disturbance = [6]string{"Moonquake", "Temp_fluc", "meteoroid_impact", "dust_potential", "dust_storm", "Solar_radiation"}
  hm_command = [3]string{"Command0", "Command1", "Command2"}
  sensor_status_list = [4]string{"Replace", "Poor", "Good", "Excellent"}
  human_status = [4]string{"Normal", "Panic", "Unconscious", "Deceased"}
  ie_sensor_type = [3]string{"Temp", "Pressure", "Humidity"}
  restock_order = [2]string{"Sent", "Not Sent"}
  inventory_item_list = [10]int{0x1111, 0x2222, 0x3333, 0x4444, 0x5555, 0x6666, 0x7777, 0x8888, 0x9999, 0xAAAA}
  solar_panel_status = [4]string{"Sleep", "Half", "Full", "Damaged"}
  battery_status = [2]string{"Healthy", "Damaged"}
  health_state = [3]string{"Healthy", "Damaged", "Low_power"}
  intervention_list = [3]string{"In0", "In1", "In2"}
  robot_status = [3]string{"Active", "Standby", "Sleep"}
  robot_health = [3]string{"Full Func", "Some Damage", "Critical"}
  subsystem_id_list = [9]string{"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "A9"}
  component_id_list = [8]string{"11", "22", "33", "44", "55", "66", "77", "88"}
  ss_sensor_type = [8]string{"Strain", "Accl", "Temp", "Disp", /*"Visual",*/ "Damage", "ThermFlux", "Pressure"}
  id_list = populateList()
  //id_list = [21]int{0x1111, 0x2121, 0x3434, 0x4444, 0x5511, 0x1234, 0x12AA, 0x55AA, 0x567A, 0x5999, 0x9999, 0x5978, 0x7777, 0x8888, 0xAAAB, 0xACAB, 0xCCCC, 0xDEAD, 0xBEEF, 0xCADE, 0xCAFE}
)
func main() {
  //id_list = populateList80()
  generateSSData()
   //go generateECLSS()
   //generateDisturbance()
  //go generateCtrlDecsn()
  //go generateHuman()
  //generateHumanSetPt()
  //generateIEData()
  //generateInventory()
  //go generatePwrConsumers()
  ///generatePwrGenerators()
  //generatePwrHealthStates()
  //generatePwrStorage()
  //generateRobotInterventions()
  //generateRobots()

}

// init db
func init() {
  Info = log.New(infoHandle, "[*] INFO: ", log.Ldate|log.Ltime)
	Error = log.New(errorHandle, "[!] ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	dbAddr := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/mcvt", os.Getenv("DBPasswd"))
	db, _ = sql.Open("mysql", dbAddr)
	for {
		if err := db.Ping(); err != nil {
			Error.Println(err, ", retry in 10s...")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}
	Info.Println("connected to db")
	// https://github.com/go-sql-driver/mysql/issues/674
	db.SetMaxIdleConns(0)

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)
}

func populateList() []uint64 {
  lst := make([]uint64, 0)
  for i:=0; i<16; i++ {
    lst = append(lst, uint64(i*16))
  }
  return lst
}

func generateECLSS() {
  _, err := db.Exec("DELETE FROM ECLSS_RT_OPERATION WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i < 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    id := id_list[rand.Intn(len(id_list))]
    //subsys := eclss_type[int(float64(id)/float64(0xFF)*float64(len(eclss_type)))]
    subsys := eclss_type[rand.Intn(len(eclss_type))]
    switch (subsys) {
      case "WRS":
        total_id := strings.ToUpper("11-" + strconv.FormatUint(id, 16))
        stmt, err := db.Prepare(`INSERT INTO ECLSS_RT_OPERATION (DATETIME, TIMESTAMP, ID, TYPE,
                                HEALTH_STATUS, OPERATION_STATUS, POWER_CONSUME, HUMIDITY,
                                WATER_LEVEL) VALUES(?,?,?,?,?,?,?,?,?)`)
        if err != nil {
      		Error.Println(stmt, err)
      	}
      	defer stmt.Close()
        _, err = stmt.Exec(datetime, timestamp, total_id, subsys,
        eclss_health[rand.Intn(len(eclss_health))], eclss_operation[rand.Intn(len(eclss_operation))], 500.0*(rand.Float64()),
        500.0*(rand.Float64()), 500.0*(rand.Float64()))
        if err != nil {
      		Error.Println(stmt, err)
      	}
      case "OGS":
        total_id := strings.ToUpper("22-" + strconv.FormatUint(id, 16))
        stmt, err := db.Prepare(`INSERT INTO ECLSS_RT_OPERATION (DATETIME, TIMESTAMP, ID, TYPE,
                                HEALTH_STATUS, OPERATION_STATUS, POWER_CONSUME, OXYGEN_LEVEL,
                                CO2_LEVEL) VALUES(?,?,?,?,?,?,?,?,?)`)
        if err != nil {
      		Error.Println(stmt, err)
      	}
      	defer stmt.Close()
        _, err = stmt.Exec(datetime, timestamp, total_id, subsys,
        eclss_health[rand.Intn(len(eclss_health))], eclss_operation[rand.Intn(len(eclss_operation))], 500.0*(rand.Float64()),
        500.0*(rand.Float64()), 500.0*(rand.Float64()))
        if err != nil {
      		Error.Println(stmt, err)
      	}
      case "FMS":
          total_id := strings.ToUpper("33-" + strconv.FormatUint(id, 16))
          stmt, err := db.Prepare(`INSERT INTO ECLSS_RT_OPERATION (DATETIME, TIMESTAMP, ID, TYPE,
                                  HEALTH_STATUS, OPERATION_STATUS, POWER_CONSUME, REFRIGERATOR_TEMP,
                                  MOISTURE, CIRCULATION) VALUES(?,?,?,?,?,?,?,?,?,?)`)
          if err != nil {
        		Error.Println(stmt, err)
        	}
        	defer stmt.Close()
          _, err = stmt.Exec(datetime, timestamp, total_id, subsys,
          eclss_health[rand.Intn(len(eclss_health))], eclss_operation[rand.Intn(len(eclss_operation))], 500.0*(rand.Float64()),
          500.0*(rand.Float64()), 500.0*(rand.Float64()), 500.0*(rand.Float64()))
          if err != nil {
        		Error.Println(stmt, err)
        	}
      default:
        total_id := strings.ToUpper("A3-33" + strconv.FormatUint(id, 16))
        stmt, err := db.Prepare(`INSERT INTO ECLSS_RT_OPERATION (DATETIME, TIMESTAMP, ID, TYPE, HEALTH_STATUS, OPERATION_STATUS, REFRIGERATOR_TEMP, MOISTURE, CIRCULATION) VALUES(?,?,?,?,?,?,?,?,?)`)
        if err != nil {
          Error.Println(stmt, err)
        }
        defer stmt.Close()
        _, err = stmt.Exec(datetime, timestamp, total_id, "FMS",
        eclss_health[rand.Intn(len(eclss_health))], eclss_operation[rand.Intn(len(eclss_operation))],
        500.0*(rand.Float64()), 500.0*(rand.Float64()), 500.0*(rand.Float64()))
        if err != nil {
          Error.Println(stmt, err)
        }
    }
  }
}

func generateDisturbance() {
  _, err := db.Exec("DELETE FROM EE_DISTURBANCE_DETECTION WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i< 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    sender_id := strings.ToUpper(subsystem_id_list[rand.Intn(len(subsystem_id_list))]+"-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    receiver_id := strings.ToUpper(subsystem_id_list[rand.Intn(len(subsystem_id_list))]+"-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    stmt, err := db.Prepare(`INSERT INTO EE_DISTURBANCE_DETECTION (DATETIME, TIMESTAMP, SENDER_ID, RECEIVER_ID, SENDER_X, SENDER_Y, SENDER_Z, DISTURBANCE) VALUES(?,?,?,?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, sender_id, receiver_id,
    50*(rand.Float64()), 50*(rand.Float64()), 50*(rand.Float64()),
     ee_disturbance[rand.Intn(len(ee_disturbance))])
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generateCtrlDecsn() {
  _, err := db.Exec("DELETE FROM HM_CTRLDECSN WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i< 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    sender_id := strings.ToUpper(subsystem_id_list[rand.Intn(len(subsystem_id_list))]+"-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    receiver_id := strings.ToUpper(subsystem_id_list[rand.Intn(len(subsystem_id_list))]+"-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    stmt, err := db.Prepare(`INSERT INTO HM_CTRLDECSN (DATETIME, TIMESTAMP, SENDER_ID, RECEIVER_ID,  COMMAND) VALUES(?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, sender_id, receiver_id,
     hm_command[rand.Intn(len(hm_command))])
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}
/*
func generateThermflux() {
  _, err := db.Exec("DELETE FROM HM_RT_THERMFLUX WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i< 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    stmt, err := db.Prepare(`INSERT INTO HM_RT_THERMFLUX (DATETIME, TIMESTAMP, SENSOR_ID, SENSOR_STATUS, SENSOR_POWER_CONSUME, SENSOR_X, SENSOR_Y, SENSOR_Z, VALUE) VALUES(?,?,?,?,?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, id_list[rand.Intn(len(id_list))], sensor_status_list[rand.Intn(len(sensor_status_list))],
    500.0*rand.Float64(), 50*(rand.Float64()), 50*(rand.Float64()), 50*(rand.Float64()),
     strconv.Itoa(int(500*(rand.Float64()))))
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}
*/
func generateHuman() {
  _, err := db.Exec("DELETE FROM HUMAN_RT_UPDATES WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i< 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    human_id := strings.ToUpper(component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    //receiver_id := strings.ToUpper(subsystem_id_list[rand.Intn(len(subsystem_id_list))]+"-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    stmt, err := db.Prepare(`INSERT INTO HUMAN_RT_UPDATES (DATETIME, TIMESTAMP, HUMAN_ID, HUMAN_STATUS, HR, BP, O2_SAT) VALUES(?,?,?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, human_id, human_status[rand.Intn(len(human_status))],
    100.0*(rand.Float64()), 100.0*(rand.Float64()), 100.0*(rand.Float64()))
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generateHumanSetPt() {
  _, err := db.Exec("DELETE FROM HUMAN_SETPTCTRL WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i< 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    human_id := strings.ToUpper("A7-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    receiver_id := strings.ToUpper(subsystem_id_list[rand.Intn(len(subsystem_id_list))]+"-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    stmt, err := db.Prepare(`INSERT INTO HUMAN_SETPTCTRL (DATETIME, TIMESTAMP, SENDER_ID, RECEIVER_ID,  SET_POINT) VALUES(?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, human_id, receiver_id,
     50*rand.Float64())
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generateIEData() {
  _, err := db.Exec("DELETE FROM IE_RT_SENSORS WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i < 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    id := strings.ToUpper(component_id_list[rand.Intn(3)]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    char, e := strconv.Atoi(id[0:1])
    if e != nil {
      Error.Println(char, e)
    }
    subsys := ie_sensor_type[char%3]
      stmt, err := db.Prepare(`INSERT INTO IE_RT_SENSORS (DATETIME, TIMESTAMP, SENSOR_ID, SENSOR_STATUS, SENSOR_POWER_CONSUME, SENSOR_X, SENSOR_Y, SENSOR_Z,
        SENSOR_TYPE, VALUE) VALUES(?,?,?,?,?,?,?,?,?,?)`)
      if err != nil {
    		Error.Println(stmt, err)
    	}
    	defer stmt.Close()
      _, err = stmt.Exec(datetime, timestamp, id, sensor_status_list[rand.Intn(len(sensor_status_list))],
      1000.00*rand.Float64(), 50*rand.Float64(), 50*rand.Float64(), 50*rand.Float64(), subsys,
      strconv.Itoa(int(200*(rand.Float64()))))
      if err != nil {
    		Error.Println(stmt, err)
    	}
    }
}

func generateInventory() {
  _, err := db.Exec("DELETE FROM INVENTORY WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i<1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    item_id := strings.ToUpper(component_id_list[rand.Intn(3)]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    quantity := 50*rand.Float64()
    restock := restock_order[rand.Intn(len(restock_order))]
    stmt, err := db.Prepare(`INSERT INTO INVENTORY (DATETIME, TIMESTAMP, ITEM_ID, QUANTITY, RESTOCK_ORDER) VALUES(?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, item_id, quantity, restock)
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generatePwrConsumers() {
  _, err := db.Exec("DELETE FROM PWR_CONSUMPTION WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i<1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    item_id := strings.ToUpper("44-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    power_consumption := 50*rand.Float64()
    stmt, err := db.Prepare(`INSERT INTO PWR_CONSUMPTION (DATETIME, TIMESTAMP, CONSUMER_ID, POWER_CONS) VALUES(?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, item_id, power_consumption)
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generatePwrGenerators() {
  _, err := db.Exec("DELETE FROM PWR_GENERATION WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i<1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    item_id := strings.ToUpper("11-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))

    status := solar_panel_status[rand.Intn(len(solar_panel_status))]
    power_gen := 50.0*rand.Float64()
    heat_gen := 50.0*rand.Float64()
    efficiency := power_gen/(power_gen+heat_gen)*100.0
    panel_x := 50*rand.Float64()
    panel_y := 50*rand.Float64()
    panel_z := 50*rand.Float64()
    stmt, err := db.Prepare(`INSERT INTO PWR_GENERATION (DATETIME, TIMESTAMP, SOLAR_PANEL_ID, SOLAR_PANEL_STATUS, SOLAR_PANEL_EFFICIENCY, PANEL_X,
      PANEL_Y, PANEL_Z, POWER_GEN, HEAT_GEN) VALUES(?,?,?,?,?,?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, item_id, status, efficiency, panel_x, panel_y, panel_z, power_gen, heat_gen)
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generatePwrHealthStates() {
  _, err := db.Exec("DELETE FROM PWR_HEALTHSTATE WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i<1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    item_id := strings.ToUpper("22-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    status := health_state[rand.Intn(len(health_state))]
    stmt, err := db.Prepare(`INSERT INTO PWR_HEALTHSTATE (DATETIME, TIMESTAMP, OPERATIONUNIT_ID, OPERATIONUNIT_HEALTHSTATE) VALUES(?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, item_id, status)
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generatePwrStorage() {
  _, err := db.Exec("DELETE FROM PWR_STORAGE WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i<1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    item_id := strings.ToUpper("33-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    status := battery_status[rand.Intn(len(battery_status))]
    efficiency := 100*rand.Float64()
    percentage := 100.0*rand.Float64()
    voltage := 100.0*rand.Float64()
    stmt, err := db.Prepare(`INSERT INTO PWR_STORAGE (DATETIME, TIMESTAMP, BATTERY_ID, BATTERY_STATUS, BATTERY_EFFICIENCY, BATTERY_PERCENT, VOLTAGE) VALUES(?,?,?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, item_id, status, efficiency, percentage, voltage)
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generateRobotInterventions() {
  _, err := db.Exec("DELETE FROM ROBOT_INTERVENTIONCTRL WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i<1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    sender_id := strings.ToUpper("A5-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    receiver_id := strings.ToUpper(subsystem_id_list[rand.Intn(len(subsystem_id_list))]+"-"+component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    intervention := intervention_list[rand.Intn(len(intervention_list))]
    stmt, err := db.Prepare(`INSERT INTO ROBOT_INTERVENTIONCTRL (DATETIME, TIMESTAMP, SENDER_ID, RECEIVER_ID, INTERVENTION) VALUES(?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, sender_id, receiver_id, intervention)
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generateRobots() {
  _, err := db.Exec("DELETE FROM ROBOT_RT_UPDATES WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i<1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    id := strings.ToUpper(component_id_list[rand.Intn(len(component_id_list))]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    status := robot_status[rand.Intn(len(robot_status))]
    health := robot_health[rand.Intn(len(robot_health))]
    battery := 100.0*rand.Float64()
    stmt, err := db.Prepare(`INSERT INTO ROBOT_RT_UPDATES (DATETIME, TIMESTAMP, ROBOT_ID, ROBOT_STATUS, ROBOT_HEALTH, BATTERY) VALUES(?,?,?,?,?,?)`)
    if err != nil {
      Error.Println(stmt, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(datetime, timestamp, id, status, health, battery)
    if err != nil {
      Error.Println(stmt, err)
    }
  }
}

func generateSSData() {
  _, err := db.Exec("DELETE FROM SS_RT_SENSORS WHERE 1")
  if err != nil {
    Error.Println(err)
  }
  for i := 0; i < 1000; i++ {
    t := time.Now()
    timestamp := t.UnixNano() / 1e6
    datetime := t.Truncate(time.Second).Local()
    //id_list := populateList80()
    sensor_id := strings.ToUpper(component_id_list[rand.Intn(7)]+"-"+strconv.FormatUint(id_list[rand.Intn(len(id_list))],16))
    char, e := strconv.Atoi(sensor_id[0:1])
    if e != nil {
      Error.Println(char, e)
    }
    subsys := ss_sensor_type[char%7]
    //subsys := ss_sensor_type[rand.Intn(len(ss_sensor_type))]
    status := sensor_status_list[rand.Intn(len(sensor_status_list))]
    consume := 1000.00*rand.Float64()
    sensor_x, sensor_y, sensor_z := 50*rand.Float64(), 50*rand.Float64(), 50*rand.Float64()
    value3 := strconv.Itoa(int(50*rand.Float64())) + "," + strconv.Itoa(int(50*rand.Float64())) + "," + strconv.Itoa(int(50*rand.Float64()))
    value :=strconv.Itoa(int(500*rand.Float64()))
    if subsys == "Accl" || subsys == "Disp" {
      stmt, err := db.Prepare(`INSERT INTO SS_RT_SENSORS (DATETIME, TIMESTAMP, SENSOR_ID, SENSOR_STATUS, SENSOR_POWER_CONSUME, SENSOR_X, SENSOR_Y, SENSOR_Z, SENSOR_TYPE, VALUE) VALUES(?,?,?,?,?,?,?,?,?,?)`)
      if err != nil {
        Error.Println(stmt, err)
      }
      defer stmt.Close()
      _, err = stmt.Exec(datetime, timestamp, sensor_id, status, consume, sensor_x, sensor_y, sensor_z, subsys, value3)
      if err != nil {
        Error.Println(stmt, err)
    }
    } else {
        stmt, err := db.Prepare(`INSERT INTO SS_RT_SENSORS (DATETIME, TIMESTAMP, SENSOR_ID, SENSOR_STATUS, SENSOR_POWER_CONSUME, SENSOR_X, SENSOR_Y, SENSOR_Z,
          SENSOR_TYPE, VALUE) VALUES(?,?,?,?,?,?,?,?,?,?)`)
        if err != nil {
          Error.Println(stmt, err)
        }
        defer stmt.Close()
        _, err = stmt.Exec(datetime, timestamp, sensor_id, status, consume, sensor_x, sensor_y, sensor_z, subsys, value)
        if err != nil {
          Error.Println(stmt, err)
        }
      }
  }
}

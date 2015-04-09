package rest

import (
  "github.com/jmcvetta/napping"
  "log"
//  "time"
)

var (
  PathStatus = "status"
  PathInfo = "info"
  PathSwitch = "switch"
  PathReset = "reset"
  PathFactoryReset = "factoryreset"
  PathConfig = "config"
)

type Status struct {
  Status string	`json:"status"`
}

type SwStatus struct {
  Val int	`json:"val"`
}

type DevInfo struct {
  Wifimode  string  `json:"wifimode"`
  Wifissid  string  `json:"wifissid"`
  Wifipwd   string  `json:"wifipwd"`
  Wifiip    string  `json:"wifiip"`
  Wifinetmask string  `json:"wifinetmask"`
  Wifigateway string  `json:"wifigateway"`
  Device  string  `json:"device"`
  Land	  string  `json:"land"`
  Serial  string  `json:"serial"`
  Mac	  string  `json:"mac"`
  Ip	  string  `json:"ip"`
  Landhq  string  `json:"landhq"`
}

func GetStatus(url string, res Status) (Status,error) {
  resp, err := napping.Get(url, nil, &res, nil)
  if resp.Status() == 200 {
    log.Println("getStatus success")
  }
  return res, err
}

func GetSwitch(url string, res SwStatus) (SwStatus, error) {
//  res := SwStatus{}
  resp, err := napping.Get(url, nil, &res, nil)
  if resp.Status() == 200 {
    log.Println("getSwitch success")
  }
  return res, err
}

func GetInfo(url string, res DevInfo) (DevInfo, error) {
//  res := DevInfo{}
  resp, err := napping.Get(url, nil, &res, nil)
  if resp.Status() == 200 {
    log.Println("getInfo success")
  }
  return res, err
}
/*
func main() {
  resGetStatus := Status{}
  resGetSwitch := SwStatus{}
  resGetInfo := DevInfo{}
  var err error
  baseUrl := "http://192.168.0.15/"

  resGetSwitch, err = getSwitch(baseUrl+PathSwitch, resGetSwitch)
  if err != nil {
    log.Println(err)
  } else {
    log.Println(resGetSwitch.Val)
  }
  time.Sleep(10 * time.Millisecond)
  resGetInfo, err = getInfo(baseUrl+PathInfo, resGetInfo)
  if err != nil {
    log.Println(err)
  } else {
    log.Println(resGetInfo.Device)
  }

  time.Sleep(10 * time.Millisecond)
  resGetStatus, err = getStatus(baseUrl+PathStatus, resGetStatus)
  if err != nil {
    log.Println(err)
  } else {
    log.Println(resGetStatus.Status)
  }
}*/

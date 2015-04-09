package web

import (
	"log"
//	"time"
//	"encoding/json"
	"net/http"
	"github.com/mgalela/gas3/db"
	"github.com/mgalela/gas3/rest"
//	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type Devreg struct {
  Mac	string	`json:"mac"`
  Ip	string	`json:"ip"`
  Land	string	`json:"land"`
}

func DevReg(c *gin.Context) {
  var devnew Devreg
  var devinfo rest.DevInfo
  var device db.Device
  var err error
  bedaip := false

  if !c.Bind(&devnew) {
    c.JSON(http.StatusBadRequest, gin.H{"status": "invalid input format"})
    return
  }
  log.Println(devnew.Mac)
  res := db.DB.Where("mac = ?",devnew.Mac).First(&device)
  if (res.Error == nil) {
    if (device.Ip == devnew.Ip) {
      log.Println("data found in database")
      return
    }
    log.Println("data found with different ip")
    bedaip = true
  }
  //getDevInfo, save to database
  devinfo, err = rest.GetInfo("http://"+devnew.Ip+"/"+rest.PathInfo, devinfo)
  if err != nil {
    log.Println(err)
    return
  }
  device.Wifimode = devinfo.Wifimode
  device.Wifissid = devinfo.Wifissid
  device.Wifipwd = devinfo.Wifipwd
  device.Wifiip = devinfo.Wifiip
  device.Wifinetmask = devinfo.Wifinetmask
  device.Wifigateway = devinfo.Wifigateway
  device.Device = devinfo.Device
  device.Land = devinfo.Land
  device.Serial = devinfo.Serial
  device.Mac = devinfo.Mac
  device.Ip = devinfo.Ip
  device.Landhq = devinfo.Landhq

  if bedaip {
    if res = db.DB.Save(&device); res.Error != nil {
      log.Println(res.Error)
      return
    }
  } else {
    if res := db.DB.Create(&device); res.Error != nil {
      log.Println(res.Error)
      return
    }
  }
  log.Println(device)
}

func Devices(c *gin.Context) {
  var devices []db.Device

  res := db.DB.Limit(50).Find(&devices)
  if (res.Error != nil) || (len(devices) == 0) {
    log.Println("data not found in database")
    c.JSON(http.StatusNotFound, gin.H{"status": "data not found"})
    return
  }
  c.JSON(http.StatusOK, &devices)
}

func Device(c *gin.Context) {
  var device db.Device
  deviceId := c.Params.ByName("id")

  res := db.DB.Where("id = ?",deviceId).First(&device)
  if res.Error != nil {
    log.Println("data not found in database")
    c.JSON(http.StatusNotFound, gin.H{"status": "data not found"})
    return
  }
  c.JSON(http.StatusOK, &device)
}

func DeviceNew(c *gin.Context) {
  var device db.Device

  if !c.Bind(&device) {
    c.JSON(http.StatusBadRequest, gin.H{"status": "invalid input format"})
    return
  }
  if res := db.DB.Create(&device); res.Error != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"status": res.Error})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"status": "OK"})
}

func DeviceUpdate(c *gin.Context) {
  var device,deviceUpdate db.Device
  deviceId := c.Params.ByName("id")

  if !c.Bind(&deviceUpdate) {
    c.JSON(http.StatusBadRequest, gin.H{"status": "invalid input format"})
    return
  }
  res := db.DB.Where("id = ?",deviceId).First(&device)
  if res.Error != nil {
    log.Println("data not found in database")
    c.JSON(http.StatusNotFound, gin.H{"status": "data not found"})
    return
  }
  if deviceUpdate.Id != device.Id {
    c.JSON(http.StatusBadRequest, gin.H{"status": "input not match with data entry"})
    return
  }
  device = deviceUpdate
  if res = db.DB.Save(&device); res.Error != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"status": res.Error})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"status": "OK"})
}

func DeviceDel(c *gin.Context) {
  deviceId := c.Params.ByName("id")
  var device db.Device

  res := db.DB.Where("id = ?",deviceId).First(&device)
  if res.Error != nil {
    log.Println("data not found in database")
    c.JSON(http.StatusNotFound, gin.H{"status": "data not found"})
    return
  }

  log.Println("delete device.device=", device.Device)

  if res = db.DB.Delete(&device); res.Error != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"status": res.Error})
    return
  }

  c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

/*
func Devices(c *gin.Context) {
	siteId := c.Params.ByName("id")
	var site VqaSites
	var devices []Device

	res := DB.Where("id = ?",siteId).First(&site)
	if res.Error != nil {
		log.Println("data not found in database")
		c.JSON(http.StatusNotFound, nil)
	}

	res = DB.Model(&site).Related(&devices, "SiteId").Order("devices.id")
	if res.Error != nil {
		log.Println("data not found in database")
		c.JSON(http.StatusNotFound, nil)
	}


	c.JSON(http.StatusOK, &devices)
}

func DevicesNew(c *gin.Context) {
	siteId := c.Params.ByName("id")
	var site VqaSites
	var device Device
	var err error

	res := DB.Where("id = ?",siteId).First(&site)
	if res.Error != nil {
		log.Println("data not found in database")
		c.JSON(http.StatusNotFound, nil)
	}

	if !c.Bind(&device) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid input format"})
		return
	}

	log.Println("device new=", device)

	if err = DB.Save(&device).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK"})
}

func DevicesUpdate(c *gin.Context) {
	devId := c.Params.ByName("id")
	var device Device
	var deviceUpdate Device
	var err error

	if !c.Bind(&deviceUpdate) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid input format"})
		return
	}

	res := DB.Where("id = ?",devId).First(&device)
	if res.Error != nil {
		log.Println("data not found in database")
		c.JSON(http.StatusNotFound, nil)
	}

	if deviceUpdate.Id != device.Id {
		c.JSON(http.StatusBadRequest, gin.H{"status": "input not match with data entry"})
		return
	}

	log.Println("update device.name=", device.Name)

	device = deviceUpdate
	if err = DB.Save(&device).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func DevicesDel(c *gin.Context) {
	devId := c.Params.ByName("id")
	var device Device
	var err error

	res := DB.Where("id = ?",devId).First(&device)
	if res.Error != nil {
		log.Println("data not found in database")
		c.JSON(http.StatusNotFound, nil)
	}

	log.Println("delete device.name=", device.Name)

	if err = DB.Delete(&device).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
*/

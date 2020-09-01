package main

import (
	"net/http"
	"strconv"
	"tim_presse/tim_test_ms_util_gen_numrange/bootstrap"

	numrange "github.com/BayramGuenes/tim_utils_numrange"
	//numrange "tim_presse/tim_test_ms_util_gen_numrange/tim_utils_numrange"

	"github.com/gin-gonic/gin"
)

func gotoMonitor(c *gin.Context) {

	//t, _ := template.ParseFiles("templates/01_mainview.html")

	viewpagedata = ServiceStatTable{} //getServicesStatInfo()
	//t.Execute(w, viewpagedata)
	c.HTML(http.StatusOK, "01_mainview.html", gin.H{"ServicesStatInfo": viewpagedata.ServicesStatInfo})

}

func monitorCheckStatus(c *gin.Context) {
	//t, _ := template.ParseFiles("templates/01_mainview.html")
	viewpagedata = getServicesStatInfo()
	//t.Execute(w, viewpagedata)
	c.HTML(http.StatusOK, "01_mainview.html", gin.H{"ServicesStatInfo": viewpagedata.ServicesStatInfo})

}

func gotoErrSubscriber(c *gin.Context) {

}

func servicesPerformOp(c *gin.Context) {
	//t, _ := template.ParseFiles("templates/01_mainview.html")
	c.Request.ParseForm()
	var execFunction string
	var execVal string
	var execSchema string
	var (
		tabnameCreate  string
		tabnameDisplay string
		tabnameGetNext string
		schemaCreate   string
		schemaDisplay  string
		schemaGetNext  string
		startNum       string
	)

	for key, value := range c.Request.PostForm {
		//fmt.Println("huhu"+key, value)
		//	execFunction := ""
		switch key {
		case "CreateNumRange":
			execFunction = key
		case "ShowNumRange":
			execFunction = key
		case "FuncGetNext":
			execFunction = key
		case "tabnameCreate":
			tabnameCreate = value[0]
		case "tabnameDisplay":
			tabnameDisplay = value[0]
		case "tabnameGetNext":
			tabnameGetNext = value[0]
		case "schemaCreate":
			schemaCreate = value[0]
		case "schemaDisplay":
			schemaDisplay = value[0]
		case "schemaGetNext":
			schemaGetNext = value[0]
		case "startNum":
			startNum = value[0]
		}
		//println(execFunction)

	}
	switch execFunction {
	case "CreateNumRange":
		execVal = tabnameCreate
		execSchema = schemaCreate
	case "ShowNumRange":
		execVal = tabnameDisplay
		execSchema = schemaDisplay
	case "FuncGetNext":
		execVal = tabnameGetNext
		execSchema = schemaGetNext
	}
	viewpagedata = performOp(execFunction, execSchema, execVal, startNum)
	//t.Execute(w, viewpagedata)
	c.HTML(http.StatusOK, "01_mainview.html", gin.H{"ServicesStatInfo": viewpagedata.ServicesStatInfo})

}
func performOp(iExecFunc, iExecSchema, iParam string, iStartNum string) (eServicesStatTable ServiceStatTable) {
	eServicesStatTable = ServiceStatTable{}
	startNum, _ := strconv.ParseInt(iStartNum, 10, 64)

	//println("bootstrap.ServiceDB, bootstrap.PortDB" + bootstrap.ServiceDB + ":" + bootstrap.PortDB)
	lOnr := numrange.NewManager(bootstrap.ServiceDB, bootstrap.PortDB, bootstrap.ApplConf.DB_UserPwd, iExecSchema)
	switch iExecFunc {
	case "CreateNumRange":
		lOut := lOnr.CreateNumRange(iParam, startNum)
		if lOut.Exception.Occured {
			eServicesStatTable.ServicesStatInfo.ResultCreateRange = lOut.Exception.ErrTxt
			eServicesStatTable.ServicesStatInfo.SchemaCreateRange = iExecSchema
			eServicesStatTable.ServicesStatInfo.CreateRangeTabName = iParam
		} else {
			eServicesStatTable.ServicesStatInfo.ResultCreateRange = "Created tab:" + lOut.TabnameNumRange
			eServicesStatTable.ServicesStatInfo.SchemaCreateRange = iExecSchema
			eServicesStatTable.ServicesStatInfo.CreateRangeTabName = iParam
		}
	case "ShowNumRange":
		lOut := lOnr.DisplayNumRange(iParam)
		if lOut.Exception.Occured {
			eServicesStatTable.ServicesStatInfo.ResultShowRange = lOut.Exception.ErrTxt
			eServicesStatTable.ServicesStatInfo.SchemaShowRange = iExecSchema
			eServicesStatTable.ServicesStatInfo.ShowStateTabName = iParam
		} else {
			s := strconv.FormatInt(lOut.LastIDTabOwner, 10)
			eServicesStatTable.ServicesStatInfo.ResultShowRange = "LastIDTabOwner:" + s
			eServicesStatTable.ServicesStatInfo.SchemaShowRange = iExecSchema
			eServicesStatTable.ServicesStatInfo.ShowStateTabName = iParam
		}
	case "FuncGetNext":
		lOut := lOnr.GetNextNumber(iParam)
		if lOut.Exception.Occured {
			eServicesStatTable.ServicesStatInfo.ResultNextTabName = lOut.Exception.ErrTxt
			eServicesStatTable.ServicesStatInfo.GetNextTabName = iParam
			eServicesStatTable.ServicesStatInfo.SchemaNextRange = iExecSchema
		} else {
			s := strconv.FormatInt(lOut.Number, 10)
			eServicesStatTable.ServicesStatInfo.ResultNextTabName = "Number:" + s
			eServicesStatTable.ServicesStatInfo.GetNextTabName = iParam
			eServicesStatTable.ServicesStatInfo.SchemaNextRange = iExecSchema
		}
	}

	return
}

func getServicesStatInfo() (eServicesStatTable ServiceStatTable) {
	eServicesStatTable = ServiceStatTable{}

	/*eServicesStatTable.ServicesStatInfo.CreateRangeTabName = "CreateTab"
	eServicesStatTable.ServicesStatInfo.ResultCreateRange = "ok"
	eServicesStatTable.ServicesStatInfo.ShowStateTabName = "ShowTab"
	eServicesStatTable.ServicesStatInfo.ResultShowRange = "okok"
	eServicesStatTable.ServicesStatInfo.GetNextTabName = "next"
	eServicesStatTable.ServicesStatInfo.ResultNextTabName = "45"*/
	/*eServicesStatTable.ServicesStatInfo = ServiceStatStruc{}

	var serviceSingleInfo ServiceStatStruc

	listRegisteredServ := getServicesAll()
	listDeadServ := getServicesNotAlive()

	for i := 0; i < len(listRegisteredServ); i++ {

		serviceSingleInfo.ServiceName = listRegisteredServ[i].LogicalName
		serviceSingleInfo.ServiceURL = listRegisteredServ[i].URL
		serviceSingleInfo.LastCheckResultText = listRegisteredServ[i].LastCheckResultText

		isDead := false
		for j := 0; j < len(listDeadServ); j++ {
			if listRegisteredServ[i].LogicalName == listDeadServ[j].LogicalName &&
				listRegisteredServ[i].URL == listDeadServ[j].URL {
				isDead = true
			}
		}
		if !isDead {
			serviceSingleInfo.ServiceStatus = "alive"
		} else {
			serviceSingleInfo.ServiceStatus = "dead"
		}
		eServicesStatTable.ServicesStatInfo = append(eServicesStatTable.ServicesStatInfo, serviceSingleInfo)

	}*/
	return eServicesStatTable
}

func getServicesAll() (rRegisteredServices listOfServices) {

	rRegisteredServices = nil

	//urlmonitor :=  instanceResourceMap.getResourcePhysicalValue(RESNAME_URL_PSSYSTEMMONITOR)

	return rRegisteredServices

}

/*====================================================================

=====================================================================*/
func getServicesNotAlive() (rRegisteredServices listOfServices) {

	rRegisteredServices = nil

	return rRegisteredServices

}

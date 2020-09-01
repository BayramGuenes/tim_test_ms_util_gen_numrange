package bootstrap

import (
	"flag"
	"os"
	"strings"
)

func LoadExternalSettings() {

	loadFromOSFlags()
	// Überschreiben OS Flags durch ENV Param
	loadFromOSEnv()
	// Überschreiben Parameter, falls durch args mitgegeben
	loadFromOSArgs()
	//println("ThisPort:" + ThisPort)
	//	println("ServiceDB, PortDB:" + ServiceDB + ":" + PortDB)
	return
}

func loadFromOSFlags() {
	protokollPointer := flag.String("protokoll", "http", "application-communication protokoll")
	servLocaIDSysRegPointer := flag.String("serviceLocationIDSysReg", "127.0.0.1", "service location id PSSysReg")
	portPointer := flag.String("port", "7000", "application port")
	servDBPointer := flag.String("serviceDB", "127.0.0.1", "service location id DBServer")
	portDBPointer := flag.String("DBport", "3306", "DB port")

	portPresServerPointer := flag.String("PortPresentation", "7500", "presentation port")
	confLocationPointer := flag.String("confLocation", "", "confPSSysCompReg.json path")

	flag.Parse()
	protokoll := *protokollPointer
	port := *portPointer
	confLocation := *confLocationPointer
	portPresServer := *portPresServerPointer
	servLocaIDSysReg := *servLocaIDSysRegPointer
	ServiceDB = *servDBPointer
	PortDB = *portDBPointer
	ThisCommunnicationProtokoll = protokoll
	ThisPort = port
	ThisConfLocation = confLocation
	PortPresentation = portPresServer
	ServiceLocationIDSysReg = servLocaIDSysReg
	if ThisCommunnicationProtokoll == "http" {
		ThisServerAddr = ThisCommunnicationProtokoll + "://" + ServiceLocationIDSysReg + ":" + ThisPort
	}
	return
}

func loadFromOSEnv() {

	servLocaIDSysReg, protokoll, port, servDB, portDB, confLocation, portPresServer := ThisCommunnicationProtokoll,
		ServiceLocationIDSysReg, ThisPort, ServiceDB, PortDB, ThisConfLocation, PortPresentation
	environList := os.Environ()
	leng := len(environList)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osenvparamval := environList[i]
			splittedString := strings.Split(osenvparamval, "=")
			paramname := splittedString[0]
			paramval := splittedString[1]

			if paramname == "serviceLocationIDSysReg" || paramname == "SERVICELOCATIONIDSYSREG" {
				servLocaIDSysReg = paramval
			}

			if paramname == "port" || paramname == "PORT" || paramname == "Port" ||
				paramname == "port_sysreg" || paramname == "PORT_SYSREG" || paramname == "Port_Sysreg" {
				port = paramval
			}

			if paramname == "serviceDB" || paramname == "SERVICEDB" {
				servDB = paramval
			}

			if paramname == "portDB" || paramname == "PORTDB" || paramname == "PortDB" {
				portDB = paramval
			}

			if paramname == "protokoll" || paramname == "Protokoll" || paramname == "PROTOKOLL" {
				protokoll = paramval
			}
			if paramname == "confLocation" || paramname == "ConfLocation" ||
				paramname == "conflocation" || paramname == "Conflocation" || paramname == "CONFLOCATION" {
				confLocation = paramval
			}
			if paramname == "portpresentation" || paramname == "portPresentation" || paramname == "PortPresentation" || paramname == "PORTPRESENTATION" {
				portPresServer = paramval
			}

		}
	}
	ServiceLocationIDSysReg,
		ThisCommunnicationProtokoll, ThisPort, ServiceDB, PortDB, ThisConfLocation, PortPresentation = servLocaIDSysReg, protokoll,
		port, servDB, portDB, confLocation, portPresServer
	if ThisCommunnicationProtokoll == "http" {
		ThisServerAddr = ThisCommunnicationProtokoll + "://" + ServiceLocationIDSysReg + ":" + ThisPort
	}
}

func loadFromOSArgs() {
	servLocaIDSysReg, protokoll, port, servDB, portDB, confLocation, portPres := ServiceLocationIDSysReg, ThisCommunnicationProtokoll,
		ThisPort, ServiceDB, PortDB, ThisConfLocation, PortPresentation

	leng := len(os.Args)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osparam := os.Args[i]
			splittedString := strings.Split(osparam, "=")
			var namevalues []string
			namevalues = append(namevalues, splittedString...)
			leng := len(namevalues)
			if leng > 1 {
				paramname := splittedString[0]
				paramval := splittedString[1]

				if paramname == "serviceLocationIDSysReg" || paramname == "SERVICELOCATIONIDSYSREG" {
					servLocaIDSysReg = paramval
				}
				if paramname == "port" || paramname == "PORT" || paramname == "Port" ||
					paramname == "port_sysreg" || paramname == "PORT_SYSREG" || paramname == "Port_Sysreg" {
					port = paramval
				}
				if paramname == "protokoll" || paramname == "Protokoll" || paramname == "PROTOKOLL" {
					protokoll = paramval
				}
				if paramname == "serviceDB" || paramname == "SERVICEDB" {
					servDB = paramval
				}

				if paramname == "portDB" || paramname == "PORTDB" || paramname == "PortDB" {
					portDB = paramval
				}
				if paramname == "confLocation" || paramname == "ConfLocation" ||
					paramname == "conflocation" || paramname == "Conflocation" || paramname == "CONFLOCATION" {
					confLocation = paramval
				}
				if paramname == "portpresentation" || paramname == "portPresentation" || paramname == "PortPresentation" || paramname == "PORTPRESENTATION" {
					portPres = paramval
				}

			}

		}
	}
	ServiceLocationIDSysReg, ThisCommunnicationProtokoll, ThisPort, ServiceDB, PortDB, ThisConfLocation, PortPresentation =
		servLocaIDSysReg,
		protokoll, port, servDB, portDB, confLocation, portPres
	if ThisCommunnicationProtokoll == "http" {
		ThisServerAddr = ThisCommunnicationProtokoll + "://" + ServiceLocationIDSysReg + ":" + ThisPort
	}
}

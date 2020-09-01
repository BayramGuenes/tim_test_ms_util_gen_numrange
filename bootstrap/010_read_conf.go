package bootstrap

import (
	"encoding/json"
	"errors"
	"log" //"strconv"
	"os"
)

func GetApplConf() (eConf ApplConfStruct, err error) {

	ApplConf = ApplConfStruct{}
	if len(ThisConfLocation) == 0 {
		errString :=
			"Achtung !!!  Service konnte nicht gestartet werden." +
				" Bitte Pfad der Konfigurationsdatei configtestnumrange.json  über Umgebungsparameter" +
				" 'confLocation' bzw als Startargument(osparam) 'confLocation=<path>' angeben."
		println(errString)
		return ApplConf, errors.New(errString)
	}

	filePathAndName := ThisConfLocation + string(os.PathSeparator) + "configtestnumrange.json"
	file, _ := os.Open(filePathAndName)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ApplConf)

	if err != nil {
		println("Decoding conf.json"+" FAILED", err.Error())
		log.Fatal(err.Error() + ". Möglicherweise fehlt die configtestnumrange.json Datei im Pfad " + ThisConfLocation)
		return ApplConf, err
	}
	eConf = ApplConf
	return eConf, nil

}

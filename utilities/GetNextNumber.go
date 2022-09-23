package utilities

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetNextCode(lastData string, typeKode string) string {

	var regex, err = regexp.Compile("(\\d+)")

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(lastData, -1)
	var getIncrement, _ = strconv.Atoi(res1[0])
	var databaru = getIncrement + 1
	fmt.Println(databaru)
	dataKodeBaru := ""
	var dataAngka = strconv.Itoa(databaru)

	if len(dataAngka) == 1 {
		dataKodeBaru = typeKode + "000" + dataAngka
	} else if len(dataAngka) == 2 {
		dataKodeBaru = typeKode + "00" + dataAngka
	} else if len(dataAngka) == 3 {
		dataKodeBaru = typeKode + "0" + dataAngka
	} else if len(dataAngka) == 3 {
		dataKodeBaru = typeKode + dataAngka
	}

	return dataKodeBaru
}

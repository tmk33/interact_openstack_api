package instances

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    
)



var TOKEN = "gAAAAABjrJ9vHoIN3fJIPCE-TpZuDgWOMSM9-nSgM2V1vEEqaAp4ULSbup4vgW-domODU00zTSI4R-HJIfV2PQSoOkQvlFbB1YOV_WhYaJRD2PwcRLAOOFLxPRuas1_tmucS_7WdprwE4C1-y8ZdkhUPgXh-DgfcJWLdj3X-Dl36dxnZQ2PUt98"

type Instances_Struct struct {
	Servers []Servers `json:"servers"`
}
type Links struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}
type Servers struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Links []Links `json:"links"`
}


// instances list
func GetAllInstances() Instances_Struct {
	URL := fmt.Sprintf("http://voscontrol:8774/v2.1/servers")
    
	fmt.Println("----------------All Instances----------------")

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token",TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Instances_Struct
	json.Unmarshal(bodyBytes, &responseObject)
	
	return responseObject
}

// print all instances
func ShowInstances() {
	instances := GetAllInstances()
	for _, value := range instances.Servers {
		fmt.Printf("Name: %+v\n", value.Name)
	}
}

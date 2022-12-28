package instances

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    
)



var TOKEN = "gAAAAABjrMYu3PChpgp8MICaEs9eCehTJz_S4C-_AXvXyhgrH1LXxnP18rywnVH2hio37O7qz0qdU8Gz7vGDL2LfAIyXoVZtnWSrLV52Axvvl5OOQ6B929RVZCLYT3NxFMIudjrcf_n6uv7Btw9-eL3MJAT6LTHRivzIWNczTNuOpyXmRBeKiiA"

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

// delete instance
func DeleteInstance(id string) {
	URL := fmt.Sprintf("http://voscontrol:8774/v2.1/servers/%+v", id)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", URL, nil)
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

	if resp.StatusCode == 204 {
		fmt.Printf("Status Code: %+v\n", resp.StatusCode)
		fmt.Printf("Instance deleted!")
	} else {
		fmt.Printf("Error!! Status Code: %+v\n", resp.StatusCode)
	}
	

}

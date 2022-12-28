package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "github.com/TMK33/interact_openstack_api/compute/images"
)

var TOKEN = "gAAAAABjrIzmVFy52OEgz6eKkjBq3vMvQ5L0Xvlgd89N8K1RJflg_KMc1sG4QBRdaSrI6Xr5yvEHkqec9enz6OJYWYGYgXPnNIS_t4MTubL46A_8pWXyvW0jh3IGoEaOZvfzHcBrCIqec5dNEqoW-_a2GE3vfdQKB9mZIu31e3UW3x8U6wjnndc"

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


//keypairs struct
type Keypairs_Struct struct {
	Keypairs []struct {
		Keypair struct {
			Name        string `json:"name"`
			PublicKey   string `json:"public_key"`
			Fingerprint string `json:"fingerprint"`
		} `json:"keypair"`
	} `json:"keypairs"`
}

//Volume struct
type Volumes_Struct struct {
	Volumes []struct {
		ID               string `json:"id"`
		Status           string `json:"status"`
		Size             int    `json:"size"`
		AvailabilityZone string `json:"availabilityZone"`
		CreatedAt        string `json:"createdAt"`
		Attachments      []struct {
			ID       string `json:"id"`
			VolumeID string `json:"volumeId"`
			ServerID string `json:"serverId"`
			Device   string `json:"device"`
		} `json:"attachments"`
		DisplayName        string      `json:"displayName"`
		DisplayDescription interface{} `json:"displayDescription"`
		VolumeType         string      `json:"volumeType"`
		SnapshotID         interface{} `json:"snapshotId"`
		Metadata           struct {
		} `json:"metadata"`
	} `json:"volumes"`
}

// instances list
func getAllInstances() Instances_Struct {
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
func showInstances() {
	instances := getAllInstances()
	for _, value := range instances.Servers {
		fmt.Printf("Name: %+v\n", value.Name)
	}
}


// List keypairs
func getAllKeypairs() Keypairs_Struct {
	URL := fmt.Sprintf("http://voscontrol:8774/v2.1/os-keypairs")
    
	fmt.Println("----------------All Keypairs----------------")

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
	var responseObject Keypairs_Struct
	json.Unmarshal(bodyBytes, &responseObject)
	
	return responseObject
}

// print all keypairs
func showKeypairs() {
	keypairs := getAllKeypairs()
	for _, value := range keypairs.Keypairs {
		fmt.Printf("Name: %+v\n", value.Keypair.Name)
	}
}

// instances list
func getAllVolumes() Volumes_Struct {
	URL := fmt.Sprintf("http://voscontrol:8774/v2.1/os-volumes")
    
	fmt.Println("----------------All Volumes----------------")

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
	var responseObject Volumes_Struct
	json.Unmarshal(bodyBytes, &responseObject)
	
	return responseObject
}

// print all instances
func showVolumes() {
	volumes := getAllVolumes()
	for _, value := range volumes.Volumes {
		fmt.Printf("Name: %+v\n", value.DisplayName)
	}
}

func main() {

	showInstances()
	//images.images.showImages()
	showKeypairs()
	showVolumes()
	
	fmt.Printf("\n")
	
}

package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
)

var TOKEN = "gAAAAABjq4_5d9OKdHXr0m2dtHCFoi5TqnknuC77QRh_J8zlC6vZUvUrRVqAw_m0apTxKmHhErApMaMQuDwjCMkp0LdCRYWUFuCTPg9m_ZMxUr-r2fNJwNVvoJJ3Zj6oq7ba24c4fMUDtt9Y1ek-rtt-2A20mrg4TgaQpeQoZKE22e0L_cIyyIE"

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

//images struct
type Images_Struct struct {
	Images []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
			Type string `json:"type,omitempty"`
		} `json:"links"`
	} `json:"images"`
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

// instances list
func listInstances() Instances_Struct {
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
func showInstances(instances Instances_Struct) {
	for _, value := range instances.Servers {
		fmt.Printf("Name: %+v\n", value.Name)
	}
}

// images list
func listImages() Images_Struct{
	URL := fmt.Sprintf("http://voscontrol:8774/v2.1/images")
    
	fmt.Println("----------------All Images----------------")

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
	var responseObject Images_Struct
	json.Unmarshal(bodyBytes, &responseObject)
	
	return responseObject
}

// print all images
func showImages(images Images_Struct) {
	for _, value := range images.Images {
		fmt.Printf("Name: %+v\n", value.Name)
	}
}

// List keypairs
func listKeypairs() Keypairs_Struct {
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
func showKeypairs(keypairs Keypairs_Struct) {
	for _, value := range keypairs.Keypairs {
		fmt.Printf("Name: %+v\n", value.Keypair.Name)
	}
}

func main() {

	
	all_instances := listInstances()
	showInstances(all_instances)
	
	fmt.Printf("\n")
	
}

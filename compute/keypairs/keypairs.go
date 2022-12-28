package keypairs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    
)



var TOKEN = "gAAAAABjrJ9vHoIN3fJIPCE-TpZuDgWOMSM9-nSgM2V1vEEqaAp4ULSbup4vgW-domODU00zTSI4R-HJIfV2PQSoOkQvlFbB1YOV_WhYaJRD2PwcRLAOOFLxPRuas1_tmucS_7WdprwE4C1-y8ZdkhUPgXh-DgfcJWLdj3X-Dl36dxnZQ2PUt98"


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

// List keypairs
func GetAllKeypairs() Keypairs_Struct {
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
func ShowKeypairs() {
	keypairs := GetAllKeypairs()
	for _, value := range keypairs.Keypairs {
		fmt.Printf("Name: %+v\n", value.Keypair.Name)
	}
}

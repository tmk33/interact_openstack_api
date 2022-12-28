package volumes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var TOKEN = "gAAAAABjrJ9vHoIN3fJIPCE-TpZuDgWOMSM9-nSgM2V1vEEqaAp4ULSbup4vgW-domODU00zTSI4R-HJIfV2PQSoOkQvlFbB1YOV_WhYaJRD2PwcRLAOOFLxPRuas1_tmucS_7WdprwE4C1-y8ZdkhUPgXh-DgfcJWLdj3X-Dl36dxnZQ2PUt98"


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


// Volume list
func GetAllVolumes() Volumes_Struct {
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

// print all Volumes
func ShowVolumes() {
	volumes := GetAllVolumes()
	for _, value := range volumes.Volumes {
		fmt.Printf("Name: %+v\n", value.DisplayName)
	}
}


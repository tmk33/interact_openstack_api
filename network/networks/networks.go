package networks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var TOKEN = "gAAAAABjrLDir8QfCXexDasHewDKxpQkxTL7fmF4n5wOgQ_7ONV-wtmxOMKpkOzjk4Qgm93MvWxF5nMRn_FpJHhXrc57iqKL8cn_kNUALGebuCKTPUhPU-aKRqJy83BoYzbatntvbr_AH3n07hFOk5hWIBXzEni3uKxehF0TZN36NVtbEhwTMN0"


//network struct
type Networks_Struct struct {
	Networks []struct {
		Bridge            interface{} `json:"bridge"`
		BridgeInterface   interface{} `json:"bridge_interface"`
		Broadcast         interface{} `json:"broadcast"`
		Cidr              interface{} `json:"cidr"`
		CidrV6            interface{} `json:"cidr_v6"`
		CreatedAt         interface{} `json:"created_at"`
		Deleted           interface{} `json:"deleted"`
		DeletedAt         interface{} `json:"deleted_at"`
		DhcpServer        interface{} `json:"dhcp_server"`
		DhcpStart         interface{} `json:"dhcp_start"`
		DNS1              interface{} `json:"dns1"`
		DNS2              interface{} `json:"dns2"`
		EnableDhcp        interface{} `json:"enable_dhcp"`
		Gateway           interface{} `json:"gateway"`
		GatewayV6         interface{} `json:"gateway_v6"`
		Host              interface{} `json:"host"`
		ID                string      `json:"id"`
		Injected          interface{} `json:"injected"`
		Label             string      `json:"label"`
		Mtu               interface{} `json:"mtu"`
		MultiHost         interface{} `json:"multi_host"`
		Netmask           interface{} `json:"netmask"`
		NetmaskV6         interface{} `json:"netmask_v6"`
		Priority          interface{} `json:"priority"`
		ProjectID         interface{} `json:"project_id"`
		RxtxBase          interface{} `json:"rxtx_base"`
		ShareAddress      interface{} `json:"share_address"`
		UpdatedAt         interface{} `json:"updated_at"`
		Vlan              interface{} `json:"vlan"`
		VpnPrivateAddress interface{} `json:"vpn_private_address"`
		VpnPublicAddress  interface{} `json:"vpn_public_address"`
		VpnPublicPort     interface{} `json:"vpn_public_port"`
	} `json:"networks"`
}


// Volume list
func GetAllNetworks() Networks_Struct {
	URL := fmt.Sprintf("http://voscontrol:8774/v2.1/os-networks")
    
	fmt.Println("----------------All Networks----------------")

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
	var responseObject Networks_Struct
	json.Unmarshal(bodyBytes, &responseObject)
	
	return responseObject
}

// print all Volumes
func ShowNetworks() {
	networks := GetAllNetworks()
	for _, value := range networks.Networks {
		fmt.Printf("Name: %+v\n", value.Label)
	}
}


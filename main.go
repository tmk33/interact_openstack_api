package main

import (
 "fmt"
 _"interact_openstack_api/compute/images"
 _"interact_openstack_api/compute/keypairs"
 "interact_openstack_api/compute/instances"
 _"interact_openstack_api/volumes/volumes"
 _"interact_openstack_api/network/networks"


)

var TOKEN = "gAAAAABjrJ9vHoIN3fJIPCE-TpZuDgWOMSM9-nSgM2V1vEEqaAp4ULSbup4vgW-domODU00zTSI4R-HJIfV2PQSoOkQvlFbB1YOV_WhYaJRD2PwcRLAOOFLxPRuas1_tmucS_7WdprwE4C1-y8ZdkhUPgXh-DgfcJWLdj3X-Dl36dxnZQ2PUt98"


func main() {

	instances.ShowInstances()
	//images.ShowImages()
	//keypairs.ShowKeypairs()
	//volumes.ShowVolumes()
	//networks.ShowNetworks()
	//instances.DeleteInstance("2009d5d1-f320-4f9e-884c-710be0eca641")
	fmt.Printf("\n")
	
}

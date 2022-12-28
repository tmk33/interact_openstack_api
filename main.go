package main

import (
 "fmt"
 "interact_openstack_api/compute/images"
 "interact_openstack_api/compute/keypairs"
 "interact_openstack_api/compute/instances"
 "interact_openstack_api/volumes/volumes"


)

var TOKEN = "gAAAAABjrJ9vHoIN3fJIPCE-TpZuDgWOMSM9-nSgM2V1vEEqaAp4ULSbup4vgW-domODU00zTSI4R-HJIfV2PQSoOkQvlFbB1YOV_WhYaJRD2PwcRLAOOFLxPRuas1_tmucS_7WdprwE4C1-y8ZdkhUPgXh-DgfcJWLdj3X-Dl36dxnZQ2PUt98"


func main() {

	instances.ShowInstances()
	images.ShowImages()
	keypairs.ShowKeypairs()
	volumes.ShowVolumes()

	fmt.Printf("\n")
	
}

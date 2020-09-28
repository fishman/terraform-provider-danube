package provider

import (
	"fmt"
	"log"

	"github.com/erigones/godanube/cloudapi"
	"github.com/fishman/terraform-provider-danube/provider/api"
	"github.com/fishman/terraform-provider-danube/provider/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceMachineCreate,
		Read:   resourceMachineRead,
		Update: resourceMachineUpdate,
		Delete: resourceMachineDelete,

		Schema: schemas.MachineSchema(),
	}
}

func resourceMachineCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudapi.Client)

	hostname := d.Get("hostname").(string)
	vcpus := d.Get("vcpus").(int)
	ram := d.Get("ram").(int)

	var tags []string
	for _, tag := range d.Get("tags").([]interface{}) {
		tags = append(tags, tag.(string))
	}

	mdata := map[string]string{}
	for k, v := range d.Get("mdata").(map[string]interface{}) {
		mdata[k] = v.(string)
	}

	vm_definition := cloudapi.MachineDefinition{
		Name: hostname,
		//DnsDomain: "lan",
		Vcpus: vcpus,
		Ram:   ram,
		Tags:  tags,
		Mdata: mdata,
	}

	if alias, ok := d.GetOk("alias"); ok {
		vm_definition.Alias = alias.(string)
	}

	disk_definition := []cloudapi.VmDiskDefinition{
		{
			//Size:  10240,
			//Image: "ubuntu-certified-18.04",
			Image: "centos-8-kvm",
		},
	}

	network_definition := []cloudapi.VmNicDefinition{
		{
			Net: "lan",
		},
	}

	opts := cloudapi.CreateMachineOpts{
		Vm:    vm_definition,
		Disks: disk_definition,
		Nics:  network_definition,
	}

	vmInfo, err := client.CreateMachine(opts)

	if err != nil {
		fmt.Println("error:" + err.Error())
		return err
	}

	d.SetId(vmInfo.Uuid)

	resourceMachineRead(d, meta)

	return err
}

func resourceMachineRead(d *schema.ResourceData, meta interface{}) error {
	uuid := d.Id()

	return api.MachineRead(uuid, d, meta)
}

func resourceMachineUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMachineDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudapi.Client)
	uuid := d.Id()
	force := true

	err := client.DeleteMachine(uuid, force)

	if err != nil {
		log.Println("[ERR] ", err.Error)
		return err
	}

	return err
}

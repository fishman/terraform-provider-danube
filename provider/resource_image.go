package provider

import (
	"github.com/fishman/terraform-provider-danube/provider/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageCreate,
		Read:   resourceImageRead,
		Update: resourceImageUpdate,
		Delete: resourceImageDelete,

		Schema: schemas.ImageSchema(),
	}
}

func resourceImageCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cloudapi.Client)

	// vmInfo, err := client.ImportImage(remoteImageUuid string, newImageName string, repoName string)

	// if err != nil {
	// 	fmt.Println("error:" + err.Error())
	// 	return err
	// }

	// d.SetId(vmInfo.Uuid)

	// resourceImageRead(d, meta)

	// return err
	return nil
}

func resourceImageRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceImageUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceImageDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

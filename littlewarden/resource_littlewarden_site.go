package littlewarden

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceLittlewardenSite() *schema.Resource {
	return &schema.Resource{
		Read: resourceLittlewardenSiteRead,
		Schema: map[string]*schema.Schema{

			"url": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed attributes
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceLittlewardenSiteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	log.Printf("Stuff: %s", config.ApiKey)

	return nil
}

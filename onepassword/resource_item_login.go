package onepassword

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceItemLogin() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemLoginRead,
		Create: resourceItemLoginCreate,
		Delete: resourceItemLoginDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				resourceItemLoginRead(d, meta)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Item login name.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Item login username.",
			},
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Sensitive:   true,
				Description: "Item login password.",
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "Item login tags.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"vault": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Vault for item login.",
			},
			"section": {
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Description: "Item login section.",
				Elem:        sectionSchema,
			},
			"url": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				Description:  "URL for item login.",
				ValidateFunc: urlValidate,
			},
		},
	}
}

func resourceItemLoginRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	vaultId := d.Get("vault").(string)
	err, v := m.onePassClient.ReadItem(getId(d), vaultId)
	log.Printf("[DEBUG] %v", v)
	if err != nil {
		return err
	}

	d.SetId(v.Uuid)
	d.Set("name", v.Overview.Title)
	return nil
}

func resourceItemLoginCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceItemLoginDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
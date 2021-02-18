package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRancher2PodSecurityPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceRancher2PodSecurityPolicyTemplateRead,
		Schema: podSecurityPolicyTemplateFields(),
	}
}

func dataSourceRancher2PodSecurityPolicyTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	name := d.Get("name").(string)

	var pspt *projectClient.AppCollection
	err = meta.(*Config).WithRetry(func() (err error) {
		pspt, err = client.PodSecurityPolicyTemplate.ByID(name)
		return err
	})
	if err != nil {
		return err
	}

	return flattenPodSecurityPolicyTemplate(d, pspt)
}

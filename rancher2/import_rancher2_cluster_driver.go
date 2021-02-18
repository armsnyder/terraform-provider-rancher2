package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRancher2ClusterDriverImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	var clusterDriver *projectClient.AppCollection
	err = meta.(*Config).WithRetry(func() (err error) {
		clusterDriver, err = client.KontainerDriver.ByID(d.Id())
		return err
	})
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	err = flattenClusterDriver(d, clusterDriver)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}

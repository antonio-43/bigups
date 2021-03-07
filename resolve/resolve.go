package resolve

import (
	"bgs/models"
	"bgs/resolver"
	"context"
)

func Resolve(obj *models.Mapper) models.Results {
	final := models.Results{}

	r := resolver.NewResolver(obj.DnsIp)

	ips, err := r.LookupHost(context.Background(), obj.Target)

	if err == nil {
		obj.Resolved = true

		obj.TargetIp = ips
		final.Tested = append(final.Tested, *obj)
	} else {
		obj.Resolved = false

		final.Tested = append(final.Tested, *obj)
	}

	return final
}

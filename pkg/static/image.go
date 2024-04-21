package static

import "github.com/mikezzb/steam-trading-server/pkg/setting"

func GetImageFullUrl() string {
	return setting.App.PrefixUrl + "/static/images/"
}

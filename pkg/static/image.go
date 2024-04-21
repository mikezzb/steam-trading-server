package static

import "github.com/mikezzb/steam-trading-server/pkg/setting"

func GetImageFullUrl() string {
	return setting.AppSetting.PrefixUrl + "/static/images/"
}

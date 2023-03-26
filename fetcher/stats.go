package fetcher

import (
	"fmt"

	"github.com/wao3/luogu-stats-card/cache"
	"github.com/wao3/luogu-stats-card/model/fetch"
)

var UserProfileCache = cache.NewCache[fetch.UserProfileData](cache.DefaultExpireTime, cache.DefaultPurgeTime)

func FetchUserProfile(uid int) (*fetch.UserProfileData, error) {
	var key = fmt.Sprintf("UserProfile_%d", uid)
	if value, ok := UserProfileCache.Get(key); ok {
		return value, nil
	}
	userProfile, err := Fetch[fetch.UserProfileData](fmt.Sprintf("/user/%d", uid), nil)
	if err != nil {
		return nil, err
	}
	UserProfileCache.Set(key, userProfile)
	return userProfile, nil
}

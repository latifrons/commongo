package safe_viper

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ViperMustGetString(key string) string {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		logrus.WithField("key", key).Fatal("config missing")
	}
	return viper.GetString(key)
}
func ViperMustGetInt(key string) int {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		logrus.WithField("key", key).Fatal("config missing")
	}
	return viper.GetInt(key)
}
func ViperMustGetBool(key string) bool {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		logrus.WithField("key", key).Fatal("config missing")
	}
	return viper.GetBool(key)
}

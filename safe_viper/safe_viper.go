package safe_viper

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
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
func ViperMustGetDuration(key string) time.Duration {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		logrus.WithField("key", key).Fatal("config missing")
	}
	s := viper.GetString(key)
	td, err := time.ParseDuration(s)
	if err != nil {
		logrus.WithField("key", key).Fatal("config invalid")
	}
	return td

}

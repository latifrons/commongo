package safe_viper

import (
	"time"

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

func ViperMustGetFloat64(key string) float64 {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		logrus.WithField("key", key).Fatal("config missing")
	}
	return viper.GetFloat64(key)
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

func ViperGetStringOrDefault(key, defaultValue string) string {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		return defaultValue
	}
	return viper.GetString(key)
}

func ViperGetIntOrDefault(key string, defaultValue int) int {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		return defaultValue
	}
	return viper.GetInt(key)
}

func ViperGetBoolOrDefault(key string, defaultValue bool) bool {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		return defaultValue
	}
	return viper.GetBool(key)
}

func ViperGetFloat64OrDefault(key string, defaultValue float64) float64 {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		return defaultValue
	}
	return viper.GetFloat64(key)
}

func ViperGetDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	if !viper.IsSet(key) || viper.GetString(key) == "" {
		return defaultValue
	}
	s := viper.GetString(key)
	td, err := time.ParseDuration(s)
	if err != nil {
		logrus.WithField("key", key).Fatal("config invalid")
	}
	return td
}

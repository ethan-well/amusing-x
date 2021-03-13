package powertrain

import (
	"amusingx.fit/amusingx/superfactory/configcontainer"
)

func Run(configObj configcontainer.Config) {
	 err := configcontainer.New().Unmarshal(configObj, "")
	 if err != nil {
	 	panic(err)
	 }
}
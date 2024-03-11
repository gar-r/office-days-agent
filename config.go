package main

const AppName = "office-days-agent"
const EnvConfigPrefix = "OFFICE_DAYS"

type Conf struct {
	ApiServerPort       int    `split_words:"true" default:"23460"`
	DBPath              string `split_words:"true" default:"$HOME/.config/office-days-agent/db"`
	PollIntervalSeconds int    `split_words:"true" default:"1800"`
	WifiName            string `split_words:"true" required:"true"`
	WifiDeviceName      string `split_words:"true" default:"en0"`
}

package main

const EnvConfigPrefix = "OFFICE_DAYS"

type Conf struct {
	PollIntervalSeconds int    `split_words:"true" default:"60"`
	WifiName            string `split_words:"true" required:"true"`
}

package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tmsmr/chilld/fancurve"
	"github.com/tmsmr/cm4iofan"
	"github.com/tmsmr/pithermal"
)

const (
	ticksPerSecond = 1
)

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	if !*debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})

	log.Info().Msgf("starting ChillD")

	fanctl, err := cm4iofan.New()
	if err != nil {
		log.Fatal().Str("err", err.Error()).Msg("failed to initialize fan controller")
	}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	ticks := time.NewTicker((1000 / ticksPerSecond) * time.Millisecond)
	defer ticks.Stop()

	currspeed := -1

	go func() {
		for range ticks.C {
			speed := 100
			temp, terr := pithermal.GetCpuTemp()
			if terr != nil {
				log.Error().Str("err", err.Error()).Msg("failed to read cpu temp, using maximum fan speed")
			} else {
				speed = fancurve.LinearFanSpeedFor(temp)
			}
			if speed != currspeed {
				msg := "setting fan speed to"
				if terr == nil {
					log.Debug().Msgf("%s %d%% @ %.2f°C", msg, speed, temp)
				} else {
					log.Warn().Msgf("%s %d%% @ UNKNOWN°C", msg, speed)
				}
				if err := fanctl.SetDutyCycle(speed); err != nil {
					log.Error().Str("err", err.Error()).Msg("failed to to set target speed in fan controller")
				} else {
					currspeed = speed
				}
			}
		}
	}()

	<-trap
	log.Info().Msg("ChillD terminating, trying to set maximum fan speed")
	_ = fanctl.SetDutyCycle(100)
}

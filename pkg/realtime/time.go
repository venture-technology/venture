package realtime

import "time"

const offset int = -3 * 60 * 60

var Now = defaultNowFunction

func defaultNowFunction() time.Time {
	loc := time.FixedZone("America/Sao_Paulo", offset)

	return time.Now().In(loc)
}

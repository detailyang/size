package size

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Unit int64

const (
	B   Unit = 1
	KB       = 1000 * B
	MB       = 1000 * KB
	GB       = 1000 * MB
	TB       = 1000 * GB
	PB       = 1000 * TB
	KiB      = 1024 * B
	MiB      = 1024 * KiB
	GiB      = 1024 * MiB
	TiB      = 1024 * GiB
	PiB      = 1024 * TiB
)

func Parse(s string) (Unit, error) {
	if strings.HasSuffix(s, "iB") && len(s) >= 4 {
		var unit Unit
		switch s[len(s)-3:] {
		case "KiB":
			unit = KiB
		case "MiB":
			unit = MiB
		case "GiB":
			unit = GiB
		case "TiB":
			unit = TiB
		case "PiB":
			unit = PiB
		default:
			return 0, errors.New("units: invalid units " + s)
		}
		s = s[:len(s)-3]
		f64, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		if float64(int64(f64)) == f64 {
			return Unit(f64) * unit, nil
		}

		return Unit(f64 * float64(unit)), nil

	} else if strings.HasSuffix(s, "B") && len(s) >= 2 {
		var unit Unit
		switch s[len(s)-2:] {
		case "KB":
			unit = KB
			s = s[:len(s)-2]
		case "MB":
			unit = MB
			s = s[:len(s)-2]
		case "GB":
			unit = GB
			s = s[:len(s)-2]
		case "TB":
			unit = TB
			s = s[:len(s)-2]
		case "PB":
			unit = PB
			s = s[:len(s)-2]
		default:
			if '0' <= s[len(s)-2] && s[len(s)-2] <= '9' {
				unit = B
				s = s[:len(s)-1]
			} else {
				return 0, errors.New("units: invalid units " + s)
			}
		}

		f64, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		if float64(int64(f64)) == f64 {
			return Unit(f64) * unit, nil
		}

		return Unit(f64 * float64(unit)), nil
	}

	return 0, errors.New("units: invalid units " + s)
}

func (b *Unit) UnmarshalText(text []byte) error {
	var err error
	*b, err = Parse(string(text))
	return err
}

// String returns a string representing the unit
func (b Unit) String() string {
	if b%10 == 0 { // KB
		var s string
		switch {
		case b >= PB:
			s = "PB"
			b /= PB
		case b >= TB:
			s = "TB"
			b /= TB
		case b >= GB:
			s = "GB"
			b /= GB
		case b >= MB:
			s = "MB"
			b /= MB
		case b >= KB:
			s = "KB"
			b /= KB
		default:
			s = "B"
		}
		return fmt.Sprintf("%d%s", b, s)
	}

	// KiB
	var s string
	switch {
	case b >= PiB:
		s = "PiB"
		b /= PiB
	case b >= TiB:
		s = "TiB"
		b /= TiB
	case b >= GiB:
		s = "GiB"
		b /= GiB
	case b >= MiB:
		s = "MiB"
		b /= MiB
	case b >= KiB:
		s = "KiB"
		b /= KiB
	default:
		s = "B"
	}
	return fmt.Sprintf("%d%s", b, s)
}

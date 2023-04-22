package log

import (
	"github.com/rs/zerolog"
	"os"
)

var Logger zerolog.Logger

func init() {
	writer := zerolog.ConsoleWriter{
		Out:                 os.Stderr,
		NoColor:             false,
		TimeFormat:          "",
		PartsOrder:          nil,
		PartsExclude:        nil,
		FieldsExclude:       nil,
		FormatTimestamp:     nil,
		FormatLevel:         nil,
		FormatCaller:        nil,
		FormatMessage:       nil,
		FormatFieldName:     nil,
		FormatFieldValue:    nil,
		FormatErrFieldName:  nil,
		FormatErrFieldValue: nil,
		FormatExtra:         nil,
	}

	logger := zerolog.New(writer).With().Timestamp().Logger()

	Logger = logger

}

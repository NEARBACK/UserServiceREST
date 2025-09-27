package entry

import (
	"os"
	"os/signal"
	"syscall"
	"useservice/internal/definitions"
	"useservice/internal/repository"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	dal *repository.DatabaseAccessLayer
	log definitions.Logger
}

func Initialize() (app *App, err error) {
	app = &App{}
	err = app.setup()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) setup() error {
	a.initLogger()
	if err := a.initializeDatabaseAccessLayer(); err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger() {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "trace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}

	zlogger := zap.New(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zap.FatalLevel),
	)
	a.log = zlogger.Sugar()
}

func (a *App) initConf() {
}

func (a *App) initializeDatabaseAccessLayer() error {

	db, err := repository.NewDatabaseAccessLayer(a.log)
	if err != nil {
		a.log.Error(err)
		return err
	}
	a.dal = db
	return nil
}

func (a *App) Start() {
}

func (a *App) awaitQuitSignal() {
	a.log.Info("Working until a quit signal is received...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func (a *App) Stop() (err error) {
	a.log.Debug("Stopping server...")
	a.dal.CloseAll()
	a.log.Debug("Database connections closed.")

	return nil
}

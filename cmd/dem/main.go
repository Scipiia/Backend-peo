package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"vue-golang/internal/config"
	"vue-golang/internal/storage/mysql"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustConfig()

	log := setupLogger(cfg.Env)

	storage, err := mysql.New()
	if err != nil {
		log.Error("failed to open db", err)
		os.Exit(1)
	}

	log.Info("server started", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      routes(*cfg, log, storage),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Error("failed start server ", err)
	}

	log.Error("server stopped")
}

//func setupLogger(env string) *slog.Logger {
//	var log *slog.Logger
//
//	switch env {
//	case envLocal:
//		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//	case envDev:
//		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//	case envProd:
//		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
//	}
//
//	return log
//}

type dualHandler struct {
	coreHandler  slog.Handler
	errorHandler slog.Handler
}

func (h *dualHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	// Разрешаем обработку, если хоть один из хендлеров может обработать уровень
	return h.coreHandler.Enabled(ctx, lvl) || h.errorHandler.Enabled(ctx, lvl)
}

func (h *dualHandler) Handle(ctx context.Context, r slog.Record) error {
	var err error

	// 1. Всегда пишем в основной вывод (stdout)
	if h.coreHandler.Enabled(ctx, r.Level) {
		err = h.coreHandler.Handle(ctx, r)
		if err != nil {
			return err
		}
	}

	// 2. Если это ошибка — пишем в файл
	if r.Level >= slog.LevelError && h.errorHandler.Enabled(ctx, r.Level) {
		// Клонируем запись, потому что Handle может мутировать
		cloned := r.Clone()
		fileErr := h.errorHandler.Handle(ctx, cloned)
		if fileErr != nil {
			// Не прерываем основной поток, но можем залогировать проблему
			// (хотя здесь уже сложно — лучше игнорировать)
		}
	}

	return err
}

func (h *dualHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &dualHandler{
		coreHandler:  h.coreHandler.WithAttrs(attrs),
		errorHandler: h.errorHandler.WithAttrs(attrs),
	}
}

func (h *dualHandler) WithGroup(name string) slog.Handler {
	return &dualHandler{
		coreHandler:  h.coreHandler.WithGroup(name),
		errorHandler: h.errorHandler.WithGroup(name),
	}
}

func setupLogger(env string) *slog.Logger {
	// Определяем уровень логирования
	var level slog.Level = slog.LevelDebug
	switch env {
	case envProd:
		level = slog.LevelInfo
	}

	// 1. Основной handler — пишет ВСЁ в stdout
	var coreHandler slog.Handler
	switch env {
	case envLocal:
		coreHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	case envDev:
		coreHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	case envProd:
		coreHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	default:
		coreHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}

	// 2. Файловый handler — только ошибки
	errorFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Если не удалось создать файл, хотя бы предупреждаем
		slog.Warn("Cannot open error log file", "error", err)
		return slog.New(coreHandler) // продолжаем без файла
	}

	errorHandler := slog.NewTextHandler(errorFile, &slog.HandlerOptions{
		Level: slog.LevelError, // Только error и выше
	})

	// 3. Объединяем через кастомный handler
	handler := &dualHandler{
		coreHandler:  coreHandler,
		errorHandler: errorHandler,
	}

	// Создаём логгер
	logger := slog.New(handler)

	// 💡 Сохранить errorFile где-то, если хотите закрыть в будущем (например, при graceful shutdown)
	// Но если логгер глобальный — можно не закрывать явно, или использовать sync.Pool / закрытие при выходе.
	return logger
}

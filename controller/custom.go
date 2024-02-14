package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type CustomContext struct {
	echo.Context
	Ctx context.Context
}

type MyHandler struct {
	slog.Handler
}

func (h MyHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value("requestId").(string); ok {
		r.Add("requestId", slog.StringValue(traceID))
	}

	return h.Handler.Handle(ctx, r)
}

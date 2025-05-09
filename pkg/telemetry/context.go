package telemetry

import "context"

type telemetryClientCtxKey struct{}

func Context(ctx context.Context, client Client) context.Context {
	return context.WithValue(ctx, telemetryClientCtxKey{}, client)
}

func FromContext(ctx context.Context) Client {
	client, _ := ctx.Value(telemetryClientCtxKey{}).(Client)
	if client == nil {
		return DefaultClient
	}
	return client
}

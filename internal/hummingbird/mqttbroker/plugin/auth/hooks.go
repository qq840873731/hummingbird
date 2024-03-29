package auth

import (
	"context"

	"go.uber.org/zap"

	"github.com/winc-link/hummingbird/internal/hummingbird/mqttbroker/server"
	"github.com/winc-link/hummingbird/internal/pkg/codes"
	"github.com/winc-link/hummingbird/internal/pkg/packets"
)

func (a *Auth) HookWrapper() server.HookWrapper {
	return server.HookWrapper{
		OnBasicAuthWrapper: a.OnBasicAuthWrapper,
	}
}

func (a *Auth) OnBasicAuthWrapper(pre server.OnBasicAuth) server.OnBasicAuth {
	return func(ctx context.Context, client server.Client, req *server.ConnectRequest) (err error) {
		err = pre(ctx, client, req)
		if err != nil {
			return err
		}
		ok, err := a.validate(string(req.Connect.Username), string(req.Connect.Password))
		if err != nil {
			return err
		}
		if !ok {
			log.Debug("authentication failed", zap.String("username", string(req.Connect.Username)))
			v := client.Version()
			if packets.IsVersion3X(v) {
				return &codes.Error{
					Code: codes.V3NotAuthorized,
				}
			}
			if packets.IsVersion5(v) {
				return &codes.Error{
					Code: codes.NotAuthorized,
				}
			}
		}
		return nil
	}
}

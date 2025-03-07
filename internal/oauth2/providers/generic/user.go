package generic

import (
	"context"

	"github.com/jkroepke/openvpn-auth-oauth2/internal/oauth2/idtoken"
	"github.com/jkroepke/openvpn-auth-oauth2/internal/oauth2/types"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func (p *Provider) GetUser(_ context.Context, tokens *oidc.Tokens[*idtoken.Claims]) (types.UserData, error) {
	var (
		preferredUsername string
		subject           string
	)

	if tokens.IDTokenClaims != nil {
		preferredUsername = tokens.IDTokenClaims.PreferredUsername
	}

	if tokens.IDTokenClaims != nil {
		subject = tokens.IDTokenClaims.Subject
	}

	return types.UserData{
		PreferredUsername: preferredUsername,
		Subject:           subject,
	}, nil
}

package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/formancehq/stack/libs/go-libs/bun/bunconnect"
	"github.com/formancehq/stack/libs/go-libs/licence"

	sdk "github.com/formancehq/formance-sdk-go/v2"
	"github.com/formancehq/reconciliation/internal/api"
	"github.com/formancehq/reconciliation/internal/storage"
	sharedapi "github.com/formancehq/stack/libs/go-libs/api"
	"github.com/formancehq/stack/libs/go-libs/auth"
	"github.com/formancehq/stack/libs/go-libs/otlp"
	"github.com/formancehq/stack/libs/go-libs/otlp/otlpmetrics"
	"github.com/formancehq/stack/libs/go-libs/otlp/otlptraces"
	"github.com/formancehq/stack/libs/go-libs/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func stackClientModule() fx.Option {
	return fx.Options(
		fx.Provide(func() *sdk.Formance {
			oauthConfig := clientcredentials.Config{
				ClientID:     viper.GetString(stackClientIDFlag),
				ClientSecret: viper.GetString(stackClientSecretFlag),
				TokenURL:     fmt.Sprintf("%s/api/auth/oauth/token", viper.GetString(stackURLFlag)),
				Scopes:       []string{"openid", "ledger:read", "ledger:write", "payments:read", "payments:write"},
			}
			underlyingHTTPClient := &http.Client{
				Transport: otlp.NewRoundTripper(http.DefaultTransport, viper.GetBool(service.DebugFlag)),
			}
			return sdk.New(
				sdk.WithClient(
					oauthConfig.Client(context.WithValue(context.Background(),
						oauth2.HTTPClient, underlyingHTTPClient)),
				),
				sdk.WithServerURL(viper.GetString(stackURLFlag)),
			)
		}),
	)
}

func newServeCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "serve",
		RunE: runServer(version),
	}
	cmd.Flags().String(listenFlag, ":8080", "Listening address")

	return cmd
}

func runServer(version string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		databaseOptions, err := prepareDatabaseOptions(cmd)
		if err != nil {
			return err
		}

		options := make([]fx.Option, 0)
		options = append(options, databaseOptions)

		options = append(options,
			otlptraces.CLITracesModule(),
			otlpmetrics.CLIMetricsModule(),
			auth.CLIAuthModule(),
		)

		options = append(options,
			stackClientModule(),
			api.HTTPModule(sharedapi.ServiceInfo{
				Version: version,
			}, viper.GetString(listenFlag)),
			licence.CLIModule(ServiceName),
		)

		return service.New(cmd.OutOrStdout(), options...).Run(cmd.Context())
	}
}

func prepareDatabaseOptions(cmd *cobra.Command) (fx.Option, error) {
	connectionOptions, err := bunconnect.ConnectionOptionsFromFlags(cmd.Context())
	if err != nil {
		return nil, err
	}

	return storage.Module(*connectionOptions), nil
}

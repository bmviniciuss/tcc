package factories

import (
	"log"
	"os"

	grpccardapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/card/grpc"
	httpcardapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/card/http"
	grpcclientwallet "github.com/bmviniciuss/tcc/card-payment/src/adapters/client_wallet/grpc"
	httpclientwalletapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/client_wallet/http"

	postgrespaymentrepository "github.com/bmviniciuss/tcc/card-payment/src/adapters/payment"
	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCardApi() payment.CardAPI {
	log.Println("Creating CardAPI based on the GRPC_ENABLED environment variable")
	e := os.Getenv("GRPC_ENABLED")

	if e == "true" {
		log.Println("Creating a gRPC card API")
		host := os.Getenv("CARD_HOST")
		grpcConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatal("Error connecting to card gRPC server")
		}

		return grpccardapi.NewGRPCardAPI(grpcConn)
	}
	log.Println("Creating a HTTP card API")
	return httpcardapi.NewHTTPCardAPI()

}

func NewClientWalletAPI() payment.ClientWalletTransactionAPI {
	log.Println("Creating ClientWalletAPI based on the GRPC_ENABLED environment variable")
	e := os.Getenv("GRPC_ENABLED")
	if e == "true" {
		log.Println("gRPC client wallet API is not implemented yet")
		host := os.Getenv("CLIENT_WALLET_HOST")
		grpcConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatal("Error connecting to CLIENT WALLET gRPC server")
		}

		return grpcclientwallet.NewGRPCClientWalletAPI(grpcConn)
	}
	log.Println("Creating a HTTP client wallet API")
	return httpclientwalletapi.NewHTTPClientWalletTransactionAPI()
}

func NewPaymentService(db *sqlx.DB) payment.Service {
	return payment.NewPaymentService(NewCardApi(), postgrespaymentrepository.NewPostgresPaymentRepository(db), NewClientWalletAPI())
}

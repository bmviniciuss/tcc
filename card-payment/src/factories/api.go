package factories

import (
	"log"
	"os"

	grpccardapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/card/grpc"
	httpcardapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/card/http"
	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
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

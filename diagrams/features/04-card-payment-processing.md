title Processamento de Pagamentos de Cartão

actor Cliente
participant Gateway
participant Serviço de Pagamentos
participant Serviço de Cartões
participant Camada de Dados

Cliente->Gateway: Requisição REST
Gateway->Serviço de Pagamentos:Requisição\nREST/gRPC
Serviço de Pagamentos->Serviço de Cartões:Busca Cartão\nREST/gRPC
Serviço de Cartões-->Serviço de Pagamentos: Informações do Cartão

Serviço de Pagamentos->Serviço de Cartões:Autorização de Pagamento\nREST/gRPC
Serviço de Cartões-->Serviço de Pagamentos: Autorização

Serviço de Pagamentos->Serviço de Pagamentos: Processamento de Pagamento
Serviço de Pagamentos->Camada de Dados: Inserção do Pagamento
Camada de Dados-->Serviço de Pagamentos: Informações do Pagamento
Serviço de Pagamentos-->Gateway:Informações do Pagamento
Gateway-->Cliente: Informações do Pagamento

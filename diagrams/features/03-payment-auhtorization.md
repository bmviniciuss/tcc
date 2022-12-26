title Autorização de Pagamentos de Cartão

actor Cliente
participant Gateway
participant Serviço de Cartões
participant Camada de Dados

Cliente->Gateway: Requisição REST
Gateway->Serviço de Cartões:Requisição\nREST/gRPC
Serviço de Cartões->Camada de Dados:Busca de Cartão por token
Camada de Dados-->Serviço de Cartões: Cartão
Serviço de Cartões->Serviço de Cartões: Regras de Autorização
Serviço de Cartões->Camada de Dados: Insere Resultado de Autorização
Camada de Dados-->Serviço de Cartões: Autorização de Pagamento
Serviço de Cartões-->Gateway: Autorização de Pagamento
Gateway-->Cliente: Autorização de Pagamento

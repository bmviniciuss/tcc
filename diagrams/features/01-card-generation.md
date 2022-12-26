title Geração Cartões

actor Cliente
participant Gateway
participant Serviço Cartões
participant Camada de Dados

Cliente->Gateway: Requisição REST
Gateway->Serviço Cartões: Requisição\nREST/gRPC
Serviço Cartões->Camada de Dados:Inserção de Cartão
Camada de Dados-->Serviço Cartões: Cartão
Serviço Cartões-->Gateway: Cartão
Gateway-->Cliente: Cartão

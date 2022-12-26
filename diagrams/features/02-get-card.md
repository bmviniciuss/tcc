title Busca de Cartão

actor Cliente
participant Gateway
participant Serviço Cartões
participant Camada de Dados

Cliente->Gateway: Requisição REST
Gateway->Serviço Cartões:Requisição\nREST/gRPC
Serviço Cartões->Camada de Dados:Busca de Cartão por token
Camada de Dados-->Serviço Cartões: Cartão
Serviço Cartões-->Gateway: Informações do Cartão
Gateway-->Cliente: Informações do Cartão

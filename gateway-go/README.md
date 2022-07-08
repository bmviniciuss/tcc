# Gateway

- [Gateway](#gateway)
- [Definição](#definição)
- [Operações](#operações)

# Definição
Porta de entrada do sistema que expõem rotas públicas para consumo de clientes. Se conecta aos micros serviços internos para realizar as operações.

Apenas é exposto uma API do tipo HTTP com endpoints que chamam internamente os micros serviços de cartão e/ou pagamento de cartão através de requisições HTTP ou gRPC.

# Operações
- [Criar Cartão](../card/README.md#geração-de-cartões)
- [Criar Pagamento Cartão](../card-payment/README.md#criar-pagamento)
- [Listar pagamentos cliente](../card-payment/README.md#busca-de-pagamentos-de-cliente)

# Q2Bank

Esse serviço faz o cadastro de usuários comuns, cadastro de lojistas, possibilita depositar dinheiro em sua carteira e também realizar transaferências.

# Rodando local

Pode ser iniciado pelo comando `make run` ou pelo comando `go run main.go`.

# ⚙️ Testes

Este serviço conta com uma cobertura de testes unitários alta.
Para os testes unitários da API, temos dois comandos principais:

- `make mock`: cria as implementações das interfaces, com o objetivo de realização da injeção de dependência para execução dos testes unitários. Para criá-los basta adicionar no arquivo **Makefile**, dentro da seção **mock**, a seguinte instrução: **mockgen -souce=./path_do_arquivo.go -destination=./path-onde-ficam-os-mocks.go -package=pasta-onde-ficam-os-mocks -mock_name=PastaDaInterface=NomeDoMock
- `make test`: executa os testes e apresenta o percentual de cobertura.

# Migration

Para subir as tabelas do banco é necessário rodar manualmente os scripts que estão no path `./migrations/script.sql`.

# Documentação

É inicializada junto com a inicialização da API, após rodar o comando `make run` ou o `go run main.go`, é só acessar o link [Documentação](http://localhost:1323/swagger/index.html#/).


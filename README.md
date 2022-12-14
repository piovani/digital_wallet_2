# Digital Wallet 2

## Technologies Used
- <a src="https://echo.labstack.com/">Echo</a>
- <a src="https://gorm.io/index.html/">Gorm</a>
- <a src="https://dev.mysql.com/doc/mysql-getting-started/en/">Mysql</a>
- <a src="https://circleci.com/docker/?utm_source=google&utm_medium=sem&utm_campaign=sem-google-dg--latam-en-nbAuth-maxConv-auth-nb&utm_term=g_e-docker_c__rsa2_20210709&utm_content=sem-google-dg--latam-en-nbAuth-maxConv-auth-nb_keyword-text_eta-docker_exact-&gclid=Cj0KCQjwgO2XBhCaARIsANrW2X341zChIvWoLkZo2VPSU0w82FmkA-USTf6F4rl7gdmoVz6yK4XWe6waAkWUEALw_wcB">Docker</a>
    - <a src="https://hub.docker.com/layers/redis/library/redis/latest/images/sha256-5050c3b85c308ec9e9eafb8ac7b3a8742a61cdb298d79851141a500491d45baf?context=explore">Redis</a>
    - <a src="https://hub.docker.com/layers/golang/library/golang/1.19-alpine3.16/images/sha256-70df3b8f9f099da7f60f0b32480015165e3d0b51bfacf9e255b59f3dd6bd2828?context=explore">Golang</a>
    - <a src="https://hub.docker.com/layers/mysql/library/mysql/5.7/images/sha256-5ecf646122c4fcbda81983c9e93e81a011b0593c9c19fbfc55b48bd1c23bc790?context=explore">Mysql</a>
- <a src="https://github.com/stretchr/testify">Testify<a>
- <a src="https://github.com/swaggo/swag#declarative-comments-format">Swag</a>
- <a src="https://github.com/swaggo/echo-swagger">Echo Swagger</a>

## Requirements
- Docker 20.10.17 or superior
- Docker Compose 1.29.2 or superior

## How Star The Project
1. Start containers Docker
```
docker-compose up -d
```
2. Run migrate
```
make migrate
```
3. Run seed
```
make seed
```

## How to Test
### Unit testing
To run unit tests, execute this command:
```
make cover
```
### With Swagger
1. Start API
```
make api
```
2. Open your Brozer at: http://localhost:8080/swagger/index.html#/coin/get_coin_price

### With <a src="https://www.postman.com/downloads/?utm_source=postman-home">Postman</a> or <a src="https://insomnia.rest/download">Insomnia</a>
1. Install Postman or Insomnia

2. import File Collections
```
./docs/Digital Wallet 2.postman_collection.jsondocs/Digital Wallet 2.postman_collection.json
```
3. Use the endpoints

## Checklist TODO
- [X] Conexao com banco de dados MySQL
- [X] Rota de Valores de moedas atual
- [X] Remover container de consulta de valores no redis
- [X] Criar Wallet do usu??rio
- [X] Ao depositar alterar a wallet tbm
- [X] Ao debitar consultar a wallet se ?? possivel
- [X] Valida????o de dados -> Rota de Deposito
- [X] Valida????o de dados -> Rota de Saque
- [X] Rota Historico
- [X] Rota Balan??o - Apresentar valor em carteira, valor em moeda atual, valor em dolares e valore em euros e a somatoraria dessas duas moedas principais
- [X] Testes unit??rios
- [X] Swagger com as rotas mapeadas
- [X] Seed com alguns dados iniciais
- [X] Centralizar os Log's 
- [X] Colocar Prometheus
- [X] Colocar Grapfana
- [X] Arrumar os containers para API rodar corretamente
- [X] Colocar Cobra para gerenciar o CMD
- [X] Melhorar o Migrate
- [X] Melhorar o Seed
- [ ] Melhorar a Coleta de values dos coins
- [ ] Melhorar a leitura dos values dos coins
- [ ] Colocar o Sleep de Collect como variavel de Ambiente
- [ ] CI com job para SonarCloud
- [ ] Teste de Integra????o
- [ ] Testes no Domain, j?? que o mesmo possui regras de cria????o
- [ ] Colocar o Valor na Wallet em Dolares e Euros na hora do Balan??o
- [ ] Criar Server gRPC
- [ ] Add Kafka
    - [ ] Na opera????o de Deposit, publicar uma mensagem no topic do kafka
    - [ ] Na opera????o de Withdraw, publicar uma mensagem no topic do kafka
    - [ ] Novo container para receber e printar as mensagens recebidas pelo kafka
- No Database criar: Transaction, Commit e Rollback
- [ ] Criar rotas que usem graphQL
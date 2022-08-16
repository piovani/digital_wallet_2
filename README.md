# Digital Wallet 2

## Technologies Used
- <a src="https://echo.labstack.com/">Echo</a>
- <a src="https://gorm.io/index.html/">Gorm</a>
- <a src="https://dev.mysql.com/doc/mysql-getting-started/en/">Mysql</a>
- <a src="https://circleci.com/docker/?utm_source=google&utm_medium=sem&utm_campaign=sem-google-dg--latam-en-nbAuth-maxConv-auth-nb&utm_term=g_e-docker_c__rsa2_20210709&utm_content=sem-google-dg--latam-en-nbAuth-maxConv-auth-nb_keyword-text_eta-docker_exact-&gclid=Cj0KCQjwgO2XBhCaARIsANrW2X341zChIvWoLkZo2VPSU0w82FmkA-USTf6F4rl7gdmoVz6yK4XWe6waAkWUEALw_wcB">Docker</a>
    - <a src="https://hub.docker.com/layers/redis/library/redis/latest/images/sha256-5050c3b85c308ec9e9eafb8ac7b3a8742a61cdb298d79851141a500491d45baf?context=explore">Redis</a>
    - <a src="https://hub.docker.com/layers/golang/library/golang/1.19-alpine3.16/images/sha256-70df3b8f9f099da7f60f0b32480015165e3d0b51bfacf9e255b59f3dd6bd2828?context=explore">Golang</a>
    - <a src="https://hub.docker.com/layers/mysql/library/mysql/5.7/images/sha256-5ecf646122c4fcbda81983c9e93e81a011b0593c9c19fbfc55b48bd1c23bc790?context=explore">Mysql</a>

## Requirements
- Docker 20.10.17 or superior
- Docker Compose 1.29.2 or superior

## How Star The Project

## How to Test

## Checklist TODO
- [X] Conexao com banco de dados MySQL
- [X] Rota de Valores de moedas atual
- [X] Remover container de consulta de valores no redis
- [X] Criar Wallet do usuário
- [X] Ao depositar alterar a wallet tbm
- [X] Ao debitar consultar a wallet se é possivel
- [X] Validação de dados -> Rota de Deposito
- [X] Validação de dados -> Rota de Saque
- [] Rota Historico
- [] Rota Balanço - Apresentar valor em carteira, valor em moeda atual, valor em dolares e valore em euros e a somatoraria dessas duas moedas principais
- [] Testes unitários
- [] Swagger
- [] Base de dados com algumas informações
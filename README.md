<h1 align="center"> 🌬️ Venture </h1>

<h1 align="center"> We are security, speed, and technology. We are Venture </h1>

<p align="center">
  <img style="width:100px; height:100px; border-radius:20%;" src="https://i.imgur.com/yieDOSJ.png"/>
</p>

<div align="center">

![Go - language](https://img.shields.io/badge/language-go-cyan)
![Postgres - Database](https://img.shields.io/badge/database-postgres-blue)
![AWS - Cloud](https://img.shields.io/badge/cloud-aws-yellow)


</div>

### 🛢 Migrations

- Please, install golang-migration

> Linux
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

> Windows
```bash
$ scoop install migrate
```

- To create a new migration
```bash
migrate create -ext sql -dir database/migrations description-of-migration
```
> f.e: migrate create -ext sql -dir database/migrations add_profile_image_to_drivers

- To run your migration
```bash
migrate -path=database/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" up
```

- If you need run rollback
```bash
migrate -path=database/migrations -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" down
```

- Case, for some reason, need run migration from the beginning
```bash
migrate create -ext=sql -dir=database/migrations -seq init
```

# Architecture

```
|__ cmd --> pasta onde a inicialização das apis, rotinas e consumers devem estar
    |__ api --> onde a inicialização apenas de apis devem estar
        |__ server --> pasta onde a inicialização do nosso server principal e das configurações dele, devem estar
            |__ routes --> onde fica a inicialização das rotas usadas no server
                |__ v1 --> rotas da v1 e inicialização
    |__ consumers --> aqui, é onde fica nossos consumers, rotinas que rodam permanentemente ouvindo filas kafka e fazendo algo, como por exemplo o uchiha_consumer que fica ouvindo, uma fila do kafka e tudo que chega ele envia um email.
|__ config --> pasta onde fica todas as nossas envs
|__ database --> onde guardamos todas as migrations, migrations sao alterações momentaneas que vao ficando no historico ate definir o banco como esta hoje
    |__ migrations --> pasta onde ficam as migrations. Para criar uma nova migrations veja acima na documentação
|__ http --> pasta usada para guardar testes https diretos, usado para teste local.
|__ internal --> pasta do core da aplicação, aqui são onde as coisas acontecem. É algo privado e não pode ser exportado, é nosso verdadeiro código.
    |__ controller --> aqui são os controladores das rotas, neles deixamos explicitos quem é o handler/controller responsavel por aquela rota, qual usecase ele vai chamar, o que ele vai receber de body e qual vai ser a resposta final em json, para o frontend ou usuário.
    |__ domain --> faz parte essencial do core da aplicação, sendo responsável pela lógica de negócios e pelas regras centrais.
        |__ repository --> o repository é onde deixamos os contratos (que em código viram interfaces) do que cada tipo de usuário deve fazer. 
        |__ service --> já os services, são os serviços que utilizamos, para validar e funcionar essas regras de negócio
            |__ adapters --> aqui são onde criamos as interfaces/contratos de parceiros 
            |__ addresses --> aqui é onde aplicamos o contrato para poder usar funções que trabalham com distância e endereços
            |__ auth --> aqui é onde fazemos a autenticação do nosso usuário de maneira única com apenas uma rota, independente do tipo de usuário
            |__ decorator --> aqui é onde implementamos o decorator, o decorator ou decorador serve justamente pra decorar uma resposta, tal como um cache e utilizando cache em alguns casos
            |__ middleware --> já no middleware é a pasta que usamos para validar se o usuário está mesmo logado ainda.
            |__ payments --> aqui é onde realizamos todas as funções de pagamento, integrado com a plataforma do parceiro de pagamento.
    |__ entity --> aqui são todos os nossos tipos de usuários transformados em objetos.
    |__ exceptions --> aqui estão erros que enviamos para os controllers.
    |__ infra --> na pasta de infraestrutura, estão todos contratos e implementações do que cada ferramenta faz pra gente, por exemplo o cache, ele só faz get e set, então temos um contrato na pasta de cache que força isso. 
        |__ bucket --> aqui está a implementação de todas as funções do s3 da aws que usamos, para bucket
        |__ cache --> aqui está a implementação de todas as funções do redis que usamos para cache
        |__ contracts --> aqui estão todos os contratos de todos os componentes de infraestrutura, nas demais pastas são a implementação daquela ferramenta
        |__ database --> aqui está a implementação de todas as funções do postgres que usamos como banco de dados
        |__ email --> aqui está a implementação de todas as funções do ses da aws que usamos para envio de email
        |__ logger --> aqui está a implementação de todas as funções do logger da uber que usamos para logar
        |__ persistence --> persistence é a pasta mais diferente por que ela implementa funções do contrato do repository, que fica lá em domain
        |__ app.go --> esse arquivo demonstra tudo que temos na nossa aplicação de maneira pública para que possa ser usado por todos os arquivos, quebrando um pouco da injeção de dependência.
    |__ setup --> na pasta de setup, é onde iniciamos todas as conexões e mostramos quem implementa cada interface
    |__ usecase --> aqui são todos os casos de uso, os casos de uso são chamados apenas pelos controllers como mostra uma imagem abaixo desta seção de explicação e cada caso de uso, é de fato uma funcionalidade do usuário
    |__ value --> aqui são nossos contratos de resposta da api, eles não são interfaces mas são as respostas de json em forma de struct, elas existem para facilitar o que devemos devolver sem ter que ficar usando milhares e multiplas tags de json dentro das entidades
|__ mocks --> esta pasta tem como objetivo mockar todas as interfaces para facilitar os testes, utilizamos o mockery
|__ pkg --> aqui estao todas as funções que podem ser reutilizadas por todo código, sem depender dele
    |__ realtime --> tudo que se refere a tempo está nessa pasta, ao invés de usar o now do Go, você pode usar o now do realtime que vai ser o tempo de São Paulo
    |__ utils --> aqui está a pasta de funções pequenas que podem ser utilizadas por todo código
|__ scripts --> aqui está a pasta utilizada para rodas scripts momentâneos, exemplo: script para enviar uma mensagem para fila de email e testar o uchiha_consumer
|__ .gitignore --> tudo que está dentro desse arquivo será ignorado na subida do git
|__ .golangci.yml --> aqui estão as configurações de linter do repositório, se você quiser configurar um linter
|__ nginx.conf --> aqui são as configurações do nosso nginx para rodar local
|__ venture-api.service --> arquivo de configuração para build do servidor
```

<p align="center">
  <br>
  <img src="https://zup.com.br/wp-content/uploads/2021/10/Clean-Architecture-3.png"/>
</p>
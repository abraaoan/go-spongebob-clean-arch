# 🧽 spongebob-service

Microserviço em Go responsável por armazenar e expor informações da série **SpongeBob SquarePants**. O projeto segue os princípios da **Clean Architecture**, com testes, cache por ID, e persistência em PostgreSQL.

## 🧠 Arquitetura

Estrutura de diretórios:

```text
spongebob-service/  
├── cmd/                → Entrada da aplicação (main)  
├── internal/  
│   ├── entity/         → Entidades do domínio (Character, Episode, Quote, Season)  
│   ├── usecase/        → Casos de uso (interfaces + implementações)  
│   ├── repository/     → Abstrações de persistência  
│   ├── infrastructure/ → Implementações (PostgreSQL, Cache)  
│   └── delivery/       → Camada de entrega (HTTP handlers)  
├── pkg/                → Pacotes auxiliares (logger, utils)  
└── test/               → Testes de integração e mocks  
```

## 🧩 Funcionalidades

- [x] Cadastro de personagens (`Character`)  
- [x] Cadastro de episódios (`Episode`)  
- [x] Cadastro de temporadas (`Season`)  
- [x] Cadastro de citações (`Quote`)  
- [x] Listagem e busca por ID  
- [x] Cache com expiração (somente leitura)  
- [ ] Testes unitários e de integração  

## 💾 Banco de Dados

- PostgreSQL  
- Scripts de criação disponíveis em `/migrations` *(a definir)*  
- Estrutura baseada em relações entre:
  - `characters`
  - `episodes`
  - `seasons`
  - `quotes`

## 🚀 Como rodar

Pré-requisitos: Go 1.22+, PostgreSQL local

1. Clone o projeto:
   git clone https://github.com/seuuser/spongebob-service.git && cd spongebob-service

2. Instale dependências:
   go mod tidy

3. Execute localmente:
   go run cmd/main.go

## 🧪 Testes

Execute todos os testes:
   go test ./...

## 🛠️ To-Do

- [ ] Dockerfile + docker-compose  
- [ ] Documentação da API (Swagger)  
- [ ] Monitoramento (healthcheck, logs)  
- [ ] Autenticação futura?  

## 🤝 Contribuição

Abra uma issue, envie um PR ou mande uma citação engraçada do Bob Esponja 🧽💬

## 📜 Licença

MIT. Livre para usar, modificar e rir com responsabilidade 😄

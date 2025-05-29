# Backend Go Bet

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.9.1-lightblue.svg)](https://gin-gonic.com)

Este é um projeto backend desenvolvido em Go, utilizando uma arquitetura limpa e moderna para gerenciamento de apostas.

## 🚀 Tecnologias Utilizadas

- Go 1.21
- Gin (Framework Web)
- GORM (ORM)
- PostgreSQL
- Redis
- Go Validator

## 📁 Estrutura do Projeto

```
src/
├── cmd/         # Ponto de entrada da aplicação
├── config/      # Configurações do projeto
├── controller/  # Controladores da aplicação
├── middleware/  # Middlewares
├── model/       # Modelos de dados
├── routes/      # Definição de rotas
└── util/        # Utilitários
```

## 🛠️ Pré-requisitos

- Go 1.21 ou superior
- PostgreSQL
- Redis

## 🔧 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/lfdelima3/Backend-Go-Bet.git
```

2. Entre no diretório do projeto:
```bash
cd Backend-Go-Bet
```

3. Instale as dependências:
```bash
go mod download
```

4. Configure as variáveis de ambiente necessárias (crie um arquivo .env baseado no .env.example)

5. Execute o projeto:
```bash
go run src/cmd/main.go
```

## 📚 Documentação da API

### Endpoints Principais

#### Autenticação
- `POST /api/v1/auth/login` - Login de usuário
- `POST /api/v1/auth/register` - Registro de novo usuário
- `POST /api/v1/auth/refresh` - Renovação de token

#### Apostas
- `GET /api/v1/bets` - Listar todas as apostas
- `POST /api/v1/bets` - Criar nova aposta
- `GET /api/v1/bets/:id` - Obter detalhes de uma aposta
- `PUT /api/v1/bets/:id` - Atualizar uma aposta
- `DELETE /api/v1/bets/:id` - Remover uma aposta

### Exemplo de Uso

```bash
# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "usuario@exemplo.com", "password": "senha123"}'

# Criar Aposta
curl -X POST http://localhost:8080/api/v1/bets \
  -H "Authorization: Bearer seu-token-jwt" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 100.00,
    "type": "sports",
    "description": "Aposta em futebol"
  }'
```

## 🔐 Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=bet_db

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

JWT_SECRET=seu_segredo_jwt
JWT_EXPIRATION=24h

SERVER_PORT=8080
```

## 🧪 Testes

Para executar os testes do projeto:

```bash
# Executar todos os testes
go test ./...

# Executar testes com cobertura
go test ./... -cover

# Executar testes específicos
go test ./src/controller/...
```

## 📦 Build e Deploy

Para criar um build do projeto:

```bash
# Build para Linux
GOOS=linux GOARCH=amd64 go build -o bet-backend ./src/cmd/main.go

# Build para Windows
GOOS=windows GOARCH=amd64 go build -o bet-backend.exe ./src/cmd/main.go
```

## 👥 Contribuição

1. Faça um Fork do projeto
2. Crie uma Branch para sua Feature (`git checkout -b feature/AmazingFeature`)
3. Faça o Commit de suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Faça o Push para a Branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Convenções de Código

- Siga o [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` para formatação do código
- Escreva testes para novas funcionalidades
- Mantenha a documentação atualizada
- Use nomes descritivos para variáveis e funções

## 🔄 CI/CD

O projeto utiliza GitHub Actions para CI/CD. Os workflows incluem:
- Testes automatizados
- Análise de código estático
- Build automático
- Deploy em ambiente de staging

## 📧 Contato

Luis Fernando - [@lfdelima3](https://github.com/lfdelima3) 
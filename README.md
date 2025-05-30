# 🎮 Backend Go Bet

API RESTful para sistema de apostas esportivas desenvolvida em Go, com foco em performance e segurança.

## ✨ Características

- 🔐 Autenticação JWT com refresh token
- 🚀 Cache com Redis para melhor performance
- ⚡ Rate Limiting para proteção contra abusos
- 📝 Logging estruturado com Zap
- 📚 Documentação Swagger
- ✅ Validação de dados robusta
- 🛡️ Tratamento de erros centralizado
- 🗄️ Migrações automáticas do banco de dados
- 🔄 Middleware de cache inteligente
- 👮‍♂️ Middleware de autenticação e autorização
- 🎯 Validação de dados com mensagens personalizadas
- 🔍 Busca avançada com filtros
- 📊 Estatísticas em tempo real
- 🔔 Notificações em tempo real
- 📱 API RESTful com padrões REST
- 🔒 Segurança reforçada
- 📈 Monitoramento de performance

## 🛠️ Requisitos

- Go 1.21 ou superior
- PostgreSQL 15 ou superior
- Redis 7 ou superior

## 🚀 Configuração

1. Clone o repositório:
```bash
git clone https://github.com/lfdelima3/Backend-Go-Bet.git
cd Backend-Go-Bet
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

4. Execute a aplicação:
```bash
go run src/cmd/main.go
```

## 📚 Documentação da API

A documentação Swagger está disponível em:
- Desenvolvimento: `http://localhost:8080/swagger/index.html`
- Produção: `https://api.seusite.com/swagger/index.html`

### Endpoints Principais

#### Autenticação
- `POST /api/v1/auth/login` - Login de usuário
- `POST /api/v1/auth/refresh` - Refresh token
- `POST /api/v1/auth/logout` - Logout

#### Usuários
- `GET /api/v1/users` - Listar usuários
- `POST /api/v1/users` - Criar usuário
- `GET /api/v1/users/:id` - Obter usuário
- `PUT /api/v1/users/:id` - Atualizar usuário
- `DELETE /api/v1/users/:id` - Deletar usuário

#### Apostas
- `GET /api/v1/bets` - Listar apostas
- `POST /api/v1/bets` - Criar aposta
- `GET /api/v1/bets/:id` - Obter aposta
- `PUT /api/v1/bets/:id` - Atualizar aposta
- `DELETE /api/v1/bets/:id` - Deletar aposta

#### Partidas
- `GET /api/v1/matches` - Listar partidas
- `POST /api/v1/matches` - Criar partida
- `GET /api/v1/matches/:id` - Obter partida
- `PUT /api/v1/matches/:id` - Atualizar partida
- `DELETE /api/v1/matches/:id` - Deletar partida

## 📁 Estrutura do Projeto

```
src/
├── cmd/          # Ponto de entrada da aplicação
├── config/       # Configurações e variáveis de ambiente
├── controller/   # Controladores da API
├── middleware/   # Middlewares (auth, cache, rate limit)
├── model/        # Modelos e entidades
├── routes/       # Definição de rotas
└── util/         # Utilitários e helpers
```

### Descrição Detalhada

#### `cmd/`
- `main.go` - Ponto de entrada da aplicação
- Configuração do servidor
- Inicialização de dependências

#### `config/`
- Configurações do banco de dados
- Configurações do Redis
- Configurações do servidor
- Variáveis de ambiente

#### `controller/`
- Lógica de negócios
- Manipulação de requisições
- Respostas HTTP
- Validação de dados

#### `middleware/`
- Autenticação
- Cache
- Rate limiting
- Logging
- Validação

#### `model/`
- Definição de entidades
- Relacionamentos
- Validações
- Hooks

#### `routes/`
- Definição de rotas
- Grupos de rotas
- Middlewares específicos

#### `util/`
- Funções auxiliares
- Constantes
- Tipos personalizados
- Validações

## 🧪 Testes

Execute os testes com:
```bash
go test ./...
```

Para ver a cobertura de testes:
```bash
go test ./... -cover
```

Para executar testes específicos:
```bash
go test ./src/controller -v
go test ./src/model -v
```

## 🔧 Variáveis de Ambiente

Principais variáveis de ambiente necessárias:

```env
# Servidor
PORT=8080
ENV=development

# Banco de Dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=betting_db

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRATION=24h

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_DURATION=1m

# Cache
CACHE_DURATION=5m
```

## 🔍 Funcionalidades Detalhadas

### Autenticação
- Login com email/senha
- Refresh token automático
- Logout com invalidação de token
- Proteção de rotas
- Roles e permissões

### Cache
- Cache de respostas HTTP
- Cache de consultas frequentes
- Invalidação automática
- TTL configurável

### Rate Limiting
- Limite por IP
- Limite por usuário
- Limite por rota
- Headers informativos

### Logging
- Logs estruturados
- Níveis de log configuráveis
- Rotação de logs
- Contexto de requisição

### Validação
- Validação de dados
- Mensagens personalizadas
- Validações customizadas
- Sanitização de dados

## 🤝 Contribuindo

1. Fork o projeto
2. Crie sua branch de feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

### Padrões de Código

- Siga o [Effective Go](https://golang.org/doc/effective_go)
- Use [gofmt](https://golang.org/cmd/gofmt/) para formatação
- Escreva testes para novas funcionalidades
- Documente funções e tipos
- Mantenha a cobertura de testes alta

## 📝 Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 👥 Autores

- **Luis Fernando Antunes de Lima** - [lfdelima3](https://github.com/lfdelima3)
- **Ruan Henrique Brunhera Aronchi** - [Ruan246Etec]
(https://github.com/RuanBrunhera)

## 🙏 Agradecimentos

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Redis Go Client](https://github.com/redis/go-redis)
- [JWT Go](https://github.com/golang-jwt/jwt)
- [Zap Logger](https://github.com/uber-go/zap)
- [Validator](https://github.com/go-playground/validator)
- [Swagger](https://github.com/swaggo/swag)

## 📞 Suporte

Para suporte, envie um email para lfdelimaa@gmail ou abra uma issue no GitHub.

## 🔄 Atualizações

- **v1.0.0** - Lançamento inicial
- **v1.1.0** - Adicionado sistema de cache
- **v1.2.0** - Melhorias na autenticação
- **v1.3.0** - Adicionado rate limiting 
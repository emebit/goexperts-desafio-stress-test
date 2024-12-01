# Goexperts-Stress-Test

Um sistema de teste de carga simples em Go para realizar testes de desempenho em serviços web.

## Como usar

### Pré-requisitos

- Go (1.22 ou maior) instalado
- Docker (opcional)

### Executando localmente

1. Clone o repositório:

```bash
    git clone git@github.com:emebit/goexperts-desafio-stress-test.git
    cd goexperts-desafio-stress-test
```

2. Compile e execute o programa:

```bash
    go run cmd/main.go --url<URL> --requests=<NÚMERO_DE_requests> --concurrency=<NÚMERO_DE_CONCORRÊNCIA>
```

### Executando com Docker

Você pode construir uma imagem docker e executar a aplicação.

1. Build da imagem:

```bash
    docker build -t goexperts-desafio-stress-test .
```

2. Execute a imagem:

```bash
    docker run goexperts-desafio-stress-test --url=https://google.com.br/ --requests=1000 --concurrency=10
```

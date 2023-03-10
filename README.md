# Artefatos TCC

Esse repositório foi criado para distribuir, para fins de teste e experimentação, os artefatos produzidos durante a escrita do Texto de Conclusão de Curso com o objetivo de realizar um estudo comparativo qualitativo entre ferramentas de monitoramento de contêineres Docker.

## Como Usar

1. Clone o repositório para o seu computador:

```bash
git clone https://github.com/lucasalustiano/TCC
```

2. Entre no diretório do repositório:

```bash
cd TCC
```

3. Construa e execute os contêineres da API e do Prometheus com o Docker Compose

```bash
docker-compose -f docker-compose-prometheus.yml up
```

ou ainda, construa e execute os contêineres da API e do NetData com o Docker Compose. Lembrando que você deve ter um token próprio fornecido pelo NetData, para isso, renomeie o arquivo .env.example para .env e insira o token na variavel de ambiente NETDATA_CLAIM_TOKEN. Após isso, execute:

```bash
docker-compose -f docker-compose-netdata.yml up
```

4. O Prometheus estará acessível em http://localhost:9090 e o NetData em http://localhost:19999 ou na própria página web do NetData.

5. Para realizar o experimento com simulação de carga, execute o script de simulação:

```bash
go run script_simulacao_carga.go
```

O script aceita dois parâmetros passados como flags:

- -n: Número de requisições a serem feitas para cada endpoint (padrão número randômico)
- -s: Número em segundos de espera entre as requisições (padrão 15 segundos)

E com isso deverá ser possível observar os dados sendo coletados e expostos pelo Prometheus e/ou pelo NetData.

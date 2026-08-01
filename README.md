# 🚀 Desafio Técnico Korp

Projeto desenvolvido como solução para o **Desafio Técnico Korp**, contemplando a construção de uma aplicação em **Go**, containerização com **Docker**, proxy reverso com **NGINX**, monitoramento com **Prometheus** e **Grafana**, publicação segura utilizando **HTTPS (Let's Encrypt + DuckDNS)** e automação completa do ambiente com **Ansible**.

---

# Arquitetura

```text
                    Internet
                         │
                         ▼
             desafio-korp.duckdns.org
                         │
                    HTTPS (443)
                         │
                         ▼
                    NGINX Reverse Proxy
                         │
        ┌────────────────┼────────────────┐
        │                │                │
        ▼                ▼                ▼
 HTTP Server         Grafana         Prometheus
    (Go)             /grafana       /prometheus
```

---

# Tecnologias utilizadas

- Golang
- Docker
- Docker Compose
- NGINX
- Prometheus
- Grafana
- Ansible
- Let's Encrypt
- DuckDNS

---

# Funcionalidades

- Aplicação HTTP escrita em Go
- Endpoint REST retornando JSON
- Health Check
- Endpoint de métricas Prometheus
- Proxy Reverso com NGINX
- Dashboard Grafana provisionado automaticamente
- HTTPS utilizando Let's Encrypt
- Automação completa utilizando Ansible
- Deploy reproduzível através do Docker Compose

---

# Endpoints

| Endpoint | Descrição |
|----------|-----------|
| `/projeto-korp` | Endpoint principal |
| `/health` | Health Check |
| `/metrics` | Métricas Prometheus |
| `/grafana` | Dashboard Grafana |
| `/prometheus` | Interface Prometheus |

---

# Estrutura do Projeto

```text
desafio-korp
│
├── app
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
│
├── nginx
│   └── http-server-projeto-korp.conf
│
├── prometheus
│   └── prometheus.yml
│
├── grafana
│   ├── dashboards
│   └── provisioning
│
├── certbot
│   ├── conf
│   └── www
│
├── ansible
│   ├── inventory
│   ├── playbook.yml
│   ├── group_vars
│   └── roles
│
├── docker-compose.yml
│
└── README.md
```

---

# Monitoramento

O projeto disponibiliza métricas através do Prometheus e um dashboard Grafana provisionado automaticamente.

O dashboard apresenta:

- Total de requisições
- Disponibilidade da aplicação
- Requisições por segundo
- Tempo de resposta
- Uso de memória
- Número de Goroutines
- Informações da aplicação Go

---

# HTTPS

O ambiente utiliza:

- DuckDNS
- Let's Encrypt
- Certbot
- NGINX Reverse Proxy

Todo acesso HTTP é redirecionado automaticamente para HTTPS.

---

# Deploy

## Clonar o projeto

```bash
git clone git@github.com:ongamss/desafio-korp.git

cd desafio-korp
```

---

## Subir os containers

```bash
docker compose up -d --build
```

---

## Validar

Aplicação

```
https://desafio-korp.duckdns.org/projeto-korp
```

Health

```
https://desafio-korp.duckdns.org/health
```

Metrics

```
https://desafio-korp.duckdns.org/metrics
```

Grafana

```
https://desafio-korp.duckdns.org/grafana
```

Prometheus

```
https://desafio-korp.duckdns.org/prometheus
```

---

# Provisionamento automatizado

Todo o ambiente pode ser provisionado com um único comando utilizando Ansible.

```bash
ansible-playbook playbook.yml
```

O playbook realiza:

- Instalação do Docker
- Clonagem do projeto
- Build da aplicação Go
- Criação dos containers
- Configuração do NGINX
- Configuração do Prometheus
- Configuração do Grafana
- Configuração do HTTPS
- Validação automática da aplicação

---

# Docker

Containers executados:

| Container | Função |
|------------|---------|
| http-server-projeto-korp | Aplicação Go |
| nginx | Reverse Proxy |
| prometheus | Monitoramento |
| grafana | Dashboards |
| certbot | Certificados HTTPS |

---

# Observabilidade

Prometheus configurado com:

- retenção de 7 dias
- limite de armazenamento de 1 GB
- coleta automática das métricas da aplicação

Grafana provisionado automaticamente via arquivos YAML e JSON.

---

# Segurança

- HTTPS obrigatório
- TLS 1.2 / TLS 1.3
- HSTS
- Reverse Proxy
- Headers de segurança
- Certificados Let's Encrypt

---

# Autor

**Magno Silva**

DevOps Engineer

GitHub

https://github.com/ongamss

---

# Licença

Projeto desenvolvido exclusivamente para fins de avaliação técnica no processo seletivo da Korp.
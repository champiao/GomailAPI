# GoMail

Go Mail é uma API REST para envio de email com mais de uma configuração por environment, podendo utilizar para enviar emails de múltiplos sistemas com uma API centralizada.

## Utilização

create a .env File to configure your credentials
```bash
SMTP_HOST_FIRST=FirstHost.com
SMTP_PORT_FIRST=YourPort
SMTP_PASS_FIRST=FirstPassword
```

## Exemplo de requisição para API

```python
{
    "flag":"FIRST",
    "mailTo":"TargetEmail",
    "subject":"teste de envio de email",
    "message":"teste Novo email pela API GO"
}
```
## Build e Run do programa

```bash
go work init ./mail
go work use .
go mod tidy
cd mail
go get .
cd ..
go build
./main {para linux/mac}
main.exe {para Windows}
```

## Resultado

Este código sobe uma API REST sem autenticação com a função de disparo de email com sua rota /mail em requisição POST com o exemplo de requisição acima, usando text/plain para o corpo do email para poder usar desde string até HTML para montar as mensagens.

## License

[CHAMPIAO](https://champiao.com.br)

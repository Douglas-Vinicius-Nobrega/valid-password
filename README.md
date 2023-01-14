# Valid-Password

# 🚀 Começando

## Nesse passo, utilizaremos as plataformas Postman e Docker, certifique-se de ter elas instalados na sua máquina.

## Caso não tenha o docker, leia até o final que faremos em duas etapas.

## Passos

## OBS: utilizaremos uma plataforma de API chama da Postman, para testar nosso projeto

1 - Com o projeto aberto no VSCode, vamos abrir o terminal integrado "CTRL + J"

2 - Com o terminal aberto, vamos utilizar o seguinte comando, $ docker-compose up

3 - Assim que nossa máquina virtual estiver on-line, entraremos no Postman e colocamos a rota que foi determinada : localhost:8080/verify

4 - Dentro do Postman, selecionaremos o métodos POST (É utilizado para enviar dados para o servidor), selecionaremos o body e JSON

5 - Agora no Body, utilizaremos o seguinte formato :

{
"password": "!TesteSenháForteÀ!123&",
"rules": [
{
"rule": "minSize",
"value": 8
},
{
"rule": "minUppercase",
"value": 8
},
{
"rule": "minLowercase",
"value": 4
},
{
"rule": "minDigit",
"value": 2
},
{
"rule": "minSpecialChars",
"value": 4
},
{
"rule": "noRepeted",
"value": 2
}
]
}

6 - Agora é só clicar no botão "Send" e esperar o retorno

7 - Retorno esperado : {"verify":false,"NoMatch":["minDigit"]}

## Caso não utilizar ou não conheça a plataforma Docker, faremos em outro formato para testa nossa API

## Certifique-se de ter o Golang, instalado na sua máquina, para fazermos esse passo.

## OBS : No passo 2, em vez de utilizamos o docker-compose up utilizaremos:

2 - go run main.go

## OBS : Assim, subiremos nosso sevidor local e podemos entra no Postman e fazeremos nossos teste.

## O retorno esperado é o mesmo do passo 7.

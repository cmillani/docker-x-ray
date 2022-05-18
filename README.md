# Docker x-ray
Este é o repositório utilizado para as demonstrações da apresentação!

## PowerPoint
O arquivo pptx contém os slides utilizados na apresentação, e nos comentários de cada slides os comandos executados

## Pastas
Cada pasta representa uma seção da apresentação:

* ./hello-world
  Aqui exploramos um codigo simples em C, fazendo um build statico para gerar uma imagem menor
* ./go
  Aqui exploramos uma imagem distroless de um web server escrito em Go
* ./capture
  Captura das chamadas de sistema feitas pela stack do Docker, explorando a trilha do software
* ./runc
  Criamos um container utilizando apenas o runc, que é o que está de baixo dos panos na invocação do Docker! Nota: o `recvtty` precisa estar disponível, e pode ser obtido a partir do repositório do `runc`!
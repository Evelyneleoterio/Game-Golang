<h1 align="center">Go Gamer</h1>
<p align="center"><i>Segue abaixo as instruções de como executar o jogo e as tecnologias utilizadas durante a construção do projeto.</i></p>

## Sobre este projeto

Este é um jogo 2D desenvolvido em Go utilizando a biblioteca Ebitengine, uma ferramenta para criação de jogos multiplataforma. O projeto contém conceitos como colisões, spawn de objetos dinâmicos, manipulação de gráficos e lógica de pontuação.

#### Por que em Go?

Criei com o intuito de praticar o que já venho aprendendo com linguagem Go.

### Tecnologias

<p style="display: flex; gap: 10px;">
  <img width="48" src="https://miro.medium.com/v2/resize:fit:600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="go-logo"/>
  <img width="90" src="https://ebitengine.org/images/share.png" alt="ebitengine-logo"/>
</p>


### Ferramentas de desenvolvimento
<p style="display: flex; gap: 10px;">
  <img width="48" src="https://img.icons8.com/?size=100&id=20906&format=png&color=000000" alt="git-logo"/>
  <img width="48" src="https://img.icons8.com/?size=100&id=0OQR1FYCuA9f&format=png&color=000000" alt="vscode-logo"/>
  <img width="55" src="https://img.icons8.com/?size=100&id=22813&format=png&color=000000" alt="docker-logo"/>
</p>


### Pré-requisitos:

* Docker: Instalado e funcionando corretamente em seu sistema (Windows, macOS ou Linux).
* Go: Versão 1.23.4 instalada.
* Git: Para clonar o repositório.

## Como executar em cada sistema operacional

**1 Clone o Repositório:**
 - `git clone https://github.com/Evelyneleoterio/Game-Golang.git`

**2 Construa a Imagem Docker:**
- Abra o terminal no diretório do seu projeto.
- Execute o comando: `docker build -t mygame`

**3 Execute o Container:**
#### Para Linux:

<img width="48" src="https://img.icons8.com/?size=100&id=17842&format=png&color=000000" alt="linux-logo"/> 

-  `docker run -it --rm mygame `

#### Para Windows:
<img width="48" src="https://img.icons8.com/?size=100&id=gXoJoyTtYXFg&format=png&color=000000" alt="windows-logo"/>

- `docker run -it --rm --platform=linux/amd64 mygame`

#### Para MacOs:
<img width="48" src="https://img.icons8.com/?size=100&id=uoRwwh0lz3Jp&format=png&color=000000" alt="macos-logo"/>

- `docker run -it --rm --platform=linux/amd64 mygame .`

## visualização do jogo

![game GIF](assets/video.webm)


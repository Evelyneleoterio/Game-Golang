<h1 align="center">Go Gamer</h1>
<p align="center"><i>Segue abaixo as instruções de como executar o jogo e as tecnologias utilizadas durante a construção do projeto.</i></p>

## Sobre este projeto

Este é um jogo 2D desenvolvido em Go utilizando a biblioteca Ebitengine, uma ferramenta para criação de jogos multiplataforma. O projeto contém conceitos como colisões, spawn de objetos dinâmicos, manipulação de gráficos e lógica de pontuação.

#### Por que em Go?

Criei com o intuito de praticar o que já venho aprendendo na linguagem Go.

### Tecnologias

<p display="inline-block">
  <img width="48" src="https://miro.medium.com/v2/resize:fit:600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="go-logo"/>

  <img width="90" src="https://ebitengine.org/images/share.png" alt="ebitengine-logo"/>
</p>

### Ferramentas de desenvolvimento
<p>
 <img width="48" src="https://img.icons8.com/?size=100&id=20906&format=png&color=000000" alt="git-logo"/>

 <img width="48" src="https://img.icons8.com/?size=100&id=0OQR1FYCuA9f&format=png&color=000000" alt="vscode-logo"/>

 <img width="55" src="https://img.icons8.com/?size=100&id=22813&format=png&color=000000" alt="docker-logo"/>
</p>
## Como executar em cada sistema operacional

<img width="48" src="https://img.icons8.com/?size=100&id=17842&format=png&color=000000" alt="linux-logo"/>

- Clone o repositório no seu ambiente local com: `git clone https://github.com/Evelyneleoterio/Game-Golang.git`

**_OBS: é necessário ter o Docker instalado no seu Linux_**

- Abra o projeto no Visual Studio Code ou, no terminal, abra o diretório onde está o projeto e digite o comando `docker compose up`.
- Em seguida digite `go run main.go`
<img width="48" src="https://img.icons8.com/?size=100&id=gXoJoyTtYXFg&format=png&color=000000" alt="windows-logo"/>

- Clone o repositório no seu ambiente local com: `git clone https://github.com/Evelyneleoterio/Game-Golang.git`

**_OBS: é importante ter o Go na versão 1.23.4 no Windows_**

- Abra o Prompt de Comando ou PowerShell no Windows.
- Dentro do terminal, digite os seguintes comandos:
  - `set GOOS=windows`
  - `set GOARCH=amd64`
  - `go build -o mygame.exe`
    Isso cria um executável para Windows chamado `mygame.exe`.

<img width="48" src="https://img.icons8.com/?size=100&id=uoRwwh0lz3Jp&format=png&color=000000" alt="macos-logo"/>

- Clone o repositório no seu ambiente local com: `git clone https://github.com/Evelyneleoterio/Game-Golang.git`

**_OBS: é importante ter o Go na versão 1.23.4 no macOS_**

- No macOS, abra o terminal, entre no diretório e execute os seguintes comandos:
  - `set GOOS=darwin`
  - `set GOARCH=amd64`
  - `go build -o mygame-macos`

<video controls>
  <source src="assets/video.webm" type="video/webm">
  Seu navegador não suporta o formato de vídeo.
</video>



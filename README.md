# Go Socket Chat

Um chat via terminal (CLI) desenvolvido em **Go**, utilizando **Sockets TCP** nativos para comunicação cliente-servidor, estruturação de dados em **JSON** e uma Interface de Usuário de Terminal (TUI) construída com o framework **Bubble Tea**.

---

## Funcionalidades

* **Interface Moderna (TUI):** Desenvolvida com [Bubble Tea](https://github.com/charmbracelet/bubbletea), oferecendo uma experiência visual rica, fluida e interativa no terminal.
* **Comunicação TCP em Tempo Real:** Arquitetura cliente-servidor robusta utilizando o pacote nativo `net` do Go.
* **Protocolo Baseado em JSON:** Mensagens estruturadas rigorosamente em JSON para garantir escalabilidade, integridade e facilidade no parsing.
---

## Tecnologias Utilizadas

* **[Go (Golang)](https://golang.org/)** - Linguagem principal.
* **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Framework TUI (The Elm Architecture em Go).
* **TCP Sockets (`net`)** - Camada de transporte e comunicação de rede.
* **Encoding JSON** - Serialização de dados das mensagens.

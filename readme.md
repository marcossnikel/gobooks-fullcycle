# Go Intensivo

- Rob Pike e Robert Griesmer
2007 -> 2009 -> 2012 (1.0)

- Suporte robusto para testes
- Documentação automática
- Ferramentas integradas de profiling
- Fuzzing; SBOM;
- Proxies
- DB de vulnerabilidades

Simples comando vc ja ve pacotes, versões, vulnerabilidades, correções

1. Google
2. Netflix
3. Uber
4. Apple
5. MELI
6. Itau
7. Stone
8. Globo

- Projetada para ser simples
- Somente 25 palavras reservadas

- Porque é performado e eficiente?
  Compilada e gera um único binário


### Concorrência e Multithreading.
  - Próprio Runtime de memória
  - Goroutines (2kb) x 2MB

Da para gerar softwares otimizados que realizam várias tarefas ao mesmo tempo

- Muito performática e simples.
- TCP, UDP, WEBSOCKETS, HTTP, GRPC
- Ja possui testes nativos

Deploy aplicação em Go > Unico deploy com um arquivo binário.

GOOS=[operationalsystem] go build file.go

- Outras opções na hora de compilar

Libs padrões do Go
fmt. > Formata e exibe informações

Toda vez que usa um pacote usa o import.

Você pode declarar o tipo ou fazer atribuição automática


Type error no Go
nil::

Go é data driven -> Não é orientada a objetos, tem estrutura de dados que tem métodos mas tem pegada de estruturação de dados...

Projeto -> API :: CRUD de Livros
- Rest e CLI.
- Simular que alguém está lendo em x tempo
- Terminou todos livros -> Avisa que todos foram lidos
- Tem que ser lido simultâneamente
- DB etc etc.

interal -> dentro da aplicação que não compartilha com ninguém


GO MOD
Nome do módulo e versão.
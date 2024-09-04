# Ezequiel Lopes
# Sistema de Decodificação de Código Morse
Foi desenvolvido um sistema de decodificação de código Morse com três implementações, chamadas de `One`, `Two` e `Three`:

- **Versão One**: A versão mais simples, criada com o objetivo de dividir a mensagem em palavras e, em seguida, em letras para realizar a decodificação.
- **Versão Two**: Essa versão utiliza uma abordagem que percorre o código morse em uma única vez, identificando as letras e realizando a decodificação.
- **Versão Three**: Implementada apenas para utilizar goroutines no Go, demonstrou o pior desempenho, pois trata-se de uma aplicação sem grande espaço para ganho com concorrência, já que não há tempo de espera significativo no processo.

## Dependências

A aplicação depende do Golang para funcionar. Você pode instalar o Go seguindo o guia oficial:  
[https://go.dev/doc/install](https://go.dev/doc/install)

Para iniciar a aplicação, execute:

```bash
go mod tidy
```

Também é possível instalar a aplicação diretamente, o que permitirá rodar o comando morse-code:

```bash
go install github.com/ezequieljn/morse-code
```

Modos de Execução
O sistema pode ser iniciado de duas maneiras:

- `-mode`: Define se o sistema será executado via CLI ou HTTP.

- `-version`: Indica qual versão (`one` padrão, `two` ou `three`) será utilizada no processo.
- `-decode`: Caso o `-mode` seja CLI, este parâmetro define o código Morse a ser decodificado.
- `-space`: Define a forma de espaçamento entre as palavras (padrão ou /).
- `-port`: Caso utilize `HTTP`, é possível alterar a porta.


### Exemplo CLI
```bash
go run main.go -decode="- ... ..- .-. ..-   - ... ..- .-. ..-" -version="two"
# INFO Decoded value="TSURU TSURU"

go run main.go -decode=".... . -.-- / .--- ..- -.. ." -space=" / "
# INFO Decoded value="HEY JUDE"
```

### Exemplo HTTP:
```bash
go run main.go -mode="http"
curl --request POST --url http://localhost:8080 --data '{"code": "... --- ..."}'
# {"decoded": "SOS"}
```

### Testes e Benchmark
A aplicação possui diversos testes:

Testes de unidade para ambas as versões de decodificação:
```
go test ./...
```

```bash
# pkg/morse/decode_bm_test.go
BenchmarkCodeMorse/code_morse_version:_one-16         86401  13874 ns/op  2816 B/op  525 allocs/op
BenchmarkCodeMorse/code_morse_version:_two-16         85828  13901 ns/op  2816 B/op  525 allocs/op
BenchmarkCodeMorse/code_morse_version:_three-16       47344  25907 ns/op
```


### Conclusão
Todas as três versões são capazes de decodificar o código Morse, porém foram construídas de maneiras diferentes. Essa aplicação foi desenvolvida para demonstrar conceitos sobre arquitetura, boas práticas e funcionalidades do Golang, como:

- Princípios SOLID
- Design Patterns
- Goroutines
- Algoritmos
- Benchmarking
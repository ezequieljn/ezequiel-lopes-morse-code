Foi desenvolvido um sistema de decodificação de código Morse com três implementações, chamadas de "One", "Two" e "Three":

- **Versão One**: A versão mais simples, criada com o objetivo de dividir a mensagem em palavras e, em seguida, em letras para realizar a decodificação.
- **Versão Two**: Desenvolvida para ser comparada com a versão One, esta abordagem percorre a frase de uma só vez, identificando as letras e realizando a decodificação.
- **Versão Three**: Implementada apenas para utilizar goroutines no Go, demonstrou o pior desempenho, pois trata-se de uma aplicação sem grande espaço para ganho com concorrência, já que não há tempo de espera significativo no processo.

O sistema pode ser iniciado de duas maneiras:

- `-mode`: Define se o sistema será executado via `cli` ou `http`.
- `-version`: Indica qual versão (One, Two ou Three) será utilizada no processo.
- `-decode`: Caso o `-mode` seja `cli`, este parâmetro define o código Morse a ser decodificado.

A aplicação possui diversos testes:

- **Testes de unidade** para ambas as versões de decodificação:
  ```bash
  go test ./...



Também foram realizados testes de benchmark comparando o desempenho das três funções:
```bash
BenchmarkCodeMorse/code_morse_version:_one-16         	   86401	     13874 ns/op	    2816 B/op	     525 allocs/op
BenchmarkCodeMorse/code_morse_version:_two-16         	   85828	     13901 ns/op	    2816 B/op	     525 allocs/op
BenchmarkCodeMorse/code_morse_version:_three-16       	   47344	     25907 ns/op	
```
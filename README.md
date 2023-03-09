# convert

Uma analise dados para os json recebidos de terceiros.

A ideia básica é colocar uma analize em caso de erro, explicando oerro e uma análize aleatória, tipo 1% dos dados 
recebidos para procurar mudanças nos dados esperados

```shell
2023/03/09 10:24:29 O dado json está conforme esperado
2023/03/09 10:24:29 
2023/03/09 10:24:29 O dado json deveria ter o campo id como inteiro, mas, enviou float
2023/03/09 10:24:29 pismo transaction parser problem: there is an inconsistency between the received type (float64) and the expected type (int). value: `2.2`
2023/03/09 10:24:29 
2023/03/09 10:24:29 O dado json deveria ter o campo id como inteiro, mas, enviou string
2023/03/09 10:24:29 pismo transaction parser problem: there is an inconsistency between the received type (string) and the expected type (int). value: `2`
2023/03/09 10:24:29 
2023/03/09 10:24:29 O dado json chegou correto, mas, com um campo a mais, não previsto na documentação
2023/03/09 10:24:29 pismo transaction parser problem: there is a new key in the received json that was not predicted in the json handling code: key: name
2023/03/09 10:24:29 
```
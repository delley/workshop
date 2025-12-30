# Workshop Go: do Zero à API

## [DIA 01](../../tree/main/dia_01#dia-01)

1. [O que é Go](../../tree/main/dia_01#o-que-e-go)

    1. [Um pouco de história](../../tree/main/dia_01#um-pouco-de-historia)
    2. [Semântica](../../tree/main/dia_01#semantica)
    3. [Por que escolher Go?](../../tree/main/dia_01#por-que-escolher-go)
    4. [Hello World](../../tree/main/dia_01#hello-world)
    5. [Separando domínio de efeitos colaterais](../../tree/main/dia_01#separando-dominio-de-efeitos-colaterais)
2. [Tipos básicos](../../tree/main/dia_01#tipos-basicos)

    1. [String](../../tree/main/dia_01#string)
    2. [Números](../../tree/main/dia_01#numeros)
    3. [Inteiros](../../tree/main/dia_01#inteiros)
    4. [Alias úteis](../../tree/main/dia_01#alias-uteis)
    5. [Tipo especial](../../tree/main/dia_01#tipo-especial)
    6. [Booleanos](../../tree/main/dia_01#booleanos)
3. [Variáveis / Constantes / Ponteiros](../../tree/main/dia_01#variaveis-constantes-ponteiros)

    1. [Variáveis](../../tree/main/dia_01#variaveis)
    2. [Constantes](../../tree/main/dia_01#constantes)
    3. [Enumerações](../../tree/main/dia_01#enumeracoes)
    4. [Ponteiros](../../tree/main/dia_01#ponteiros)
    5. [A função `new()`](../../tree/main/dia_01#a-funcao-new)
4. [Tipos Compostos](../../tree/main/dia_01#tipos-compostos)

    1. [Array](../../tree/main/dia_01#array)
    2. [Slice](../../tree/main/dia_01#slice)
    3. [Map](../../tree/main/dia_01#map)

## [DIA 02](../../tree/main/dia_02#dia-02)

5. [Estruturas de controle](../../tree/main/dia_02#estruturas-de-controle)

    1. [if](../../tree/main/dia_02#if)
    2. [Switch](../../tree/main/dia_02#switch)
    3. [for](../../tree/main/dia_02#for)
    4. [for range](../../tree/main/dia_02#for-range)
6. [Struct](../../tree/main/dia_02#struct)

    1. [Struct Nomeada](../../tree/main/dia_02#struct-nomeada)
    2. [Struct anônima](../../tree/main/dia_02#struct-anonima)
7. [Funções](../../tree/main/dia_02#funcoes)

    1. [Como declarar uma função?](../../tree/main/dia_02#como-declarar-uma-funcao)
    2. [Declaração de função que recebe parâmetros](../../tree/main/dia_02#declaracao-de-funcao-que-recebe-parametros)
    3. [Funções anônimas](../../tree/main/dia_02#funcoes-anonimas)
    4. [Função com retorno nomeado](../../tree/main/dia_02#funcao-com-retorno-nomeado)
    5. [Funções variádicas](../../tree/main/dia_02#funcoes-variadicas)
    6. [Métodos](../../tree/main/dia_02#metodos)
8. [`defer`: execução adiada e gerenciamento seguro de recursos](../../tree/main/dia_02#defer-execucao-adiada-e-gerenciamento-seguro-de-recursos)

    1. [Funcionamento básico do `defer`](../../tree/main/dia_02#funcionamento-basico-do-defer)
    2. [`defer` e o ciclo de vida da função](../../tree/main/dia_02#defer-e-o-ciclo-de-vida-da-funcao)
    3. [Ordem de execução: LIFO (Last In, First Out)](../../tree/main/dia_02#ordem-de-execucao-lifo-last-in-first-out)
    4. [`defer` como ferramenta de gerenciamento de recursos](../../tree/main/dia_02#defer-como-ferramenta-de-gerenciamento-de-recursos)
    5. [Interação entre `defer` e valores de retorno](../../tree/main/dia_02#interacao-entre-defer-e-valores-de-retorno)
    6. [Armadilhas clássicas do `defer`](../../tree/main/dia_02#armadilhas-classicas-do-defer)
    7. [Boas práticas no uso de `defer`](../../tree/main/dia_02#boas-praticas-no-uso-de-defer)
    8. [Considerações finais sobre `defer`](../../tree/main/dia_02#consideracoes-finais-sobre-defer)
9. [Erros](../../tree/main/dia_02#erros)

10. [Interfaces](../../tree/main/dia_02#interfaces)


## [DIA 03](../../tree/main/dia_03#dia-03)

11. [Concorrência](../../tree/main/dia_03#concorrencia)

    1. [Goroutines](../../tree/main/dia_03#goroutines)
    2. [Canais (Channels)](../../tree/main/dia_03#canais-channels)
    3. [Select: Multiplexação com `select`](../../tree/main/dia_03#select-multiplexacao-com-select)
    4. [O pacote `sync`: Coordenação Explícita e Controle de Concorrência](../../tree/main/dia_03#o-pacote-sync-coordenacao-explicita-e-controle-de-concorrencia)
    5. [Canais versus primitivas do pacote `sync`](../../tree/main/dia_03#canais-versus-primitivas-do-pacote-sync)
    6. [Considerações finais sobre concorrência](../../tree/main/dia_03#consideracoes-finais-sobre-concorrencia)
12. [Pacotes](../../tree/main/dia_03#pacotes)

13. [Go Modules: Gerenciamento de Dependências](../../tree/main/dia_03#go-modules-gerenciamento-de-dependencias)

    1. [O que é o Go Modules?](../../tree/main/dia_03#o-que-e-o-go-modules)
    2. [GOPATH, um pouco de história](../../tree/main/dia_03#gopath-um-pouco-de-historia)
    3. [Configuração do projeto e ativação do Go Modules](../../tree/main/dia_03#configuracao-do-projeto-e-ativacao-do-go-modules)
    4. [Referências](../../tree/main/dia_03#referencias)
14. [Projeto](../../tree/main/dia_03#projeto)

    1. [API](../../tree/main/dia_03#api)
    2. [API Rest](../../tree/main/dia_03#api-rest)
    3. [API em Go com net/HTTP](../../tree/main/dia_03#api-em-go-com-nethttp)
15. [Referências](../../tree/main/dia_03#referencias)

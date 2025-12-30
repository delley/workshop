# Workshop Go: do Zero à API

## [DIA 01](../../tree/main/dia_01#dia-01)

1. [O que é Go](../../tree/main/dia_01#o-que-é-go)

    1. [Um pouco de história](../../tree/main/dia_01#um-pouco-de-história)
    2. [Semântica](../../tree/main/dia_01#semântica)
    3. [Por que escolher Go?](../../tree/main/dia_01#por-que-escolher-go)
    4. [Hello World](../../tree/main/dia_01#hello-world)
    5. [Separando domínio de efeitos colaterais](../../tree/main/dia_01#separando-domínio-de-efeitos-colaterais)
2. [Tipos básicos](../../tree/main/dia_01#tipos-básicos)

    1. [String](../../tree/main/dia_01#string)
    2. [Números](../../tree/main/dia_01#números)
    3. [Inteiros](../../tree/main/dia_01#inteiros)
    4. [Alias úteis](../../tree/main/dia_01#alias-úteis)
    5. [Tipo especial](../../tree/main/dia_01#tipo-especial)
    6. [Booleanos](../../tree/main/dia_01#booleanos)
3. [Variáveis / Constantes / Ponteiros](../../tree/main/dia_01#variáveis-constantes-ponteiros)

    1. [Variáveis](../../tree/main/dia_01#variáveis)
    2. [Constantes](../../tree/main/dia_01#constantes)
    3. [Enumerações](../../tree/main/dia_01#enumerações)
    4. [Ponteiros](../../tree/main/dia_01#ponteiros)
    5. [A função `new()`](../../tree/main/dia_01#a-função-new)
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
    2. [Struct anônima](../../tree/main/dia_02#struct-anônima)
7. [Funções](../../tree/main/dia_02#funções)

    1. [Como declarar uma função?](../../tree/main/dia_02#como-declarar-uma-função)
    2. [Declaração de função que recebe parâmetros](../../tree/main/dia_02#declaração-de-função-que-recebe-parâmetros)
    3. [Funções anônimas](../../tree/main/dia_02#funções-anônimas)
    4. [Função com retorno nomeado](../../tree/main/dia_02#função-com-retorno-nomeado)
    5. [Funções variádicas](../../tree/main/dia_02#funções-variádicas)
    6. [Métodos](../../tree/main/dia_02#métodos)
8. [`defer`: execução adiada e gerenciamento seguro de recursos](../../tree/main/dia_02#defer-execução-adiada-e-gerenciamento-seguro-de-recursos)

    1. [Funcionamento básico do `defer`](../../tree/main/dia_02#funcionamento-básico-do-defer)
    2. [`defer` e o ciclo de vida da função](../../tree/main/dia_02#defer-e-o-ciclo-de-vida-da-função)
    3. [Ordem de execução: LIFO (Last In, First Out)](../../tree/main/dia_02#ordem-de-execução-lifo-last-in-first-out)
    4. [`defer` como ferramenta de gerenciamento de recursos](../../tree/main/dia_02#defer-como-ferramenta-de-gerenciamento-de-recursos)
    5. [Interação entre `defer` e valores de retorno](../../tree/main/dia_02#interação-entre-defer-e-valores-de-retorno)
    6. [Armadilhas clássicas do `defer`](../../tree/main/dia_02#armadilhas-clássicas-do-defer)
    7. [Boas práticas no uso de `defer`](../../tree/main/dia_02#boas-práticas-no-uso-de-defer)
    8. [Considerações finais sobre `defer`](../../tree/main/dia_02#considerações-finais-sobre-defer)
9. [Erros](../../tree/main/dia_02#erros)

10. [Interfaces](../../tree/main/dia_02#interfaces)


## [DIA 03](../../tree/main/dia_03#dia-03)

11. [Concorrência](../../tree/main/dia_03#concorrência)

    1. [Goroutines](../../tree/main/dia_03#goroutines)
    2. [Canais (Channels)](../../tree/main/dia_03#canais-channels)
    3. [Select: Multiplexação com `select`](../../tree/main/dia_03#select-multiplexação-com-select)
    4. [O pacote `sync`: Coordenação Explícita e Controle de Concorrência](../../tree/main/dia_03#o-pacote-sync-coordenação-explícita-e-controle-de-concorrência)
    5. [Canais versus primitivas do pacote `sync`](../../tree/main/dia_03#canais-versus-primitivas-do-pacote-sync)
    6. [Considerações finais sobre concorrência](../../tree/main/dia_03#considerações-finais-sobre-concorrência)
12. [Pacotes](../../tree/main/dia_03#pacotes)

13. [Go Modules: Gerenciamento de Dependências](../../tree/main/dia_03#go-modules-gerenciamento-de-dependências)

    1. [O que é o Go Modules?](../../tree/main/dia_03#o-que-é-o-go-modules)
    2. [GOPATH, um pouco de história](../../tree/main/dia_03#gopath-um-pouco-de-história)
    3. [Configuração do projeto e ativação do Go Modules](../../tree/main/dia_03#configuração-do-projeto-e-ativação-do-go-modules)
14. [Projeto](../../tree/main/dia_03#projeto)

    1. [API](../../tree/main/dia_03#api)
    2. [API Rest](../../tree/main/dia_03#api-rest)
    3. [API em Go com net/HTTP](../../tree/main/dia_03#api-em-go-com-nethttp)
15. [Referências](../../tree/main/dia_03#referências)

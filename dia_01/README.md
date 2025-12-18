# Dia 01

## O que é Go

> _"Go é uma linguagem de programação de código aberto que facilita a criação de software **simples**, **confiável** e **eficiente**"_ 
>
> **[golang.org][1]**

### Um pouco de história

A linguagem surgiu na **Google**, em 2007, pelas mãos de:

- **Rob Pike**
- **Ken Thompson** (co-criador do Unix)
- **Robert Griesemer**

Foi anunciada publicamente em 2009 e teve sua versão 1.0 lançada em 2012. Desde então, vem sendo adotada em larga escala para infraestrutura, sistemas distribuídos e aplicações em nuvem.

### Semântica

Go apresenta certa familiaridade com C, sobretudo pela filosofia de obter o máximo de efeito com o mínimo de recursos e abstrações. Essa característica se reflete em uma sintaxe direta e em decisões de design voltadas à simplicidade e à eficiência.

No entanto, Go não deve ser entendida como uma versão modernizada de C. Trata-se de uma linguagem que incorpora boas ideias de diferentes paradigmas e linguagens, ao mesmo tempo em que elimina funcionalidades que historicamente introduzem complexidade excessiva ou comprometem a confiabilidade do código.

Go é uma linguagem compilada e estaticamente tipada. Embora ofereça suporte a ponteiros, não permite aritmética de ponteiros, reduzindo uma classe inteira de erros comuns em linguagens de baixo nível. Além disso, é uma linguagem moderna, projetada com primitivas de concorrência simples e eficientes, e conta com gerenciamento automático de memória por meio de garbage collection (GC).

### Por que escolher Go?

Go tem se consolidado como uma escolha estratégica para o desenvolvimento de sistemas modernos, especialmente no contexto de aplicações distribuídas, serviços de backend e ambientes em nuvem. Seu design prioriza simplicidade, desempenho, segurança e produtividade, oferecendo um equilíbrio raro entre eficiência de baixo nível e facilidade de uso. A seguir, destacam-se os principais motivos que levam equipes e organizações a adotarem Go.

#### Código aberto

Go é uma linguagem de código aberto, mantida pela comunidade e apoiada pelo Google. Isso garante transparência, liberdade de uso e evolução contínua, além de permitir que qualquer desenvolvedor contribua para o ecossistema, bibliotecas e ferramentas.

#### Fácil de aprender

A linguagem foi projetada com uma sintaxe enxuta e consistente, reduzindo a carga cognitiva do desenvolvedor. Mesmo profissionais vindos de linguagens como C, Java ou Python conseguem aprender Go rapidamente, graças à clareza do código e à ausência de construções excessivamente complexas.

#### Alto desempenho

Go combina desempenho elevado com alta produtividade. Por ser compilada diretamente para código de máquina e possuir um processo de compilação extremamente rápido, a linguagem é adequada tanto para desenvolvimento ágil quanto para aplicações que exigem eficiência e baixo consumo de recursos.

#### Modelo de concorrência simples

A concorrência é um dos pilares do Go. A linguagem introduz as _goroutines_, threads leves gerenciadas pelo runtime, e os _channels_, que facilitam a comunicação segura entre tarefas concorrentes. Esse modelo torna mais simples escrever código concorrente, escalável e menos propenso a erros.

#### Portabilidade e multiplataforma

Go é nativamente multiplataforma. O mesmo código-fonte pode ser compilado para diferentes sistemas operacionais e arquiteturas, como Linux, macOS e Windows, facilitando a distribuição de aplicações em ambientes heterogêneos.

#### Design orientado à nuvem

Go foi pensada desde o início para cenários de nuvem e microsserviços. As aplicações são distribuídas como um único binário estático, que inclui todas as dependências necessárias, simplificando o deploy, reduzindo problemas de compatibilidade e tornando o runtime mais previsível.

#### Segurança

Por ser uma linguagem estaticamente e fortemente tipada, Go força o desenvolvedor a ser explícito quanto aos tipos de dados utilizados. Isso permite que o compilador detecte uma ampla classe de erros em tempo de compilação, aumentando a confiabilidade e a segurança do código.

#### Coleta de lixo eficiente

Go possui gerenciamento automático de memória por meio de um _garbage collector_ moderno, projetado para minimizar pausas e impacto no desempenho. Esse mecanismo simplifica o desenvolvimento e contribui para aplicações mais estáveis e eficientes.

#### Biblioteca padrão robusta

A biblioteca padrão do Go é ampla e madura. Ela oferece suporte nativo para servidores HTTP, manipulação de entrada e saída, criptografia, concorrência, testes e muito mais, reduzindo a dependência de bibliotecas externas e acelerando o desenvolvimento de aplicações completas e confiáveis.


### Hello World

Assim como em outras linguagens... em Go temos também o clássico **Hello World**.

```go
// dia_01/exemplos/01-intro/ex1.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}
```

Para executar este código, você pode começar com o comando:

```shell
go run hello.go
```

### Separando domínio de efeitos colaterais

Mas... como podemos testar nosso código?

Primeiramente, vamos separar o "domínio" (regras de negócio) do restante do código (efeitos colaterais). A função `fmt.Println` é um efeito colateral (que está imprimindo um valor no `stdout` - saída padrão) e a `string` que estamos passando para ela é o nosso domínio.

```go
// dia_01/exemplos/01-intro/ex2.go
package main

import "fmt"

func Hello() string {
    return "Hello, World"
}

func main() {
    fmt.Println(Hello())
}
```

Criamos uma nova função: `Hello`, mas dessa vez adicionamos a palavra `string` na sua definição. Isso significa que essa função retornará uma `string`.

Para validar o comportamento da função `Hello`, criaremos um arquivo denominado `hello_test.go`, no qual será inserido o código a seguir:

```go
// dia_01/exemplos/01-intro/hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, World"

    if got != want {
        t.Errorf("Got '%s', want '%s'", got, want)
    }
}
```

Percebam que não é preciso usar vários _frameworks_ (ou bibliotecas) de testes. Tudo o que precisamos está pronto na linguagem e a sintaxe é a mesma para o resto dos códigos que você irá escrever.

Escrever um teste é como escrever uma função, com algumas regras:

- O código precisa estar em um arquivo que termine com **_test.go**.
- A função de teste precisa começar com a palavra **Test**.
- A função de teste recebe um único argumento, que é `t *testing.T`.

Por enquanto, isso é o bastante para saber que o nosso `t` do tipo `*testing.T` é a nossa porta de entrada para a ferramenta de testes.

Para executar este teste, você pode começar com o comando:
```shell
go test -v ex2.go hello_test.go
```

## Tipos básicos

A linguagem Go oferece diversos tipos básicos para representar e manipular dados — desde tipos que refletem diretamente os recursos do hardware até estruturas convenientes para modelar informações mais complexas.

### String

Uma **string** em Go é uma **sequência imutável de bytes**. Embora possam conter qualquer dado binário, as strings são normalmente utilizadas para representar texto legível por humanos.

Internamente, strings em Go são codificadas em **UTF-8**, o que significa que podem armazenar caracteres **Unicode** de forma eficiente.

::: note
**Nota técnica**

Para trabalhar com **caracteres Unicode**, usamos o tipo `rune`.
:::

### Números

Go oferece uma ampla variedade de tipos numéricos para diferentes tamanhos e precisões.

### Inteiros

Tipos com tamanho explícito:

| Tipo     | Descrição              |
|----------|------------------------|
| `uint8`  | conjunto de todos os inteiros sem sinal de 8 bits (de 0 a 255) |
| `uint16` | conjunto de todos os inteiros sem sinal de 16 bits (0 to 65535) |
| `uint32` | conjunto de todos os inteiros sem sinal de 32 bits (0 to 4294967295) |
| `uint64` | conjunto de todos os inteiros sem sinal de 32 bits (0 to 18446744073709551615)|
| `int8`   | conjunto de todos os inteiros de 8 bits com sinal (-128 a 127) |
| `int16`  | conjunto de todos os inteiros de 16 bits com sinal (-32768 to 32767) |
| `int32`  | conjunto de todos os inteiros de 32 bits com sinal (-2147483648 to 2147483647) |
| `int64`  | conjunto de todos os inteiros de 64 bits com sinal (-9223372036854775808 to 9223372036854775807) |

Tipos com tamanho dependente da arquitetura:

| Tipo     | Descrição              |
|----------|------------------------|
|`int` e `uint`| assumem o tamanho nativo do compilador (geralmente 32 ou 64 bits) |

### Alias úteis

| Tipo     | Descrição              |
|----------|------------------------|
| `byte` | alias para `uint8`, geralmente usado para representar dados binários|
| `rune` | alias para `int32`, especificamente projetado para armazenar valores inteiros que representam caracteres **Unicode** codificados em **UTF-8** |

### Tipo especial

| Tipo     | Descrição              |
|----------|------------------------|
| `uintptr` | Inteiro sem sinal que pode armazenar um ponteiro. Usado principalmente em programação de baixo nível, não garante portabilidade entre plataformas |

#### Ponto Flutuante

| Tipo     | Descrição              |
|----------|------------------------|
| `float32` | conjunto de todos os números de ponto flutuante de 32 bits do padrão IEEE-754 |
| `float64` | conjunto de todos os números de ponto flutuante de 64 bits do padrão IEEE 754 |

#### Complexos
| Tipo     | Descrição              |
|----------|------------------------|
| `complex64` e `complex128` | conjunto de todos os números complexos com partes real e imaginária armazenadas como valores float32 ou float64. Podem ser criados pela função `complex` |

### Booleanos

Em Go, valores booleanos são armazenados usando o tipo `bool`. Embora uma variável do tipo `bool` seja armazenada como um valor de 1 byte, ela não é, no entanto, um alias para um valor numérico. Go fornece dois literais pré-declarados:

| Tipo     | Descrição              |
|----------|------------------------|
| `bool`   | assume os literais `true` ou `false` |

## Variáveis / Constantes / Ponteiros

### Variáveis

Go é uma linguagem **fortemente tipada**, o que significa que **toda variável precisa ter um tipo associado** — explícito ou inferido. Cada variável em Go está vinculada a um **nome (identificador)**, um **tipo** e um **valor**.

A forma explícita para declarar uma variável em Go segue o seguinte formato:

```go
var <lista de identificadores> <tipo>
```

A palavra-chave `var` é usada para declarar um ou mais identificadores de variáveis, seguidos do seus respectivos tipos. O trecho de código a seguir mostra a declaração de diversas variáveis:

```go
// dia_01/exemplos/02-var/var01.go
...
var nome, desc string
var diametro int32
var massa float64
var ativo bool
var terreno []string
...
```

#### O valor zero

O trecho de código anterior mostra vários exemplos de variáveis sendo declaradas com uma variedade de tipos. À primeira vista, parece que essas variáveis não têm um valor atribuído. Na verdade, isso contradiz nossa afirmação anterior de que todas as variáveis em Go estão vinculadas a um tipo e um valor.

Durante a declaração de uma variável, se um valor não for fornecido, o compilador do Go vinculará automaticamente um valor padrão (ou um valor zero) à variável para a inicialização adequada da memória.

A tabela a seguir mostra os tipos do Go e seus valores zero padrão:

| Tipo | Valor zero |
|------|------------|
| string | "" (string vazia) |
| Numérico – Inteiro: byte, int, int8, int16, int32, int64, rune, uint, uint8, uint16, uint32, uint64, uintptr | 0 |
| Numérico – Ponto flutuante: float32, float64 | 0.0 |
| booleano| false |
| Array | Cada índice terá um valor zero correspondente ao tipo do array |
| Struct | Em uma estrutura vazia, cada membro terá seu respectivo valor zero |
| Outros tipos: interface, função, canal, slice, mapa e ponteiro | nil |

#### Declaração inicializada

Go também suporta a combinação de declaração de variável e inicialização como uma expressão usando o seguinte formato:

```go
var <lista de identificadores> <tipo> = <lista de valores ou expressões de inicialização>
```

O seguinte trecho de código mostra a combinação de declaração e inicialização:

```go
// dia_01/exemplos/02-var/var02.go
...
var nome, desc string = "Tatooine", "Planeta"
var diametro int32 = 10465
var massa float64 = 5.972E+24
var ativo bool = true
var terreno = []string{
    "Deserto",
}
...
```

#### Omitindo o tipo das variáveis

Em Go, também é possível omitir o tipo, conforme mostrado no seguinte formato de declaração:

```go
var <lista de identificadores> = <lista de valores ou expressões de inicialização>
```

O compilador do Go irá inferir o tipo da variável com base no valor ou na expressão de inicialização do lado direito do sinal de igual, conforme mostrado no trecho de código a seguir.

```go
// dia_01/exemplos/02-var/var03.go
...
var nome, desc = "Yavin IV", "Lua"
var diametro = 10200
var massa = 641693000000000.0
var ativo = true
var terreno = []string{
    "Selva",
    "Florestas Tropicais",
}
...
```

Quando o tipo da variável é omitido, as informações de tipo são deduzidas do valor atribuído ou do valor retornado de uma expressão.

A tabela a seguir mostra o tipo que é inferido dado um valor literal:

| Valor Literal | Tipo inferido |
|---------------|---------------|
| Texto com aspas duplas ou simples:<br>`"Lua Yavin IV"`<br><code>\`Sua superfície tem<br>seis continentes ocupando<br>67% do total.\`</code> | string |
| Inteiros:<br>`-51`<br>`0`<br>`1234` | int |
| Decimal:<br>`-0.12`<br>`1.0`<br>`1.3e5`<br>`5e-11` | float64 |
| Números complexos:<br>`-1.0i`<br>`2i`<br>`(0+2i)` | complex128 |
| Booleanos:<br>`true`<br>`false` | bool |
| Arrays:<br>`[2]int{-3, 51}`| O tipo do `array` definido pelo valor literal. Neste caso `[2]int` |
| Map:<br>`map[string]int{`<br>`"Tatooine": 10465,`<br>`"Alderaan": 12500,`<br>`"Yavin IV": 10200,`<br>`}` | O tipo do `map` definido pelo valor literal. Neste caso `map[string]int` |
| Slice:<br>`[]int{-3, 51, 134, 0}` | O tipo do `slice` definido pelo valor literal: `[]int` |
| Struct:<br>`struct{`<br>`nome string`<br>`diametro int`<br>`}{`<br>`"Tatooine", 10465,`<br>`}` | O tipo do `struct` definido conforme o valor literal. Neste caso: `struct{nome string; diametro int}` |
| Function:<br>`var sqr = func (v int) int {`<br>`    return v * v`<br>`}` | O tipo de `function` definido na definição literal da função. Neste caso, a variável `sqr` terá o tipo: `func (v int) int` |

#### Declaração curta de variável

Em Go é possível reduzir ainda mais a sintaxe da declaração de variáveis. Neste caso, usando o formato _short variable declaration_. Nesse formato, a declaração perde a palavra-chave `var` e a especificação de tipo e passa a usar o operador `:=` (dois-pontos-igual), conforme mostrado a seguir:

```go
<lista de identificadores> := <lista de valores ou expressões de inicialização>
```

O techo de código a seguir mostra como usá-la:

```go
// dia_01/exemplos/02-var/var04.go
...
func main() {
	nome := "Endor"
	desc := "Lua"
	diametro := 4900
	massa := 1.024e26
	ativo := true
	terreno := []string{
        "Florestas", 
        "Montanhas", 
        "Lagos",
    }
...
```

##### Restrições na declaração curta de variáveis

Existem algumas restrições quando usamos a declaração curta de variáveis e é muito importante estar ciente para evitar confusão:

- Em primeiro lugar, ela só pode ser usada dentro de um bloco de funções;
- o operador `:=` declara a variável e atribui os valores;
- `:=` não pode ser usado para atualizar uma variável declarada anteriormente;

#### Declaração de variável em bloco

A sintaxe do Go permite que a declaração de variáveis seja agrupada em blocos para maior legibilidade e organização do código. O trecho de código a seguir mostra a reescrita de um dos exemplos anteriores usando a declaração de variável em bloco:

```go
// dia_01/exemplos/02-var/var05.go
var (
	nome     string  = "Endor"
	desc     string  = "Lua"
	diametro int32   = 4900
	massa    float64 = 1.024e26
	ativo    bool    = true
	terreno          = []string{
		"Florestas",
		"Montanhas",
		"Lagos",
	}
)
```

### Constantes

Uma constante é um valor com uma representação literal de uma string, um caractere, um booleano ou números. O valor para uma constante é estático e não pode ser alterado após a atribuição inicial.

#### Constantes tipadas

Usamos a palavra chave `const` para indicar a declaração de uma constante. Diferente da declaração de uma variável, a declaração deve sempre incluir o valor literal a ser vinculado ao identificador, conforme mostrado a seguir:

```go
const <lista de identificadores> tipo = <lista de valores ou expressões de inicialização>
```

O seguinte trecho de código mostra algumas constantes tipadas sendo declaradas:

```go
// dia_01/exemplos/03-constants/const01.go
...
const a1, a2 string = "Workshop", "Go"
const b rune = 'G'
const c bool = false
const d int32 = 2020
const e float32 = 2.020
const f float64 = math.Pi * 2.0e+3
const g complex64 = 20.0i
const h time.Duration = 20 * time.Second
...
```

Note que cada constante declarada recebe explicitamente um tipo. Isso implica que a constantes só podem ser usada em contextos compatíveis com seus tipos. No entanto, isso funciona de maneira diferente quando o tipo é omitido.

#### Constantes não tipadas

Constantes são ainda mais interessantes quando não são tipada. Uma constante sem tipo é declarada da seguinte maneira:

```go
const <lista de identificadores> = <lista de valores ou expressões de inicialização>
```

Neste formato, a especificação de tipo é omitida na declaração. Logo, uma constante é meramente um bloco de bytes na memória sem qualquer tipo de restrição de precisão imposta. A seguir, algumas declarações de constantes não tipificadas:

```go
// dia_01/exemplos/03-constants/const02.go
...
const i = "G é" + " para Go"
const j = 'G'
const k1, k2 = true, !k1
const l = 111*100000 + 20
const m1 = math.Pi / 3.141592
const m2 = 1.41421356237309504880168872420969807856967187537698078569671875376
const m3 = m2 * m2
const m4 = m3 * 20.0e+400
const n = -5.0i * 20
const o = time.Millisecond * 20
...
```

A constante `m4` recebe um valor muito grande (`m3 * 20.0e+400`) que é armazenado na memória sem qualquer perda de precisão. Isso pode ser útil em aplicações onde realizar cálculos com um alto nível de precisão é extremamente importante.

#### Atribuindo constantes não tipadas

Mesmo Go sendo uma linguagem fortemente tipada, é possível atribuir uma constante não tipada a diferentes tipos de precisão diferentes, embora compatíveis, sem qualquer reclamação do compilador, conforme mostrada a seguir:

```go
// dia_01/exemplos/03-constants/const03.go
...
const m2 = 1.41421356237309504880168872420969807856967187537698078569671875376
var u1 float32 = m2
var u2 float64 = m2
u3 := m2
...
```

O exemplo anterior mostra a constante não tipada `m2` sendo atribuída a duas variáveis de ponto flutuante com diferentes precisões, `u1` e `u2`, e a uma variável sem tipo, `u3`. Isso é possível porque a constante `m2` é armazenada como um valor não tipado e, portanto, pode ser atribuída a qualquer variável compatível com sua representação (um ponto flutuante).

Como `u3` não tem um tipo específico, ele será inferido a partir do valor da constante, e como `m2` representa um valor decimal, o compilador irá inferir seu tipo padrão, um `float64`.

A declaração de constantes também podem ser organizadas em blocos, aumentando a legibilidade do código, conforme a seguir:

```go
// dia_01/exemplos/03-constants/const04.go
...
const (
	a1, a2 string        = "Workshop", "Go"
	b      rune          = 'G'
	c      bool          = false
	d      int32         = 2020
	e      float32       = 2.020
	f      float64       = math.Pi * 20.0e+3
	g      complex64     = 20.0i
	h      time.Duration = 20 * time.Second
)
...
```

### Enumerações

Um interessante uso para constantes é na criação de enumerações. Usando a declaração de blocos, é facilmente possível criar valore inteiros que aumentam numericamente. Para isso, basta atribuir o valor constante pré-declarado `iota` a um identificador de constante na declaração de bloco, conforme mostrado no exemplo a seguir:

```go
// dia_01/exemplos/04-enum/enum01.go
...
const (
	estrelaHiperGigante = iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
	estrelaAna
	estrelaSubAna
	estrelaAnaBranca
	estrelaAnaVermelha
	estrelaAnaMarrom
)
...
```

Nessa situação, o compilador fará o seguinte:

- Declarar cada membro no bloco como um valor constante inteiro não tipado;
- Inicializar a `iota` com o valor zero;
- Atribuir a `iota`, ou zero, ao primeiro membro (`EstrelaHiperGigante`);
- Cada constante subsequente recebe um `int` aumentado em um.

Assim, as constantes da lista receberão os valores de zero até nove.

É importante ressaltar que, sempre que `const` aparecer em um bloco de declaração, o contador é redefinido para zero. No trecho de código seguinte, cada conjunto de constantes é enumerado de zero a quatro:

```go
// dia_01/exemplos/04-enum/enum02.go
...
const (
	estrelaHiperGigante = iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
)
const (
	estrelaAna = iota
	estrelaSubAna
	estrelaAnaBranca
	estrelaAnaVermelha
	estrelaAnaMarrom
)
...
```

#### Substituindo o tipo padrão de uma enumeração

Por padrão, uma constante enumerada é declarada como um tipo inteiro não tipado. Porém, podemos substituir o tipo padrão provendo explicitamente um tipo numérico, como mostrado a seguir:

```go
// dia_01/exemplos/04-enum/enum03.go
...
const (
	estrelaAna byte = iota
	estrelaSubAna
	estrelaAnaBranca
	estrelaAnaVermelha
	estrelaAnaMarrom
)
...
```
É possível especificar qualquer tipo numérico que pode representar um inteiro ou um ponto flutuante. No exemplo anterior, cada constante será declarada como um tipo `byte`.

#### Usando `iota` em expressões

Quando a `iota` aparece em uma expressão, o compilador irá aplicar a expressão para cada valor sucessivo. O exemplo a seguir atribui números pares aos membros do bloco de declaração:

```go
// dia_01/exemplos/04-enum/enum04.go
...
const (
	estrelaHiperGigante = 2.0 * iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
)
...
```

#### Ignorando valores em enumerações

É possível ignorar certos valores em uma enumeração simplesmente atribuindo a `iota` a um identificador em branco (`_`). No trecho de código a seguir, o valor `0` é ignorado:

```go
// dia_01/exemplos/04-enum/enum05.go
...
const (
	_                   = iota
	estrelaHiperGigante = 1 << iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
)
...
```

### Ponteiros

Go possibilita o uso de ponteiros. Um *ponteiro* é o *endereço* de memória de um valor.

Um ponteiro em Go é definido por operador `*`(asterisco). O trecho de código a seguir mostra um exemplo da utilização de ponteiros:

```go
var p *int
```
Um ponteiro é definido de acordo com seu tipo de dado.

No código anterior a variável `p` é um ponteiro para um valor do tipo `int`.

Também é possível obter o endereço do valor de uma variável, para isso, utilizamos o operador `&` (e comercial).

```go
eraOuroSith := 5000
p := &eraOuroSith
```

Já o valor referenciado ao ponteiro pode ser acessado usando o operador `*`.

```go
eraOuroSith := 5000
p := &eraOuroSith
fmt.Println(*p) // imprime 5000
```

Um exemplo mais completo:

```go
// dia_01/exemplos/05-pointer/pont01.go
...
var p *int
eraOuroSith, epIV := 42, 37
// ponteiro para eraOuroSith
p = &eraOuroSith
// valor de eraOuroSith por meio do ponteiro
fmt.Printf("Era de Ouro dos Sith - %d anos antes do Ep.IV (%#x)\n", *p, p)
// atualiza o valor de eraOuroSith por meio do ponteiro
*p = 5000
// o novo valor de eraOuroSith
fmt.Printf("Era de Ouro dos Sith - %d anos antes do Ep.IV | Atualizado (%#x)\n", *p, p)
// ponteiro para epIV
p = &epIV
// divide epIV por meio do ponteiro
*p = *p / 38
// o novo valor de epIV
fmt.Printf("Star Wars: Ep.IV é o Marco %d (%#x)\n", epIV, p)
...
```

::: important
**Nota técnica**

Go não permite **aritmética de ponteiros**.
:::

### A função `new()`

Outra forma de criar variáveis em Go, é usando a função `new()`.

A expressão `new(T)` cria uma variável *sem nome* do tipo `T`, inicializa ela com seu valor zero e devolve seu endereço de memória.

```go
// dia_01/exemplos/06-new/new01.go
...
	// epIV, do tipo *int, aponta para uma variável sem nome
	epIV := new(int)
	// eraOuroSith, do tipo *int, também aponta para uma variável sem nome
	eraOuroSith := new(int)
	// "0" zero
	fmt.Println(*eraOuroSith)
	// novo valor para o int sem nome
	*eraOuroSith = *epIV - 5000
	// "-5000"
	fmt.Println(*eraOuroSith)
...
```

O uso da função `new()` é relativamente raro.

## Tipos Compostos

Tipos compostos em Go são tipos criados pela combinação de tipos básicos e tipos compostos.

### Array

*Array* é uma sequência de elementos do mesmo tipo de dados. Um *array* tem um tamanho fixo, o qual é definido em sua declaração, e não pode ser mais alterado.

A declaração de um *array* segue o seguinte formato:

```go
[<tamanho>]<tipo do elemento>
```

Exemplo:

```go
var linhaTempo [10]int
```

*Arrays* também podem ser multidimensionais:

```go
var mult [3][3]int
```

Iniciando um *array* com valores:

```golang
var linhaTempo = [3]int{0, 5, 19}
```

Você pode usar `...`(reticências) na definição de capacidade e deixar o compilador definir a capacidade com base na quantidade de elementos na declaração.

```go
// Declaração simplificada
linhaTempo := [...]int{0, 5, 19}
```

Neste caso, o tamanho do *array* será 3.

O próximo exemplo mostra como atribuir valores a um *array* já definido:

```go
// dia_01/exemplos/07-arrays/arr01.go 
...
var linhaTempo [3]int
linhaTempo[0] = 0
linhaTempo[1] = 5
linhaTempo[2] = 19
...
```

#### Tamanho de um *array*:

O tamanho de um *array* pode ser obtido por meio da função nativa `len()`.

```go
// dia_01/exemplos/07-arrays/arr02.go
...
// Declaração simplificada
linhaTempo := [...]int{0, 5, 19}
// imprime 3
fmt.Println(len(linhaTempo))
...
```

### Slice

*Slice* é *wrap* flexível e robusto que abstrai um *array*. Em resumo, um *slice* não detém nenhum dado nele. Ele apenas referencia *arrays* existentes.

A declaração de um *slice* é parecida com a de um *array*, mas sem a capacidade definida.

```go
// dia_01/exemplos/08-slice/slice01.go
...
// declaracao com var
var s1 []int
fmt.Println("Slice 1:", s1)
// declaração curta
s2 := []int{}
fmt.Println("Slice 2:", s2)
// tamanho de um slice
fmt.Println("Tamanho do slice 1:", len(s1))
fmt.Println("Tamanho do slice 2:", len(s2))
...
```

O código anterior criou um *slice* sem capacidade inicial e sem nenhum elemento.

Também é possível criar um um *slice* a partir de um array:

```go
// dia_01/exemplos/08-slice/slice02.go
...
1	// Naves do jogo "Star Wars: Battlefront"
2	naves := [...]string{
3		1: "X-Wing",
4		2: "A-Wing",
5		3: "Millenium Falcon",
6		4: "TIE Fighter",
7		5: "TIE Interceptor",
8		6: "Imperial Shuttle",
9		7: "Slave I",
10	}
11	// cria um slice de naves[1] até naves[3]
12	rebeldes := naves[1:4]
13	fmt.Println(rebeldes)
...
```

A sintaxe `s[i:j]` cria um *slice* a partir do *array* `naves` iniciando do índice `i` até o índice `j - 1`. Então, na **linha 12** do código, `naves[1:4]` cria uma representação do *array* `naves` iniciando do índice 1 até o 3. Sendo assim, o slice `rebeldes` tem os valores `["X-Wing" "A-Wing" "Millenium Falcon"]`.

Um *slice* pode ser criado usando a função `make()`, uma função nativa que cria um *array* e retorna um *slice* referenciando o mesmo.

A sintaxe da função é a seguinte: `func make([]T, len, cap) []T`.

Neste caso, é passando como parâmetro o **tipo (T)**, o **tamanho (len)** e a **capacidade (cap)**. A capacidade é opcional, e caso não seja informada, seu valor *padrão* será o **tamanho (len)**, que é um campo obrigatório.

```go
// dia_01/exemplos/08-slice/slice03.go
...
s := make([]int, 5, 5)
fmt.Println(s)
...
```
 
#### Adicionando elementos a um slice

Como sabemos, *arrays* são limitados em seu tamanho e não podem ser aumentados. Já *Slices*, tem seu tamanho dinâmico e podem receber novos elementos em tempo de execução por meio da função nativa `append`.

A definição da função `append` é a seguinte: `func append(s []T, x ...T) []T`.

A sintaxe, `x ...T` significa que a função aceita um número variável de elementos no parâmetro `x`, desde que respeitem o tipo do *slice*.

```go
// dia_01/exemplos/08-slice/slice04.go
...
// Naves do jogo "Star Wars: Battlefront"
rebeldes := [...]string{"'X-Wing'", "'A-Wing'", "'Millenium Falcon'"}
imperiais := [...]string{"'TIE Fighter'", "'TIE Interceptor'", "'Imperial Shuttle'", "'Slave I'"}

naves := make([]string, 0, 0)
fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
naves = append(naves, "''")
fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
naves = append(naves, rebeldes[:]...)
fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
naves = append(naves, imperiais[:]...)
fmt.Printf("Cap: %d - %v\n", cap(naves), naves)
...
```

Uma questão que pode ter *ficado no ar*: Se um slice é um *wrap* de um *array*, como ela tem esta flexibilidade?

Bem, o que acontece por *debaixo dos panos* quando um novo elemento é adicionado a um *slice* é o seguinte:

1. Um novo *array* é criado
2. Os elementos do *array* atual são copiados
3. O elemento ou elementos **adicionados** ao *slice* são incluido no *array*
4. É retornado um *slice*, que é uma referência para o novo *array*

### Map

Um *Map* é uma estrutura de dados que mantém uma coleção de pares chave/valor.
Também conhecido como *hash table* (tabela de dispersão ou tabela hash).

A declaração de um *map* segue o seguinte formato:

```go
map[k]v
```
Onde `k` é o tipo da chave e `v` o tipo dos valores.

Exemplos de uso de map:

```golang
// dia_01/exemplos/09-map/map01.go
...
naves := make(map[string]string)

naves["YT-1300"] = "Millennium Falcon"
naves["T-65"] = "X-Wing"
naves["RZ-1"] = "A-Wing"
naves["999"] = "Tunder Tanque"

fmt.Println("Quantidade de naves:", len(naves))
fmt.Println(naves)
fmt.Printf("Nave do Han Solo: %s\n", naves["YT-1300"])

fmt.Println("999 não é uma nave. Removendo...")
delete(naves, "999")

fmt.Println("Quantidade de naves atualizada:", len(naves))
fmt.Println(naves)
...
```


[1]:https://golang.org/
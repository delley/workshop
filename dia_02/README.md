# Dia 02

## Estruturas de controle

Estruturas de controle são ultilizadas para alterar a forma como o nosso código é executado. Podemos, por exemplo, fazer com que uma parte do nosso código seja repitido várias vezes, ou que seja executado caso uma condição seja satisfeita.

### if

O `if` é uma instrução que avalia uma condição booleana. Para entender melhor como ele funciona, vamos analisar o seguinte problema:

> Em uma soma de dois números, onde `a` = 2 e `b` = 4, avalie o valor de `c`. Se `c` for igual a 6, imprima na tela *"Sua soma está correta!"*, caso contrário, imprima *"Sua soma está errada!"*.

```go
package main 

func main() {
    var a, b = 2, 4
    c := (a + b)
    if c == 6 {
        fmt.Println("Sua soma está correta.")
        return
    }
    //uma forma d fazer se não
    fmt.Println("Sua soma está errada.")
}
```

#### Outro Exemplo:

```go
package main 

func main() {
    if 2%2 == 0 {
        fmt.Println("É par.")
    } else {
        fmt.Println("É impar.")
    }

    if num := 2 num < 0 {
        fmt.Println(num, "É negativo.")
    } else if num < 10 {
        fmt.Println(num, "Tem um dígito.")
    } else {
        fmt.Println(num, "Tem vários dígitos.")
    }
}
```

### Switch

A instrução ***`switch`*** é uma maneira mais fácil de evitar longas instruções ***`if-else`***. Com ela é possível realiza ações diferentes com base nos possíveis valores de uma expressão.

#### Exemplo 1

```go
package main 

func main() {
    i := 2
    switch i {
    case 1:
        fmt.Println("Valor de ", i, " por extenso é: um")
    case 2:
        fmt.Println("Valor de ", i, " por extenso é: dois")
    case 3:
        fmt.Println("Valor de ", i, " por extenso é: três")
    }
}
```

O ***`switch`*** pode testar valores de qualquer tipo, além de podermos usar vírgula para separar várias expreções em uma mesma condição ***`case`***.

#### Exemplo 2

```go
package main

func main() {
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("É fim de semana.")
    default:
        fmt.Println("É dia de semana.")
    }
}
```

O ***`switch`*** sem uma expressão é uma maneira alternativa para expressar uma lógica ***`if-else`***.

#### Exemplo 3

```go
package main

func main() {
    j := 3
    switch {
    case 1 == j:
        fmt.Println("Valor por extenso é: um")
    case 2 == j:
        fmt.Println("Valor por extenso é: dois")
    default:
        fmt.Println("Valor não encontrado.")
    }
}
```

### for

Em outras linguagens de programação temos várias formas de fazer laços de repetição, porém, em **Go** só temos uma forma, e é usando a palavra reservada ***`for`***.

#### Exemplo 1

A forma tradicional, que já conhecemos, e que no exemplo vai imprimir números de 1 a 10.

```go
package main

func main(){
    for i := 1; i <=10; i++ {
        fmt.Println("O número é: ", i)
    }
}
```

#### Exemplo 2

```go
package main

func main(){
    i := 5
    for i <= 5 {
        fmt.Println("O número é: ", i)
        i = i + 1
    }
}
```

#### Exemplo 3 loop infinito

```go
package main

func main(){
    for {
        fmt.Println("Olá sou o infinito")
        break
    }
}
```

### for range

Já vimos as outras formas de usar o ***`for`***, agora falta o ***`range`***. Essa expressão espera receber uma lista (`array` ou `slice`).

```go
func exemploFor4() {
	listaDeCompras := []string{"arroz", "feijão", "melancia", "banana", "maçã", "ovo", "cenoura"}
	for k, p := range listaDeCompras {
		retornaNomeFruta(k, p)
	}
}

func retornaNomeFruta(key int, str string) {
	switch str {
	case "melancia", "banana", "maçã":
		fmt.Println("Na posição", key, "temos a fruta:", str)
	default:
		return
	}
}
```

## Struct

***`Stuct`*** é um tipo de dado agregado que agrupa zero ou mais valores nomeados de tipo quaisquer como uma única entidade. Cada valor é chamado de ***campo***.

### Struct Nomeada

Uma `struct` nomeada recebe um nome em sua declaração. 
Para exemplificar, criaremos uma `strutc` para representar um cadastro de funcionário em uma empresa. Seus campos pode ser acessados través da expressão `variaval.Name`, exemplo:

```go
package main

type Employee struct {
    ID      int
    Name    string
    Age     *time.Time
    Salary  float64
    Company string
}

func main() {
    cl := Employee{}
    //forma de acesso
    cl.ID = 1
    cl.Name = "Diego dos Santos"
    cl.Age = nil
    cl.Salary = 100.55
    cl.Company = "Fliper"
    fmt.Println("o nome é:", cl.Name, " trabalha na empresa: ", cl.Company)
    //oura forma de popular structs
    cl1 := Employee{
        ID:      1,
        Name:    "Francisco Oliveira",
        Age:     nil,
        Salary:  2000.50,
        Company: "Iron Mountain",
    }
    fmt.Println("o nome é:", cl1.Name, " trabalha na empresa: ", cl1.Company)

}
```

### Struct anônima

Uma ***`struct`*** anônima é tipo sem um nome como referência. Sua declaração é semelhante a uma ***declaração rapída de variável***.

> Só devemos usar uma `struct` anônima quando não há necessida de criar um objeto para o dado que será transportado por ela.

```go
package main

func main() {
    inferData("Diego", "Santos")
    inferData("Francisco", "Oliveira")
}

func inferData(fN, lN string) {
    name1 := struct{FirstName, LastName string}{FirstName: fN, LastName: lN}
    fmt.Println("O nome é:", name1.FirstName, name1.LastName)
}
```

## Funções

Funções são pequenas unidades de códigos que podem abistrair ações, retornar e/ou receber valores.

### Como declarar uma função?

Declarar uma função é algo bem simples, utilizando a palavra reservada ***`func`*** seguida do identificador.

```go
package main

func nomeDaFuncao(){}

```

Essa é a declaração mais simples de função que temos. No exemplo acima criamos uma função que não recebe nenhum parametro e não retorna nada, o nome dela poderia ser `fazNada`.

Uma função em Go também é um tipo e pode ser declarada da segunte forma:

```go
package main

type myFunc = func(l, b int) int

func main() {
	soma(func(l, b int) int {
		return l + b
	})
}

func soma(fn myFunc) {
	res := fn(1, 3)
	fmt.Println(res)
}

```

### Declaração de função que recebe parâmetros

Podemos declarar uma função que recebe dois números e faz uma multiplicação.

```go
package main

func main() {
	fmt.Println("Resultado é:", multiplica(3, 8))
}

func multiplica(a, b int) int {
	return (a * b)
}
```

Veja que na declaração da `func multiplica(a, b int)` os parametros foram passados um seguido do outro, isso por que eles são do mesmo **tipo** (`int`). Caso fossem de tipos diferentes seria nessário declarar cada tipo separadamente, exemplo `func minhaFunc(str string, i int)`.

### Funções anônimas

Go também suporta declaração de funções anônimas. Funções anônimas são úteis quando você deseja definir uma função em linha sem ter que nomeá-la.

```go
package main

func main() {
    fn := exemploAnonimo()
    fmt.Println("Resultado é:", fn)
    fmt.Println("Resultado é:", fn)
    fmt.Println("Resultado é:", fn)
}

func exemploAnonimo() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
```

### Função com retorno nomeado

Podemos criar uma função e nomear o retorno da mesma. veja o exemplo:

```go
package main

func main() {
    fn := exemploNomeado()
    fmt.Println("Nome é:", exemploNomeado("Marcela"))
    fmt.Println("Nome é:", exemploNomeado("Diego"))
    fmt.Println("Nome é:", exemploNomeado("Francisco"))
}

func exemploNomeado(str string) (nome string) {
    nome = str
    return 
}
```

### Funções variádicas

Função variádica é uma função que pode receber qualquer número de argumentos à direita e de um mesmo tipo. Um bom exemplo de função variádica é a função `fmt.Println`. A função pode ser chamada de forma usual, com argumentos individuais ou uma lista (`Slice`).

```go
package main

func main() {
    fmt.Println("Resultado é:", exemploVariadico(1,2))
    fmt.Println("Resultado é:", exemploVariadico(2,3))
    fmt.Println("Resultado é:", exemploVariadico(3,4))
}

func exemploVariadico(numeros ...int) (total int) {
    total = 0

    for _, n := range numeros {
        total += n
    }
	return 
}
```

### Métodos

Métodos em Go são uma variação da declaração de função. No método, um parâmetro extra aparece antes do nome da função e é chamado de receptor (*receiver*).

Métodos podem ser definidos para qualquer tipo de receptor, até mesmo ponteiros, exemplo:


```go
package main

type area struct {
    Largura int
    Altura  int
}

func (r *area) CalculaArea() int {
    res := r.Largura * r.Altura
    return res
}

func (r area) CalculaPerimetro() int {
    res := 2*r.Largura * 2*r.Altura
    return res
}


func main() {
    a := area{Largura: 10, Altura: 5}
    resultArea := a.CalculaArea()
    fmt.Println("area: ", resultArea)
    perim := &a //repassando os valores
    resultPerim := perim.CalculaPerimetro()
    fmt.Println("perim: ", resultPerim)
}
```

## `defer`: execução adiada e gerenciamento seguro de recursos

O `defer` é um dos mecanismos mais característicos da linguagem Go. Ele permite adiar a execução de uma função até o momento em que a função envolvente retorna, independentemente de como esse retorno ocorre. Essa característica torna o `defer` uma ferramenta central para **gerenciamento de recursos**, **segurança do fluxo de execução** e **legibilidade do código**.

Em Go, o `defer` não é apenas um atalho sintático; ele expressa uma intenção clara: _“este código deve ser executado quando eu sair deste escopo”_. Essa semântica simples tem implicações profundas na forma como escrevemos código robusto, especialmente em funções com múltiplos pontos de retorno.

### Funcionamento básico do `defer`

A instrução `defer` agenda a execução de uma chamada de função para ocorrer **após** o retorno da função atual. A avaliação dos argumentos ocorre **imediatamente**, no momento em que o `defer` é declarado, enquanto a execução da função é postergada.

```go
// dia_02/exemplos/08-defer/ex1.go
...
func exemplo() {
	defer fmt.Println("mundo!")
	fmt.Print("Olá ")
}
...
```

A saída será:

```bash
Olá mundo!
```

Mesmo que a função possua múltiplos `return`, panics ou fluxos condicionais complexos, as funções deferidas serão executadas de forma garantida.

### `defer` e o ciclo de vida da função

As chamadas deferidas são associadas à **ativação da função na pilha de chamadas**. Quando a função começa a retornar, o runtime executa os `defer` registrados antes de desalocar o frame da pilha.

Isso implica que:

* `defer` **não está ligado a blocos**, mas à função inteira;
* não existe **“defer de escopo local”** em Go;
* o momento exato da execução ocorre após a avaliação do `return`, mas antes do controle voltar ao chamador.

```go
// dia_02/exemplos/08-defer/ex2.go
...
func soma(a, b int) int {
	defer fmt.Println("fim da função soma")
	return a + b
}
...
```

### Ordem de execução: LIFO (Last In, First Out)

Uma característica essencial do `defer` é que múltiplas chamadas deferidas são executadas em **ordem inversa** àquela em que foram declaradas, seguindo o modelo de uma pilha (LIFO).

```go
// dia_02/exemplos/08-defer/ex3.go
...
func exemplo() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
...
```

Saída:

```bash
3
2
1
```

Esse comportamento é intencional e extremamente útil para padrões como:

* aquisição e liberação de múltiplos recursos,
* locks encadeados,
* composição segura de operações.

### `defer` como ferramenta de gerenciamento de recursos

O uso mais comum de `defer` está associado à liberação de recursos externos, como arquivos, conexões e locks.

```go
// dia_02/exemplos/08-defer/ex4.go
...
file, err := os.Open("dados.txt")
if err != nil {
    return err
}
defer func() {
    file.Close()
    fmt.Println("Arquivo fechado com sucesso!")
}()
...
```

Esse padrão garante que:

* o recurso será liberado exatamente uma vez;
* o código permanece correto mesmo com retornos antecipados;
* a intenção do desenvolvedor fica explícita.

O mesmo princípio se aplica a:

* `defer rows.Close()`
* `defer conn.Close()`
* `defer mu.Unlock()`
* `defer cancel()`

### Interação entre `defer` e valores de retorno

Em Go, funções podem ter **valores de retorno nomeados**, e o `defer` pode interagir diretamente com eles.

```go
// dia_02/exemplos/08-defer/ex5.go
...
func contador() (n int) {
    defer func() {
        n++
    }()
    return 10
}
...
```

Nesse caso, o valor retornado será `11`, pois:

1. `n` recebe o valor `10`;
2. o `defer` é executado;
3. a função retorna o valor final de `n`.

Esse comportamento é poderoso, mas deve ser usado com cautela, pois pode reduzir a clareza do código se aplicado indiscriminadamente.

### Armadilhas clássicas do `defer`

Apesar de sua simplicidade aparente, o `defer` possui algumas armadilhas conhecidas que merecem atenção.

#### 1. `defer` dentro de loops

Uma das armadilhas mais comuns ocorre ao usar `defer` dentro de loops longos.

```go
...
for i := 0; i < 1000; i++ {
    f, _ := os.Open(fmt.Sprintf("file%d.txt", i))
    defer f.Close()
}
...
```

Nesse exemplo, **todos os arquivos permanecem abertos até o final da função**, o que pode causar exaustão de recursos.

A abordagem correta é limitar o escopo da função ou extrair a lógica para uma função auxiliar:

```go
...
for i := 0; i < 1000; i++ {
    func() {
        f, _ := os.Open(fmt.Sprintf("file%d.txt", i))
        defer f.Close()
        // uso do arquivo
    }()
}
...
```

#### 2. Avaliação imediata dos argumentos

Os argumentos de uma função deferida são avaliados no momento da declaração, não no momento da execução.

```go
// dia_02/exemplos/08-defer/ex6.go
...
for i := 0; i < 3; i++ {
    defer fmt.Println(i)
}
...
```

Saída:

```bash
2
1
0
```

Isso ocorre porque o valor de `i` é capturado a cada iteração, no momento do `defer`.

Esse comportamento é correto e previsível, mas frequentemente causa confusão em closures mais complexas.

#### 3. Custo de performance do `defer`

O `defer` possui um custo maior do que uma chamada direta de função. Embora esse custo tenha sido significativamente reduzido nas versões modernas do Go, ele **não é zero**.

Em código crítico de altíssima performance (loops internos, hot paths), pode ser preferível uma liberação explícita:

```go
mu.Lock()
// código crítico
mu.Unlock()
```

Ao invés de:

```go
mu.Lock() 
defer mu.Unlock()
```

A regra prática é clara:

* priorize **clareza e segurança**;
* otimize apenas quando houver evidência mensurável de impacto.

#### 4. `defer` não substitui controle explícito de fluxo

Embora poderoso, o `defer` não deve ser usado para ocultar lógica complexa ou efeitos colaterais não óbvios.
Evite:

* modificar estado global de forma implícita;
* alterar valores de retorno sem clareza;
* criar dependências implícitas difíceis de rastrear.
    
O uso idiomático do `defer` é **simples, previsível e local**.

### Boas práticas no uso de `defer`

* Declare o `defer` **imediatamente após** a aquisição do recurso.
* Use `defer` para liberar recursos, não para lógica de negócio.
* Prefira clareza a micro-otimizações.
* Evite `defer` em loops extensos sem controle de escopo.
* Trate `defer` como parte do contrato de segurança da função.

### Considerações finais sobre `defer`

O `defer` é um dos pilares do estilo idiomático de Go. Ele não apenas reduz a complexidade do código, mas também aumenta sua robustez ao garantir que recursos sejam liberados corretamente, mesmo em cenários de erro.

Dominar o `defer` é um passo essencial antes de avançar para concorrência, pois muitos padrões seguros com goroutines, canais e primitivas de sincronização dependem diretamente de seu uso correto.

## Tratamento de Erros em Go

O tratamento de erros em Go segue uma abordagem explícita e baseada em valores. Diferentemente de linguagens que utilizam mecanismos de exceção como `try/catch`, Go adota um modelo simples e previsível: funções que podem falhar retornam, além do resultado esperado, um valor do tipo `error`.

Esse modelo não é uma limitação da linguagem, mas uma decisão de design. Em Go, erros fazem parte do fluxo normal de execução e devem ser tratados de forma clara e deliberada pelo código chamador.

### Erros como valores

Em Go, `error` é uma interface. Isso permite que erros sejam criados, retornados, comparados, propagados ou enriquecidos com contexto adicional, da mesma forma que qualquer outro valor.

O padrão mais comum — e idiomático — para lidar com erros é a verificação explícita após a chamada de uma função:

```go
result, err := algumaFuncao()
if err != nil {
    // decidir como tratar o erro
    return err
}
```

Essa verificação explícita torna o fluxo do programa mais fácil de entender, evita efeitos colaterais implícitos e deixa claro onde e como cada erro é tratado.

**Exemplo prático**

O exemplo a seguir ilustra uma função que realiza um cálculo simples e retorna um erro caso uma condição inválida seja encontrada:

```go
...
func soma(numeros ...int) (int, error) {
    total := 0
    for _, n := range numeros {
        total += n
    }

    if total == 0 {
        return 0, errors.New("o resultado não pode ser zero")
    }

    return total, nil
}
...
```

Ao consumir essa função, o código chamador deve decidir como lidar com o erro retornado:

```go
func main() {
    total, err := soma(1, 2)
    if err != nil {
        return
    }
    fmt.Println("Resultado:", total)
}
```

### Tratamento versus propagação

Nem todo erro deve ser tratado imediatamente. Uma regra prática em Go é:

- **Trate o erro** quando você tem contexto suficiente para tomar uma decisão.
- **Propague o erro** quando não é possível resolvê-lo naquele nível da aplicação.

Propagar um erro geralmente significa retorná-lo ao chamador, preservando o fluxo explícito do programa.

### Boas práticas

- Nunca ignore erros silenciosamente.
- Evite funções auxiliares que apenas verificam `err != nil` sem executar nenhuma ação concreta.
- Sempre que possível, adicione contexto ao erro antes de propagá-lo.
- Trate erros no nível mais apropriado da aplicação.

### Considerações finais sobre erros

Em Go, erros são parte do contrato das funções e devem ser tratados com a mesma atenção que qualquer outro valor retornado. Essa abordagem explícita favorece código mais legível, previsível e fácil de manter, além de incentivar decisões conscientes sobre falhas e comportamentos excepcionais ao longo do fluxo da aplicação.

## Interfaces

Uma `interface` é uma coleção de metódos que um tipo concreto deve implementar para ser considerado uma instância dessa interface. Portanto, uma `interface` define, mas, não declara o comportamento do tipo.

Para exemplificar vamos usar os mesmos exemplos que usamos para criar métodos.

```go
package main

import "fmt"

// Geo interface base para figuras geométricas
type Geo interface {
	Area() float64
}

// Retangulo representa um retângulo
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Area calcula a are de um retângulo
func (r *Retangulo) Area() float64 {
	res := r.Largura * r.Altura
	return res
}

// Triangulo representa um triângulo
type Triangulo struct {
	Base   float64
	Altura float64
}

// Area calcula a are de um triângulo
func (t *Triangulo) Area() float64 {
	res := (t.Base * t.Altura) / 2
	return res
}

func imprimeArea(g Geo) {
	fmt.Println(g)
	fmt.Println(fmt.Sprintf("Área      : %0.2f", g.Area()))
}

func main() {
	r := Retangulo{
		Altura:  10,
		Largura: 5,
	}

	t := Triangulo{
		Base:   10,
		Altura: 5,
	}

	imprimeArea(&r)
	imprimeArea(&t)
}
```

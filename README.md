# A tour of goの復讐
最初に**大文字で**始まる文字は**エクスポートされた公式**からのものになる

    naked" return = returnの後に何も描かないこと

この場合はi = 1とj = 2となる  
`var i, j int = 1, 2`

### 関数の中では、 var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができる

`var i, j int = 1, 2` -> `i, j := 1, 2`と**関数内**なら省略できる

#### 関数も変数として扱うこともできる
関数の引数を受け取ることもでき、戻り値としての利用もできる

    func compute(fn func(float64, float64) float64) float64 {
    	return fn(3, 4)
    }
    
    func main() {
    	hypot := func(x, y float64) float64 {
    		return math.Sqrt(x*x + y*y)
    	}
    	fmt.Println(hypot(5, 12))
    
    	fmt.Println(compute(hypot))
    	fmt.Println(compute(math.Pow))
    }
    結果:
    13
    5
    81

#### 関数クロージャー
**クロージャー** = それ自身の外部から変数を参照する関数値  
この関数は、参照された変数へアクセスして変えることができる

    func adder() func(int) int {
    	sum := 0
    	return func(x int) int {
    		sum += x
    		return sum
    	}
    }
    
    func main() {
    	pos, neg := adder(), adder()
    	for i := 0; i < 10; i++ {
    		fmt.Println(
    			pos(i),
    			neg(-2*i),
    		)
    	}
    }
    結果:
    0 0
    1 -2
    3 -6
    6 -12

## 基本的な型
    bool
    
    string
    
    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr
    
    byte // uint8 の別名
    
    rune // int32 の別名
         // Unicode のコードポイントを表す
    
    float32 float64
    
    complex64 complex128

`var Hoge bool = false` -> `fmt.Printf("Type: %T Value: %v\n", Hoge, Hoge)`  
この記述をすることによって**Tで型を**表し、**v(value)で値**を表すことができる

#### zero values
    数値型(int,floatなど): 0
        bool型: false
        string型: "" (空文字列( empty string ))
**変数に初期値を書かないとデフォルトの与えられる**  
`var i int` -> `fmt.Plintf("%v", i)`となり出力として**0が出力**される

### 型変換
    シンプルに記述する場合
    i := 42
    f := float64(i)
    u := uint(f)
**型変換は明示的な変換が必要になる**  
##### :=で型推論した場合
    i := 42           // int
    f := 3.142        // float64
    g := 0.867 + 0.5i // complex128
と**自動で型を調整してくれる**  

## 定数(const)
定数の変数の先頭を大文字にすることが理想的   -> `const Hoge = "hello"`  
**文字(character)、文字列(string)、boolean、数値(numeric)のみで使える**  
型推論はできない

##### << と >>の意味
> https://teratail.com/questions/18602  
上記を参考に2真数に変更した際に値をかえる  

### For文
    func main() {
    	sum := 0
    	for i := 0; i < 10; i++ {
    		sum += i
    	}
    	
    i := 0 初期化ステートメント: 最初のイテレーション(繰り返し)の前に初期化が実行されます
    i < 10 条件式: イテレーション毎に評価されます
    i++ 後処理ステートメント: イテレーション毎の最後に実行されます
**※初期化と後処理ステートメントの記述は任意**  
-> 

    func main() {
    	sum := 1
    	for ; sum < 1000; {
    		sum += sum
    	}

**※while文もfor文でgoでは表す**  

    func main() {
    	sum := 1
    	for sum < 1000 {
    		sum += sum
    	}

**無限ループの書き方**  
    
    func main() {
    	for {
    	}
    }

##### Sqrt関数
`func Sqrt(x float64) float64` -> Sqrtは、xの平方根を返す

##### Pow関数
`func Pow(x, y float64) float64`  
 -> Powは、x**y、ベースxをy値で累乗した値を返す

### If文
**Ifのステートメント**  
If文の部分に簡単なステートメントを書くことができるようになっている

    func pow(x, n, lim float64) float64 {
    	if v := math.Pow(x, n); v < lim {
    		return v
    	}
    	return lim
    }

### Switch文
**breakは自動的に提供される** & **caseのものは定数でなくても平気**  

    パソコンのOSを知るプログラム
    func main()  {
    	fmt.Print("Go runs on")
    	switch os := runtime.GOOS; os {
    	case "darwin":
    		fmt.Println("os x.")
    	case "linux":
    		fmt.Println("Linux.")
    	default:
    		fmt.Printf("%s.\n", os)
    	}

##### Deferメソッド
**defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延**

    func main()  {
    	defer fmt.Println("world")
    	fmt.Println("hello")
    }

**stackされる場合**

    func main()  {
    	fmt.Println("counting")
    	for i := 0; i < 10; i++ {
    		defer fmt.Println(i)
    	}
    	fmt.Println("done")
    }
このように記述することで9, 8, 7, ....と出力される

## ポインタ
**正直全然理解することができないので参考サイトを参照**  
>図がわかりやすいもの
>https://note.com/aki2020/n/n9294337b19c8  
>C言語のやつだけど参考になる  
>https://qiita.com/tatachiy/items/7856931f4cacf2c5b743  
>アドレスとポインタ  
>https://qiita.com/Sekky0905/items/447efa04a95e3fec217f

**ポインタは値のメモリアドレスを指す**  
公式のサイト
https://go-tour-jp.appspot.com/moretypes/1

### 構造体(struct)
**フィールド( field )の集まり**

    type Vertex struct {
    	X int
    	Y int
    }
    
    func main() {
    	fmt.Println(Vertex{1, 2})
    }
    と記述すると{1 2}となる
    
**ドット(.)を使って今回だったらx, yにアクセスできる**  
    
    v := Vertex2{1, 2}
    	v.x = 4
    	fmt.Println(v.x)

**ポインタからのアクセス**
ポインタを利用して呼ぶこともできる  
    
    func main() {
    	v := Vertex{1, 2}
    	p := &v
    	p.X = 1e9
    	fmt.Println(v)
    }

### 配列(array)
**int型の配列を10個作成する宣言**->  `var a [10]int`  

    func main() {
    	var a [2]string
    	a[0] = "Hello"
    	a[1] = "World"
    	fmt.Println(a[0], a[1])
    	fmt.Println(a)
    
    	primes := [6]int{2, 3, 5, 7, 11, 13}
    	fmt.Println(primes)
    }
    
    結果:
    Hello World
    [Hello World]
    [2 3 5 7 11 13]

####　配列の操作(slice)
**配列を限定する**

    primes := [6]int{2,3,5,7,11,13}
    
    ここの部分で1~４までを取ってくる
    	var s []int = primes[1:4]
    	fmt.Println(s)
    	
    	結果:
    	[3 5 7]
**配列ストラクト**  
配列の中にオブジェクトが入っているイメージ  

    s := []struct {
    		i int
    		b bool
    	}{
    		{2, true},
    		{3, false},
    		{5, true},
    		{7, true},
    		{11, false},
    		{13, true},
    	}
    	fmt.Println(s)
    	結果:
    	[{2 true} {3 false} {5 true} {7 true}
    	 {11 false} {13 true}]
**スライスのデフォルト**
スライスするときは、それらの既定値を代わりに使用することで上限または下限を省略することができる

    s = s[:2]
    	fmt.Println(s)
    
    	s = s[1:]
    	fmt.Println(s)
**長さと容量の二つをもっている**  

    s := []int{2, 3, 5, 7, 11, 13}
    	printSlice(s)
    
    	// Slice the slice to give it zero length.
    	s = s[:0]
    	printSlice(s)
    
    	// Extend its length.
    	s = s[:4]
    	printSlice(s)
    
    	// Drop its first two values.
    	s = s[2:]
    	printSlice(s)
    	結果:
    	len=6 cap=6 [2 3 5 7 11 13]
        len=0 cap=6 []
        len=4 cap=6 [2 3 5 7]
        len=2 cap=4 [5 7]

**make関数を使用することでスライスの容量と長さを指定できる**  

    長さのみの指定
    a := make([]int, 5)  // len(a)=5
    容量のみの指定
    b := make([]int, 0, 5) // len(b)=0, cap(b)=5
    
**append関数で配列の追加をすることができるようになる**

    var s []int
    	printSlice3(s)
    
    	s = append(s, 0)
    	printSlice3(s)
    
    	s = append(s, 1)
    	printSlice3(s)
    
    	s = append(s, 2, 3, 4)
    	printSlice3(s)
    	
    	結果:
    	len=0 cap=0 []
            len=1 cap=1 [0]
            len=2 cap=2 [0 1]
            len=5 cap=6 [0 1 2 3 4]

## Range
**スライスや、マップ( map )をひとつずつ反復処理するために使う**  
-> 変数を二つ用意する必要がある。
1. index
1. indexの場所の要素のコピー


    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
    
    func main() {
    	for i, v := range pow {
    		fmt.Printf("2**%d = %d\n", i, v)
    	}
    }
    結果:
    2**0 = 1
    2**1 = 2
    2**2 = 4
    2**3 = 8
    2**4 = 16
    2**5 = 32
    2**6 = 64
    2**7 = 128

**indexや値は " _ "(アンダーバー) へ代入することで捨てることができるようん笑になる**  

    for i, _ := range pow
    for _, value := range pow
    
    もしインデックスだけが必要なのであれば、2つ目の値を省略
    
    for i := range pow


## Maps
**map はキーと値とを関連付けをする**  
0 = nil -> nilのマップはキーを持っておらず、キーの追加をすることもできない  
`make` 関数は指定された型のマップを初期化して、使用可能な状態で返す

    type Vertex struct {
    	Lat, Long float64
    }
    
    var m map[string]Vertex
    
    func main() {
    	m = make(map[string]Vertex)
    	m["Bell Labs"] = Vertex{
    		40.68433, -74.39967,
    	}
    	fmt.Println(m["Bell Labs"])
    }
    結果：
    {40.68433 -74.39967}
    
**Mapのリテラルにはkey(キー)が必要になってくる**  

    var m = map[string]Vertex{
    	"Bell Labs": Vertex{
    		40.68433, -74.39967,
    	},
    	"Google": Vertex{
    		37.42202, -122.08408,
    	},
    }
    結果：
    map[Bell labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

### mapの操作
**m へ要素(elem)の挿入や更新**
    
    m[key] = elem

**要素の取得**  

    elem = m[key]
    
**削除**  

    delete(m, key)
    
**キーに対する要素が存在するかどうかは、2つの目の値で確認**  

    elem, ok = m[key]
    存在している場合はtrueになりそれ以外はfalse


## メソッド(methods)
**クラス( class )のしくみはありませんが、型にメソッド( method )を定義できる**  
メソッド -> 特別なレシーバ( receiver )引数を関数に取る  
レシーバー -> func キーワードとメソッド名の間に自身の引数リストで表現  
**Abs メソッドは v という名前の Vertex 型のレシーバを持つことを意味**

    type Vertex struct {
    	X, Y float64
    }
    
    func (v Vertex) Abs() float64 {
    	return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }
    
    func main() {
    	v := Vertex{3, 4}
    	fmt.Println(v.Abs())
    }
    結果:
    5
> ここで3*3と4*4で9+16で25になりsqrtで平方根を求めるので5


**通常での書き方**

    func Ads(v Type1) float64  {
    	return math.Sqrt(v.x*v.x + v.y*v.y)
    }
    
    vからアクセスすることができるのはxとyとなる

**structの型だけではなく、任意の型(type)にもメソッドを宣言できる**

    type MyFloat float64
    
    func (f MyFloat) Abs() float64 {
    	if f < 0 {
    		return float64(-f)
    	}
    	return float64(f)
    }
    
    func main() {
    	f := MyFloat(-math.Sqrt2)
    	fmt.Println(f.Abs())
    }

### ※メソッドのポイントの注意点
    var v Vertex
    ScaleFunc(v, 5)  // Compile error!
    ScaleFunc(&v, 5) // OK
    
**↓**

    var v Vertex
    v.Scale(5)  // OK
    p := &v
    p.Scale(10) // OK
    
    メソッドがポインタレシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができる

### 逆の場合
    var v Vertex
    fmt.Println(AbsFunc(v))  // OK
    fmt.Println(AbsFunc(&v)) // Compile error!
    
**↓**

    var v Vertex
    fmt.Println(v.Abs()) // OK
    p := &v
    fmt.Println(p.Abs()) // OK
    
    変数の引数を取る関数は、特定の型の変数を取る必要がある


##これがすごく参考になる
    type Type5 struct {
    	x, y float64
    }
    
    func (v *Type5) Scale(f float64)  {
    	v.x = v.x * f
    	v.y = v.y * f
    }
    
    func (v *Type5) Abs() float64  {
    	return math.Sqrt(v.x*v.x + v.y*v.y)
    }
    
    func main()  {
    	v := &Type5{3, 4}
    	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
    	v.Scale(5)
    	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
    }
    
    結果:
    Before scaling: &{X:3 Y:4}, Abs: 5
    After scaling: &{X:15 Y:20}, Abs: 25

## インターフェース
**メソッドの集まりで定義する**  
そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができる

    基本的な書き方
    type Abser interface {
    	Abs() float64
    }

**わかりやすいまとめ**  
-> インターフェースは明示的に宣言する必要はない

    type I interface {
    	M()
    }
    type T struct {
    	S string
    }
    
    func (t T) M()  {
    	fmt.Println(t.S)
    }
    
    func main()  {
    	var i I=T{"hello"}
    	i.M()
    }
    結果: hello

インターフェースの値はタプルのように考えることができるようになっている  
    
    (value, type)

**ポインタがあるかないかで変わってくる**  
インターフェース自体の中にある具体的な値が nil の場合、メソッドは nil をレシーバーとして呼び出される

    i = t
    	describe(i)
    	i.M()
    
    	i = &T{"hello"}
    	describe(i)
    	i.M()
    	
    	結果:
            (<nil>, *main.T)
            <nil>
            (&{hello}, *main.T)
            hello


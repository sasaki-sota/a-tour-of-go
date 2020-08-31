# A tour of goの復讐
最初に**大文字で**始まる文字は**エクスポートされた公式**からのものになる

    naked" return = returnの後に何も描かないこと

この場合はi = 1とj = 2となる  
`var i, j int = 1, 2`

### 関数の中では、 var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができる

`var i, j int = 1, 2` -> `i, j := 1, 2`と**関数内**なら省略できる

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

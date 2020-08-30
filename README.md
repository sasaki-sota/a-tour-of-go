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



/**
 * 每个 Go 文件都属于且仅属于一个包
 * 一个Go应用程序都包含一个名为main包
 * 所有的包名都应该使用小写字母
 * 定义包名
 */
package main
/**
  *导入包
  * import "fmt" 或 import "fmt";import "os"
  * import ("fmt";"os")
  * imoport (
  *    "fmt"
  *    "os"
  * )
  * import fm "fmt" 别名
  */
import "fmt"
// 如果有 init() 函数则会先执行该函数
func init() {
	fmt.Println("初始化");
}

/**
 * 标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用 public属性面向对象,* 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的, private属性面向对象
 * Test 单一返回值, Go支持多返回值
 */
func Test(a int, b int) int {
	/**  
	 * 定义变量 var c int = 1 声明类型和赋值, 可以不赋值, int默认为0
	 *         c := 1 简洁声明变量 根据右侧值来推断类型
	 */  
	c := a + b // var c = a + b 因为右侧已经推断类型
	return c
}

func demo() (a bool, b bool) {
	return true, false
}

 // 入口函数
func main()  { // 左花括号必须跟声明在同一行
	fmt.Println("Hello World")
}
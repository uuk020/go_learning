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
  */
import "fmt"
 // 入口函数
func main()  {
	// 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用 public属性面向对象, 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的, private属性面向对象
	fmt.Println("Hello World")
}
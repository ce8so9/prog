package main
import "fmt"
type Person struct {
  Name string
}
func (p *Person) SayHello() {
  fmt.Println("Hello, ",p.Name)
}
type Dog struct {}
func (d *Dog) SayHello() {
  fmt.Println("Woof Woof")
}
type Cat struct{}
func (c *Cat) Answer(){
  fmt.Println("Meow")
}
type Friend interface {
  SayHello()
}
func Greet (f Friend) {
  f.SayHello()
}
func main() {
  var guy = new(Person)
  guy.Name = "Irvin"
  Greet(guy)
  var dog = new(Dog)
  Greet(dog)
  var cat = new(Cat)
  Greet(cat)
}

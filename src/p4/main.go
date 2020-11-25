package main
 
import "fmt"
 
type Stream interface{
    read() string
    write(string)
    close()
}
 
func writeToStream(stream Stream, text string){
    stream.write(text)
}
func closeStream(stream Stream){
    stream.close()
}
 
// структура файл
type File struct{
    text string
}
// структура папка
type Folder struct{}
 
// реализация методов для типа *File
func (f *File) read() string{
    return f.text
}
func (f *File) write(message string){
    f.text = message
    fmt.Println("Запись в файл строки", message)
}
func (f *File) close(){
    fmt.Println("Файл закрыт")
}
// релизация методов для типа *Folder
func (f *Folder) close(){
    fmt.Println("Папка закрыта")
}

func (f *Folder) read() string{
    return "sldkcksd"
}
 
func main() {
     
    myFile2 := &File{}
	(*myFile2).close()
}
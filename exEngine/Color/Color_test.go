package Color_test

import (
	"ExEngine/Color"
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println(Color.Black("黑色字体"))
	fmt.Println(Color.Red("红色字体"))
	fmt.Println(Color.Blue("蓝色字体"))
}

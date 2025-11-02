# point-distance

Простая библиотека на Go для вычисления евклидова расстояния между двумя точками на плоскости.

## Возможности

- Инкапсуляция координат точки (`x`, `y`) в структуре `Point`
- Конструктор `NewPoint(x, y)` для безопасного создания точек
- Метод `Distance(other Point)` для расчёта расстояния по формуле:  
  \( d = \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2} \)

## Пример использования

```go
package main

import (
    "fmt"
    "github.com/danilovaalina/point-distance"
)

func main() {
    p1 := NewPoint(0, 0)
    p2 := NewPoint(3, 4)
    fmt.Printf("Расстояние между точками: %.2f\n", p1.Distance(p2)) // Вывод: 5.00
}
```

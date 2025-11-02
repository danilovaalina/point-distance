# point-distance

Простая библиотека на Go для вычисления евклидова расстояния между двумя точками на плоскости.

## Возможности

- Инкапсуляция координат точки (`x`, `y`) в структуре `Point`
- Конструктор `NewPoint(x, y)` для безопасного создания точек
- Метод `Distance(other Point)` для расчёта расстояния по формуле:  **d = √((x₂ - x₁)² + (y₂ - y₁)²)**

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

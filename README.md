# Guía 10 - Backtracking
## Implementaciones de las diapositivas

En las siguientes carpetas se encuentran las implementaciones del código suministrado en las clases:

- `/ratmaze`, problema del laberinto, resuelto con backtracking
- `/sudoku`, resolución de un sudoku clásico, utilizando backtracking
- `/main`, ejecuciones adicionales con fines interactivos

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

### Backtracking

Usando los conceptos de backtracking resuelva los siguientes ejercicios. 

1. Problema de las N reinas. A mitad del siglo XIX Max Bezzel publico el problema de las ocho reinas, el cual consistía en poder colocar 8 reinas en un tablero convencional de ajedrez sin que ninguna reina amenaze a otra (es decir no coincida en su misma fila, columna y diagonal). Este problema ya es un clasico que dispone de 92 soluciones distintas, pero solo 12 soluciones unicas (el resto se compone de rotaciones y reflejos o espejados).
Se pide diseñar una solución más generica donde se reciba un número N que representará la cantidad de reinas y también el tamaño del tablero con el cual trabajar. La salida debe ser la cantidad de soluciones distintas posibles de acomodar esas N reinas en un tablero de NxN.

2. Problema de las N reinas únicas. Se pide que la solucion del problema anterior sea solo de las soluciones únicas, no de todas las distintas.

3. Problema de la mochila (Knapsack). Este es uno de los 21 problemas de [NP completos](https://es.wikipedia.org/wiki/NP-completo) de Richard Karp, y ya [mencionado en 1896](https://doi.org/10.1112%2Fplms%2Fs1-28.1.486) por Mathews, G. B. Consiste en un problema de optimización combinatoria, donde se espera poder llenar una "mochila" con un peso limitado, por una cantidad de objetos, cada uno con un peso y valor específico, máximizando el valor total almacenado. Hay una simplificación del mismo, denominado problema de la mochila 0-1 donde un objeto puede estar o no dentro de la mochila, por completo.
Se pide diseñar una solución utilizando backtraking para este problema de la mochila 0-1.

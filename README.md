# Reto Técnico

En este repositorio podra encontrar una prueba tecnica para la empresa mercado libre

## Prerrequisitos

- Tener instalado Go
- Docker o similar que nos permita crear contenedores

## Tecnologías, herramientas y librerias
- [ Gin Web Framework](https://gin-gonic.com/) framework principal con el que exponemos endpoints y nos ayuda con la comunicacion a la base de datos
- [Redis](https://redis.io/)es un almacén de datos en memoria de código abierto que ofrece tiempos de respuesta por debajo del milisegundo.
## Construido
-   [ Visual Studio Code ](https://code.visualstudio.com/) Editor  donde se puede compilar el proyecto.
- -   [Postman](https://www.postman.com/) es una aplicación que permite probar API web.

## Como ejecutarlo

Primero descargamos todas las dependencias del proyecto para esto usamos el siguiente comando
`go get ./...

Luego ejecutamos el archivo principal que en este caso es main para esto ejecutmos el siguiente comando
`go run main.go

Con esto tendremos la app funcionando, pero este proyecto requiere de dos container uno que sera nuestra base de datos y tel otro el que nos ayudara a almacenar cache. Para en la raiz del projecto ejecutamos el siguiente comando
`docker-compose up -d

Esto nos ejecutara los container necesarios para que funcione. 

## Arquitectura



# Challenge teorico

## Procesos, hilos y corrutinas

- Un caso en el que usarías procesos para resolver un problema y por qué.
- Un caso en el que usarías threads para resolver un problema y por qué.
- Un caso en el que usarías corrutinas para resolver un problema y por qué.
## Optimización de recursos del sistema operativo
Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de
ellos información en una API HTTP. ¿Cómo lo harías? Explicar.

## Análisis de complejidad
- Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con
complejidades O(n2), O(n3), O(2n) y O(n log n), respectivamente, ¿Cuál de los
algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.

- Asume que dispones de dos bases de datos para utilizar en diferentes
problemas a resolver. La primera llamada AlfaDB tiene una complejidad de O(1)
en consulta y O(n2) en escritura. La segunda llamada BetaDB que tiene una
complejidad de O(log n) tanto para consulta, como para escritura. ¿Describe en
forma sucinta, qué casos de uso podrías atacar con cada una?

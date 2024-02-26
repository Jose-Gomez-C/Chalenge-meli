# Reto Técnico

En este repositorio podrá encontrar una prueba técnica para la empresa mercado libre

## Prerrequisitos

- Tener instalado Go
- Docker o similar que nos permita crear contenedores

## Tecnologías, herramientas y librerías 
- [ Gin Web Framework](https://gin-gonic.com/) framework principal con el que exponemos endpoints y nos ayuda con la comunicación  a la base de datos
- [Redis](https://redis.io/)es un almacén de datos en memoria de código abierto que ofrece tiempos de respuesta por debajo del milisegundo.
## Construido
-   [ Visual Studio Code ](https://code.visualstudio.com/) Editor  donde se puede compilar el proyecto.
- -   [Postman](https://www.postman.com/) es una aplicación que permite probar API web.

## Como ejecutarlo

Primero descargamos todas las dependencias del proyecto para esto usamos el siguiente comando

`go get ./...

Luego ejecutamos el archivo principal que en este caso es main para esto ejecutamos el siguiente comando

`go run main.go

Con esto tendremos la app funcionando, pero este proyecto requiere de dos container uno que será nuestra base de datos y el otro el que nos ayudara a almacenar cache. Para en la raíz del proyectó ejecutamos el siguiente comando

`docker-compose up -d

Esto nos ejecutara los container necesarios para que funcione. 

## Arquitectura
![untitled](https://github.com/Jose-Gomez-C/Chalenge-meli/assets/46968912/9f7bc33f-7bae-410d-b374-1a25cb3958bf)

# Challenge teorico

## Procesos, hilos y corrutinas

- Un caso en el que usarías procesos para resolver un problema y por qué.
	- Imaginemos que estamos diseñando una solución para enviar y recibir códigos de descuento, este es un proceso manual y afecta a varias tecnologías que ya tenemos en nuestra empresa, aparte debemos tener una forma fácil y sencilla para que el usuario final pueda actualizar sus datos. 
	
	   en este caso usar procesos seria la mejor opción ya que por cada tecnología consultada vamos a tener una salida que se podría  usar en la siguiente o al finalizar el proceso. por otro lado tendríamos el proceso de que el usuario pueda actualizar  su información lo cual es vital para el proceso principal ya que es el objetivo del proyecto.
	
	   por ultimo y otro beneficio de usar procesos en este caso es que tendríamos un orden y aparte cada proceso tendría un fin en especifico y su escalabilidad estaría casi asegurada
- Un caso en el que usarías threads para resolver un problema y por qué
	- Un caso muy típico es un chat ya que tenemos múltiples usuarios que pueden interactuar y usar un solo hilo puede ser un bloqueante para el servidor. por esto es buena idea usar un hilo para cada usuario.

	   otro beneficio es que el cliente podrá enviar y recibir mensajes al instante ya que el los threds nos permitiría hacer esto.
- Un caso en el que usarías corrutinas para resolver un problema y por qué.
	- Imaginemos que estamos desarrollando una app de análisis de datos y para esto debemos procesar grandes cantidades de datos al usar corrutinas obtenemos estos benefeicios
		- Procesamiento concurrente de datos
		- Operaciones de entrada y salida no bloqueantes
		- Flexibilidad en la gestión de flujos de datos
		- Escalabilidad y recursos
## Optimización de recursos del sistema operativo
Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de
ellos información en una API HTTP. ¿Cómo lo harías? Explicar.
- primero usar una estrategia de divide y conquista por lo cual debemos dividir el lotes mas pequeños y procesar estos lotes de forma concurrente. Aparte implementar cache compartido para estos lotes con el fin de no congestionar la API con solicitudes repetidas. Por otro lado implementar un manejo de errores y reintentos , esto garantizará que obtengamos la mayor cantidad de información.

## Análisis de complejidad
- Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con
	complejidades O(n2), O(n3), O(2n) y O(n log n), respectivamente, ¿Cuál de los
	algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.
	- descartaría el C ya que tiene una complejidad exponencial  lo cual lo hace super ineficiente con entradas de tamaño mediano por lo cual seria el mas lento. el que tomaria depende mucho de la entrada pero tomaría el D aunque puede ser malo para entradas de poco tamaño a medio y gran tamaño tiene un rendimiento eficiente. Los otros dos algoritmos se podrían usar si la entrada es pequeña.

- Asume que dispones de dos bases de datos para utilizar en diferentes
	problemas a resolver. La primera llamada AlfaDB tiene una complejidad de O(1)
	en consulta y O(n2) en escritura. La segunda llamada BetaDB que tiene una
	complejidad de O(log n) tanto para consulta, como para escritura. ¿Describe en
	forma sucinta, qué casos de uso podrías atacar con cada una?
	- Usaría AlfaDB solo para consultar por us velocidad de consulta, ideal para mantener el cache de nuestra aplicación.
	- BetaDb La usaría para almacenar cantidades grandes de datos, y requiere de consultas, ideal para aplicaciones que realizan búsquedas complejas o recuperan datos basados en criterios específicos.


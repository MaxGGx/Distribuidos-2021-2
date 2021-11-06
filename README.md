# Distribuidos-2021-2

Repo de tareas de distribuidos

Nombres: <br>
Luis Gonzalez 201892004-5<br>
Sebastian Martinez 201873519-1<br>
Iñaki Oyarzun 201873620-1<br>

Instrucciones de uso:

Se necesitan un total de 7 consolas para ejecutar el código y hacer las revisiones necesarias. Estas 7 instancias deben ser:

2 instancias conectadas a la máquina dist33, de las cuales 1 instancia es para ejecutar el Lider y la otra instancia es para ejecutar dataNode1

2 instancias conectadas a la máquina dist34, de las cuales 1 instancia es para ejecutar el Pozo y la otra instancia es para ejecutar dataNode2

1 instancia conectadas a la máquina dist35, que corresponde a la ejecución de nameNode

2 instancias conectadas a la máquina dist36, de las cuales 1 instancia es para ejecutar el juego (como usuario) y la otra instancia es para ejecutar dataNode3


Se mencionará las instrucciones a realizar por cada máquina inicialmente, y después se muestra el orden de ejecución (esto quiere decir, no ejecutar secuencialmente las instrucciones para instancia 1, instancia 2. Leer primero como se ejecutará cada máquina y luego proceder a realizar los comandos)

Para la máquina dist33:

	Instancia 1:

		Una vez conectada a alumno@dist33, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "tarea/Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/Lider/server"

		una vez en el directorio .../server, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)

	Instancia 2:

		Una vez conectada a alumno@dist33, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "tarea/Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/DataNode y NameNodes/dataNode1/server"

		una vez en el directorio .../server, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)



Para la máquina dist34:

	Instancia 3:

		Una vez conectada a alumno@dist34, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/Pozo/server"

		una vez en el directorio .../server, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)

	Instancia 4:

		Una vez conectada a alumno@dist34, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/DataNode y NameNodes/dataNode2/server"

		una vez en el directorio .../server, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)



Para la máquina dist35:

	Instancia 5:

		Una vez conectada a alumno@dist35, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/DataNode y NameNodes/nameNode/server-client"

		una vez en el directorio .../server-client, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)



Para la máquina dist36:

	Instancia 6:

		Una vez conectada a alumno@dist36, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/M4"

		una vez en el directorio .../M4, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)

	Instancia 7:

		Una vez conectada a alumno@dist36, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/DataNode y NameNodes/dataNode3/server"

		una vez en el directorio .../server, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)



Orden de ejecución (como iniciar las máquinas virtuales):





Consideraciones:
- Se asume que los archivos de ronda se identifican como las etapas

- Se asume que las entradas siempre serán correctas

- Considerar que en caso de requerirse reiniciar el código, se deben eliminar de la carpeta data de cada datanode todos los archivos .txt que se crean en el documento 

- Se toma en cuenta que el jugador una vez que gana deberá pasar por las demás etapas de igual forma sin tomar en cuenta sus valores para llegar a obtener el mensaje de ganador.

- Tomar en consideración que debido al azar que manejan los BOTS, estos la mayoría de las veces no podrán pasar mas allá de la ronda 1, se recomienda para ello si se quiere analizar mas allá de dicha etapa, modificar la condición en la línea 461 del archivo de líder (Líder/server/main.go) para que no mate a la mayoría de los BOTS que no lograron sumar los 21. Esto podria ser por ejemplo que los BOTS no sumen 21 sino que 10 o menos para hacer pasar a gran parte de ellos a la siguiente etapa.
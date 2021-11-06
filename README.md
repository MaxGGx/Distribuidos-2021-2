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

2 instancias conectadas a la máquina dist36, de las cuales 1 instancia es para ejecutar el Lider y la otra instancia es para ejecutar dataNode1


Se mencionará las instrucciones a realizar por cada máquina inicialmente, y después se muestra el orden de ejecución

Para la máquina dist33:

	Instancia 1:

		Una vez conectada a alumno@dist33, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "tarea/Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/Lider/server"

		una vez en el directorio .../server, se ejecuta el comando $ make run

	Instancia 2:

		Una vez conectada a alumno@dist33, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "tarea/Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "FINAL DIST/DataNode y NameNodes/dataNode1/server"

		una vez en el directorio .../server, se ejecuta el comando $ make run




====== Laboratorio Nº3 Distribuidos - Grupo 8======
Nombres:
-Luis Gonzalez-------> 201892004-5
-Sebastian Martinez--> 201873519-1
-Inaki Oyarzun-------> 201873620-1

====== CONSIDERACIONES ======

- Se considera que no se realizara por parte de los informantes algun comando durante un proceso de MERGE
- Se asume que el usuario ingresara correctamente los comandos y sin errores de tipeo, respetando las mayusculas y minusculas
(tanto para cuando se cree una ciudad/planeta como comandos)

====== INSTRUCCIONES DE EJECUCION ======

Se necesitan un total de 7 consolas para ejecutar el código y hacer las revisiones necesarias. Estas 7 instancias deben ser:

1 instancia conectadas a la máquina dist33, que será para ejecutar el Broker 

2 instancias conectadas a la máquina dist34, de las cuales 1 instancia es para ejecutar Leia y la otra instancia es para ejecutar Fulcrum1

2 instancias conectadas a la máquina dist35, de las cuales 1 instancia es para ejecutar AhsokaT y la otra instancia es para ejecutar Fulcrum2

2 instancias conectadas a la máquina dist36, de las cuales 1 instancia es para ejecutar AlmiranteT y la otra instancia es para ejecutar Fulcrum3

Se mencionará las instrucciones a realizar por cada máquina inicialmente, y después se muestra el orden de ejecución (esto quiere decir, no ejecutar secuencialmente las instrucciones para instancia 1, instancia 2. Leer primero como se ejecutará cada máquina y luego proceder a realizar los comandos)

Para la máquina dist33:

	Instancia 1:

		Una vez conectada a alumno@dist33, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "tarea/Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "Lab3/Code/Broker"

		una vez en el directorio .../Broker, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)


Para la máquina dist34:

	Instancia 2:

		Una vez conectada a alumno@dist34, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "tarea/Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "Lab3/Code/Leia"

		una vez en el directorio .../Leia, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)

	Instancia 3:

		Una vez conectada a alumno@dist34, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "Lab3/Code/Fulcrums/Fulcrum1"

		una vez en el directorio .../Fulcrum1, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)



Para la máquina dist35:

	Instancia 4:

		Una vez conectada a alumno@dist35, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "Lab3/Code/Informantes/AhsokaT"

		una vez en el directorio .../AhsokaT, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)

	Instancia 5:

		Una vez conectada a alumno@dist35, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "Lab3/Code/Fulcrums/Fulcrum2"

		una vez en el directorio .../Fulcrum2, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)




Para la máquina dist36:

	Instancia 6:

		Una vez conectada a alumno@dist36, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "Lab3/Code/Informantes/AlmiranteT"

		una vez en el directorio .../AlmiranteT, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)

	Instancia 7:

		Una vez conectada a alumno@dist36, primero se tiene que cambiar al directorio donde se encuentra el repositorio Git, con el comando $ cd "Distribuidos-2021-2"

		El segundo paso es asegurarse de trabajar en la rama "inaki", por lo que se ejecuta el commando $ git checkout inaki

		Luego, se ejecuta el comando $ cd "Lab3/Code/Fulcrums/Fulcrum3"

		una vez en el directorio .../Fulcrum3, se ejecuta el comando $ make run (ESPERAR AL ORDEN DE INSTANCIACIÓN EN LA SECCIÓN SIGUIENTE)



Orden de ejecución (como iniciar las máquinas virtuales):

Para ejecutar las diferentes instancias desde la 1 hasta la 7, se sugiere ubicarse en el directorio correspondiente (ejemplo de la máquina Instancia 1: ubicarse en .../server); una vez ubicadas las 7 instancias en su directorio, ejecutar el comando "make run" en las instancias en el siguiente orden:

primero los 3 data nodes, luego el name node, luego el pozo, luego el lider, por último el jugador; es decir:

Instancia 2, Instancia 4, Instancia 7, Instancia 5, Instancia 3, Instancia 1, Instancia 6.

Una vez iniciadas todas las instancias, en la instancia del lider se dará la opción de controlar el juego, y en la instancia del jugador se pueden realizar las jugadas.



Consideraciones:
-El archivo .zip descargado de github no permite elegir una sola rama para descargar, por lo que se sugiere revisar al abrir el .zip, que se esté en la rama adecuada "inaki"

-Todas las máquinas virtuales cuentan con un readme llamado README.md que se encuentra en el directorio "Distribuidos-2021-2"

-Si se ve que una de las consolas se queda detenida (sea lider o el jugador), revisar la otra (jugador o lider) por si falta algún input.

- Se asume que los archivos de ronda se identifican como las etapas

- Se asume que las entradas siempre serán correctas

- Considerar que en caso de requerirse reiniciar el código, se deben eliminar de la carpeta data de cada datanode todos los archivos .txt que se crean en el documento 

- Se toma en cuenta que el jugador una vez que gana deberá pasar por las demás etapas de igual forma sin tomar en cuenta sus valores para llegar a obtener el mensaje de ganador.

- Tomar en consideración que debido al azar que manejan los BOTS, estos la mayoría de las veces no podrán pasar mas allá de la ronda 1, se recomienda para ello si se quiere analizar mas allá de dicha etapa, modificar la condición en la línea 461 del archivo de líder (Líder/server/main.go) para que no mate a la mayoría de los BOTS que no lograron sumar los 21. Esto podria ser por ejemplo que los BOTS no sumen 21 sino que 10 o menos para hacer pasar a gran parte de ellos a la siguiente etapa.

-Se toma en cuenta que existen prints para el lider los cuales indican las acciones realizadas. Una de ellas es cuando se registra en name node exitosamente la movida o jugada del usuario ("Agregado Exitosamente"). Ademas de ello cuando se consulta sobre el historial de un jugador, la mostrara en la forma de historial de ronda/etapa 1 a 3, en donde si no exite data sobre dicha ronda, se avisará por pantalla del lider. Finalmente, las demás interfaces de NameNode, DataNode y Pozo, mostraán un mensaje de la forma en que sea mas facil identificar el traspaso de mensajes que se realiza.

====== FIN DE README ======
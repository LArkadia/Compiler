# Compiler

# Analizador léxico
Este analizador léxico es capaz de reconocer los siguientes elementos:
- Identificadores
- Palabras reservadas: and else false for fun if null or print return true var while
- Números: enteros, con punto decimal y con exponente
- Comentarios: de una sola línea y multilínea. No generan token.
- Cadenas
- Símbolos: < <= > >= ! != = == + - * / { } ( ) , . ;
## Caracteristicas adicionales
- Incluye una nterfaz de línea de comandos que permita introducir cadenas de manera repetitiva (como el
intérprete de Python) y lectura de archivos a través de argumentos a la aplicación. No debe contener
mensaje como: “introduce la cadena...”
- Incluye palabras reservadas y elementos del lenguaje previamente especificados.
- Tiene un funcionamiento basado en AFD. 

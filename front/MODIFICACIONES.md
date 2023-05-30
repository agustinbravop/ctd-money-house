# Modificaciones para adaptar el front dado a nuestro back end específico.

## Resumen

1. **Arreglen los endpoints de la API en `/src/utils/api/index.ts`**.
2. **Refactoricen los campos de los tipos en `/src/types/`**.
3. **Busquen los campos hardcodeados que quedaron por ahí**.

No se puede tocar la parte estética, pero hay bastantes cambios funcionales para que coincida. Observaciones:

**Hay comentarios que comienzan con '// API:'**, con los que dejé observaciones. Si los buscan les va a resultar más
fácil ver qué detalles del código tocar.

## Recordatorios

- Recordar el CORS Allow All Origins, dado que la app de React se levanta en una url distinta al server de Java/Go.
- Cuando prueben el login posiblemente necesiten borrar el token que les queda en el `localStorage`.
- Los forms de login y register tienen validaciones que quizás no son las mismas del back. Las pueden encontrar en el
  archivo `/src/utils/formValidation/index.ts`.

### Endpoints

En `/src/utils/api/index.ts` van a encontrar todas las funciones que realizan llamadas a la API. Deben **cambiar los
endpoints** a los que ustedes hallan definido.

### Tipos

En la carpeta `/src/types/` van a poder cambiar campos de los tipos que tiene el front. Recuerden que al front solo
importa el json de la response, no los nombres de las clases de Java o los campos de las tablas de la base de datos.
Recuerden que este refactor debe también modificar toda otra linea de código que acceda al campo User.phone (yo lo hice
automáticamente con el IDE WebStorm de JetBrains). Otra alternativa es no tocar esta parte y cambiar el json que
devuelve el back, lo que les sea más facil.

### Login y Register

El form que está en `/src/pages/Login/index.tsx` onSubmit lee el campo 'accessToken' de la response. Nosotros lo tuvimos
que cambiar a 'access_token'.
En el register tuvimos que cambiar el dni que se parseaba a number pero para el back es un string.
El form de register a veces se rompe cuando valida algunos campos. Es por desestructurar un `undefined` en el
componente `ErrorMessage`. No lo arreglamos. **El front tiene validaciones que el back puede no tener**.

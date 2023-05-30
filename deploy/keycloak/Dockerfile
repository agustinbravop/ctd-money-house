FROM quay.io/keycloak/keycloak:19.0.3

ENV KEYCLOAK_ADMIN=admin
ENV KEYCLOAK_ADMIN_PASSWORD=admin
ENV USERS_CLIENT_SECRET=aBC5irvTeiVOhF4nNBRN2ZZ0JLh35N8L

WORKDIR /opt/keycloak

# money-house-realm.json contiene la configuración del realm y clients.
# Keycloak utiliza aplica esa config al levantar el containter por primera vez.
COPY money-house-realm.json data/import/

ENTRYPOINT [ "bin/kc.sh" ]

# Argumentos de CMD se pueden sobreescribir al levantar el contenedor.
# Ej: Nos permite usar 'start' en lugar de 'start-dev' para producción.
# '--import-realm' crea un realm a partir del JSON en data/import/. 
# Es necesario para configurar keycloak con el realm 'money-house'.
CMD [ "start-dev", "--import-realm" ]
## Ejecución a partir de la fuente
Ajustando los valores correctos de las variables, ejecutar el siguiente comando:
```bash
API_SERVER_PORT=puerto_a_servir \
	DB_HOST=ip_de_bd \
	DB_PORT=puerto_de_bd \
	DB_USER=usuario_de_bd \
	DB_PASSWORD=clave_super_segura \
	DB_NAME=nombre_de_base_de_datos \
	DB_SSL_MODE=disable \
	DB_MAXIDLECONNS=10 \
	DB_MAXOPENCONNS=10 \
	go run cmd/main.go
```

## Regeneración de swagger.yaml
Se debe ejecutar el siguiente comando, en caso de que se realicen actualizaciones en el código, que afecten la documentación del mismo
```bash
~/go/bin/swagger generate spec -o ./view/swagger.yaml --scan-models
~/go/bin/swagger mixin \
		./view/init.yml \
		./view/swagger.yaml \
		--format yaml \
		-o ./view/swagger.yaml
```

go mod init personas



## Construcción de imagen a partir de Dockerfile multi stage
```bash
docker build -t personas:1.0.1 .
```

## Construcción de Contenedor a partir de Imagen
```bash
docker run --rm  -p 8080:8080 \
	-e API_SERVER_PORT=8080 \
	-e DB_HOST=209.50.51.139 \
	-e DB_PORT=3306 \
	-e DB_USER=root \
	-e DB_PASSWORD=clave_super_secreta \
	-e DB_NAME=epi_bd \
	-e DB_MAXIDLECONNS=10 \
	-e DB_MAXOPENCONNS=10 \
	--name micro \
	personas:1.0.1
```
### Linke de prueba http://localhost:8080/console/
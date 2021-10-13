 # Servicio de Alertas a Corto Plazo GO
Busca las alertas del [Servicio Meteorológico Nacional](https://www.smn.gob.ar/avisos_a_muy_corto_plazo) utilizando el [endpoint](https://ws.smn.gob.ar/alerts/type/AC)

El servicio esta dockerizado pudiendo levantarlo usando el archivo Dockerfile o docker-compose.yml

---

### Levantar servicio
* ***Dockerfile***
```sh
docker build --pull --rm -f "Dockerfile" -t alertas_a_corto_plazo_go:latest "."
docker run --rm -d  -p 5000:5000/tcp alertas_a_corto_plazo_go:latest   (run in backgroung)
docker run --rm -it  -p 5000:5000/tcp alertas_a_corto_plazogo:latest   (run interactive)
```
* ***docker-compose.yml***
```sh
docker-compose -f "docker-compose.yml" up -d --build    (no es interactivo)
```

---
### Endpoints
```sh
curl --location --request GET 'http://localhost:5000/weather/sort-alert/all'
```


echo "Compilando Api Challenge"
./buildlinux.sh
echo "Copiando a carpeta dist"
mv build/apichallenge container/dist/apichallenge
cd container/
echo "Ejecutando Docker Compose"
docker-compose up -d --build

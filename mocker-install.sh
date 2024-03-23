#!/bin/bash

# Definir la versión de mockery que deseas descargar o clonar
VERSION="v2.32.0"

# Verificar si ya se tiene instalado Go en el sistema
if ! command -v go &>/dev/null; then
    echo "Go no está instalado en el sistema. Instalando Go..."

    # Actualizar lista de paquetes
    sudo apt update

    # Instalar Go
    sudo apt install -y golang

    # Verificar la instalación
    if command -v go &>/dev/null; then
        echo "Go ha sido instalado correctamente."
    else
        echo "Ha ocurrido un error al instalar Go."
        exit 1
    fi
fi


# Descargar o clonar el repositorio de mockery
echo "Descargando mockery versión $VERSION..."
if [ ! -d mockery ]; then
    git clone --branch $VERSION --depth 1 https://github.com/vektra/mockery.git
else
    cd mockery
    git pull origin $VERSION
    cd ..
fi

# Compilar el binario de mockery
echo "Compilando mockery..."
cd mockery
go build

# Copiar el binario de mockery a /usr/bin
echo "Copiando mockery a /usr/bin (requiere privilegios de administrador)..."
sudo cp mockery /usr/bin/

# Verificar que mockery esté disponible en el sistema
if command -v mockery &>/dev/null; then
    echo "mockery versión $VERSION ha sido instalado correctamente."
    cd ..
    rm -rf mockery/
else
    echo "Ha ocurrido un error al instalar mockery."
fi

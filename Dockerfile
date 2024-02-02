# Используем официальный образ Golang как базовый
FROM golang:latest as builder

# Устанавливаем libvips
RUN apt-get update && apt-get install -y \
    build-essential \
    wget \
    git \
    pkg-config \
    python3-pip \
    python3-setuptools \
    python3-wheel \
    libexpat1-dev \
    librsvg2-dev \
    libpng-dev \
    libjpeg-dev \
    libwebp-dev \
    libexif-dev \
    liblcms2-dev \
    libglib2.0-dev \
    liborc-dev \
    libgirepository1.0-dev \
    gettext

RUN apt-get install -y meson ninja-build

ARG VIPS_VER=8.14.2
ARG VIPS_DLURL=https://github.com/libvips/libvips/releases/download

RUN cd /usr/local/src \
    && wget ${VIPS_DLURL}/v${VIPS_VER}/vips-${VIPS_VER}.tar.xz \
    && tar xf vips-${VIPS_VER}.tar.xz \
    && cd vips-${VIPS_VER} \
    && meson setup build --buildtype=release \
    && cd build \
    && meson compile \
    && meson test \
    && meson install

RUN ldconfig
RUN export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/vips/lib

# Копируем исходный код в контейнер
WORKDIR /app
COPY . .

# # Собираем приложение
# RUN go build -o main .

# # Запускаем приложение
# CMD ["./main"]

CMD make build
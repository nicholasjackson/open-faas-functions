FROM golang:1.10

RUN set -x; \
    VER="17.12.0-ce"; \
    curl -L -o /tmp/docker-$VER.tgz https://download.docker.com/linux/static/stable/x86_64/docker-$VER.tgz; \
    tar -xz -C /tmp -f /tmp/docker-$VER.tgz; \
    mv /tmp/docker/* /usr/bin; \
    curl -L https://github.com/docker/compose/releases/download/1.19.0/docker-compose-`uname -s`-`uname -m` -o /usr/bin/docker-compose; \
    chmod +x /usr/bin/docker-compose

# Add opencv required by facedetect function
RUN apt-get update; \
    apt-get install -y build-essential cmake; \
    apt-get install -y qt5-default libvtk6-dev; \
    apt-get install -y zlib1g-dev libjpeg-dev libwebp-dev libpng-dev libtiff5-dev libjasper-dev libopenexr-dev libgdal-dev; \
    apt-get install -y libdc1394-22-dev libavcodec-dev libavformat-dev libswscale-dev libtheora-dev libvorbis-dev libxvidcore-dev libx264-dev yasm libopencore-amrnb-dev libopencore-amrwb-dev libv4l-dev libxine2-dev; \
    apt-get install -y libtbb-dev libeigen3-dev; \
    apt-get install -y python-dev python-tk python-numpy python3-dev python3-tk python3-numpy; \
    apt-get install -y doxygen; \
    apt-get install -y unzip wget

RUN wget https://github.com/opencv/opencv/archive/3.4.0.zip; \
    unzip 3.4.0.zip; \
    rm 3.4.0.zip; \
    mv opencv-3.4.0 OpenCV; \
    cd OpenCV; \
    mkdir build; \
    cd build; \
    cmake -DWITH_QT=ON -DWITH_OPENGL=ON -DFORCE_VTK=ON -DWITH_TBB=ON -DWITH_GDAL=ON -DWITH_XINE=ON -DBUILD_EXAMPLES=ON -DENABLE_PRECOMPILED_HEADERS=OFF ..; \
    make -j8; \
    make install; \
    ldconfig; \
    cd ../..; \
    rm -rf OpenCV

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

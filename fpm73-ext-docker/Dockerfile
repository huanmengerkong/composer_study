FROM php:7.4-fpm
RUN apt-get update && apt-get install -y \
        libfreetype6-dev \
        libjpeg62-turbo-dev \
        libpng-dev \
        git \
        zip \
        unzip \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-install mysqli \
    && docker-php-ext-install bcmatch \
    && docker-php-ext-install pdo \
    && docker-php-ext-install pdo_mysql \
RUN apt-get install -y librdkafka-dev \
    && pecl install rdkafka \
    && docker-php-ext-enable rdkafka
COPY ./phpredis-5.3.3.tar.gz .
RUN tar -zxvf ./phpredis-5.3.3.tar.gz \
    && mkdir -p /usr/src/php/ext \
    && mv phpredis-5.3.3 /usr/src/php/ext/redis \
    && docker-php-ext-install redis \
    && pecl install xdebug-2.8.1 mongodb \
    && docker-php-ext-enable redis xdebug mongodb \
    && php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" \ 
    && php composer-setup.php --install-dir=/usr/local/bin --filename=composer \ 
    && php -r "unlink('composer-setup.php');"                                             

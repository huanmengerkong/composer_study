<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInitd6b93bc241e89d4064739e7ff51855e0
{
    public static $prefixLengthsPsr4 = array (
        'P' => 
        array (
            'Predis\\' => 7,
        ),
    );

    public static $prefixDirsPsr4 = array (
        'Predis\\' => 
        array (
            0 => __DIR__ . '/..' . '/predis/predis/src',
        ),
    );

    public static $classMap = array (
        'Composer\\InstalledVersions' => __DIR__ . '/..' . '/composer/InstalledVersions.php',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->prefixLengthsPsr4 = ComposerStaticInitd6b93bc241e89d4064739e7ff51855e0::$prefixLengthsPsr4;
            $loader->prefixDirsPsr4 = ComposerStaticInitd6b93bc241e89d4064739e7ff51855e0::$prefixDirsPsr4;
            $loader->classMap = ComposerStaticInitd6b93bc241e89d4064739e7ff51855e0::$classMap;

        }, null, ClassLoader::class);
    }
}

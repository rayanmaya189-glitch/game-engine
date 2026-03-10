// Top-level build file where you can add configuration options common to all sub-projects/modules.
buildscript {
    ext {
        kotlin_version = '1.9.22'
        compose_version = '1.5.8'
        hilt_version = '2.50'
        room_version = '2.6.1'
        retrofit_version = '2.9.0'
        okhttp_version = '4.12.0'
        coroutines_version = '1.7.3'
    }
}

plugins {
    id 'com.android.application' version '8.2.2' apply false
    id 'com.android.library' version '8.2.2' apply false
    id 'org.jetbrains.kotlin.android' version '1.9.22' apply false
    id 'com.google.dagger.hilt.android' version '2.50' apply false
    id 'com.google.protobuf' version '0.9.4' apply false
}

task clean(type: Delete) {
    delete rootProject.buildDir
}

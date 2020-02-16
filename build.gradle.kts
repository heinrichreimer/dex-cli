plugins {
    kotlin("jvm") version "1.3.61"
    id("com.google.protobuf") version "0.8.11"
}

group = "dev.reimer"
version = "0.1.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))
}

protobuf {

}

//protobuf {
//    protoc {
//        artifact = "com.google.protobuf:protoc:3.5.1-1"
//    }
//    plugins {
//        grpc { artifact = "io.grpc:protoc-gen-grpc-java:${grpcVersion}" }
//    }
//    generateProtoTasks {
//        all()*.plugins { grpc {} }
//    }
//}

tasks {
    compileKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
    compileTestKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
}